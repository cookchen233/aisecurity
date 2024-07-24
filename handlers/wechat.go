package handlers

import (
	"aisecurity/ent/dao"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	http2 "aisecurity/utils/http"
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/silenceper/wechat/v2/util"
	"go.uber.org/zap"
	"hash/crc32"
	"net/http"
	"os"
	"sort"
	"strings"
)

type WechatHandler struct {
	Handler
	Service     *services.AdminService
	userService *services.UserService
}

func NewWechatHandler() *WechatHandler { return &WechatHandler{} }
func (h *WechatHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *WechatHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *WechatHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *WechatHandler) SetRequestContext(c *gin.Context, childHandler IHandler) {
	h.Service = services.NewAdminService(c)
	h.Filter = &filters.Event{}
	h.Entity = &entities.Event{}
	h.userService = services.NewUserService(c)
}

func (h *WechatHandler) checkSign(c *gin.Context) bool {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")

	arr := []string{timestamp, nonce, os.Getenv("WECHAT_TOKEN")}
	sort.Strings(arr)
	hash := sha1.New()
	hash.Write([]byte(strings.Join(arr, "")))
	signature2 := fmt.Sprintf("%x", hash.Sum(nil))
	if signature != signature2 {
		utils.Logger.Error("validate failed", zap.String("signature", signature), zap.Any("arr", arr))
		return false
	}
	return true
}

func (h *WechatHandler) OfficialAccount(c *gin.Context) {
	if !h.checkSign(c) {
		c.String(http.StatusOK, "签名验证失败")
		return
	}
	echostr := c.Query("echostr")
	if echostr != "" {
		c.String(http.StatusOK, echostr)
		return
	}
	server := utils.WechatOfficialAccount.GetServer(c.Request, c.Writer)
	// 设置接收消息的处理方法
	server.SetMessageHandler(h.handleMsg)
	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		utils.Logger.Error("OfficialAccount serve error", zap.Error(err))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	//发送回复的消息
	c.JSON(http.StatusOK, server.Send())
}

// HandleMsg 处理用户发送的消息
func (h *WechatHandler) handleMsg(msg *message.MixMessage) *message.Reply {
	switch msg.MsgType {
	case message.MsgTypeText: // 文本消息
		g.Dump("x1")
		return h.handleTextMsg(msg)
	case message.MsgTypeEvent: // 事件推送
		g.Dump("x2")

		return h.handleEventMsg(msg)
	default:
		return h.replyMessage("未知的消息类型")
	}
}

var availableCmds = []string{
	"1. 绑定手机号[你的手机号]",
}

// HandleEventMsg 处理事件推送
func (h *WechatHandler) handleEventMsg(msg *message.MixMessage) *message.Reply {
	switch msg.Event {
	case message.EventSubscribe: // 用户订阅
		if strings.HasPrefix(msg.EventKey, "create-admin_") {
			return h.createAdminByQRCode(msg)
		}
		_, err := h.userService.CreateByWechatOpenidIfNotExists(string(msg.FromUserName))
		if err != nil {
			utils.Logger.Error("CreateByWechatOpenidIfNotExists error", zap.Error(err))
		}
		return h.replyMessage("感谢关注重庆鲲云安全科技")
	case message.EventScan:
		if strings.HasPrefix(msg.EventKey, "create-admin_") {
			return h.createAdminByQRCode(msg)
		}
		return h.replyMessage("未知的扫一扫事件")
	case message.EventUnsubscribe: // 用户取消订阅
		return h.replyMessage("已成功取消了关注")

	default:
		return h.replyMessage("未知的事件")
	}
}

// HandleEventMsg 处理互动消息
func (h *WechatHandler) handleTextMsg(msg *message.MixMessage) *message.Reply {
	if strings.Contains(msg.Content, "绑定手机号") {
		return h.bindMobile(msg)
	}
	return h.replyMessage("重庆鲲云安全科技竭诚为您服务")
}

// BindMobile 绑定手机号
func (h *WechatHandler) bindMobile(msg *message.MixMessage) *message.Reply {
	user, err := h.userService.CreateByWechatOpenidIfNotExists(string(msg.FromUserName))
	if err != nil {
		return h.replyError(utils.ErrorWithStack(err))
	}
	var u = user.(*entities.User)
	u.Mobile = strings.ReplaceAll(msg.Content, "绑定手机号", "")
	_, err = h.Service.Update(u)
	if err != nil {
		return h.replyError(utils.ErrorWithStack(err))
	}
	return h.replyMessage("成功绑定了手机号")
}

// bindOpenid 绑定微信
func (h *WechatHandler) bindOpenid(msg *message.MixMessage) *message.Reply {
	sceneStrData := strings.Split(strings.Replace(msg.EventKey, "bind-openid_", "", 1), "_")
	var admID = gconv.Int(sceneStrData[1])
	var openid = string(msg.FromUserName)
	ex, err := h.Service.GetByWechatOpenid(openid)
	if err != nil {
		return h.replyError(utils.ErrorWithStack(err))
	}
	if ex == nil {
		return h.replyMessage("您当前还没有成为我们的用户, 请你先咨询工作人员获取邀请二维码进行注册")
	}
	adm := ex.(*entities.Admin)
	if adm.ID != admID {
		adm.WechatOpenid = ""
		_, err := h.Service.Update(adm)
		if err != nil {
			return h.replyError(utils.ErrorWithStack(err))
		}
		admin, err := h.Service.GetByID(admID)
		if err != nil {
			return h.replyError(utils.ErrorWithStack(err))
		}
		adm2 := admin.(*entities.Admin)
		adm2.WechatOpenid = openid
		_, err = h.Service.Update(adm2)
		if err != nil {
			return h.replyError(utils.ErrorWithStack(err))
		}
	}
	return h.replyMessage("成功绑定了微信")
}

// createAdminByQRCode 二维码邀请创建管理员
func (h *WechatHandler) createAdminByQRCode(msg *message.MixMessage) *message.Reply {
	sceneStrData := strings.Split(strings.Replace(msg.EventKey, "create-admin_", "", 1), "_")
	_, err := h.Service.GetByWechatOpenid(string(msg.FromUserName))
	if err != nil {
		if !dao.IsNotFound(err) {
			return h.replyError(utils.ErrorWithStack(err))
		}
	} else {
		return h.replyMessage("您已经注册过该账号了, 请直接点击主页登录")
	}
	var password = fmt.Sprintf("%08x", crc32.ChecksumIEEE([]byte(uuid.NewString())))
	g.Dump("username", string([]rune(msg.FromUserName)[15:]))
	create, err := h.Service.Create(&entities.Admin{
		Admin: dao.Admin{
			Username:     string([]rune(msg.FromUserName)[15:]),
			Nickname:     string([]rune(msg.FromUserName)[15:]),
			WechatOpenid: string(msg.FromUserName),
			CreatorID:    gconv.Int(sceneStrData[1]),
			Password:     password,
		},
	})
	if err != nil {
		return h.replyError(utils.ErrorWithStack(err))
	}

	return h.replyMessage("账号创建成功, 请登录后及时修改密码\n登录名: " + create.(*entities.Admin).Username + "\n密码: " + password)
}

func (h *WechatHandler) replyMessage(msg string) *message.Reply {
	return &message.Reply{
		MsgType: message.MsgTypeText,
		MsgData: message.NewText(msg),
	}
}

func (h *WechatHandler) replyError(err error) *message.Reply {
	utils.Logger.Error("handleMsg error", zap.Error(err))
	return &message.Reply{
		MsgType: message.MsgTypeText,
		MsgData: message.NewText("非常抱歉, 当前服务器发生错误, 请稍后再试"),
	}
}

func (h *WechatHandler) AdminOauth2(c *gin.Context) {
	accessInfo, err := getAccessInfo(c.Query("code"))
	if err != nil {
		http2.Error(c, utils.ErrorWithStack(err), 2000)
		return
	}

	_, err = h.Service.GetByWechatOpenid(accessInfo.Openid)
	if err != nil && !dao.IsNotFound(err) {
		if !dao.IsNotFound(err) {
			http2.Error(c, utils.ErrorWithStack(err), 2000)
			return
		}
	}

	_, err = getUserInfo(accessInfo.AccessToken, accessInfo.Openid)
	if err != nil {
		http2.Error(c, utils.ErrorWithStack(err), 2000)
		return
	}
	server := utils.WechatOfficialAccount.GetServer(c.Request, c.Writer)

	c.JSON(http.StatusOK, server.Send())
}

type accessInfo struct {
	util.CommonError
	AccessToken    string `json:"access_token"`
	ExpiresIn      string `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	Openid         string `json:"openid"`
	Scope          string `json:"scope"`
	IsSnapshotuser string `json:"is_snapshotuser"`
	Unionid        string `json:"unionid"`
}

func getAccessInfo(code string) (*accessInfo, error) {
	uri := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		os.Getenv("WECHAT_APP_ID"), os.Getenv("WECHAT_APP_SECRET"), code)
	var response []byte
	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}

	var result *accessInfo
	err = util.DecodeWithError(response, result, "access_token")
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	return result, nil
}

type userInfo struct {
	Openid     string `json:"openid"`
	Nickname   string `json:"nickname"`
	Sex        int    `json:"sex"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Headimgurl string `json:"headimgurl"`
	Unionid    string `json:"unionid"`
}

func getUserInfo(accessToken string, openid string) (*userInfo, error) {
	uri := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN",
		accessToken, openid)
	var response []byte
	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}

	var result *userInfo
	err = util.DecodeWithError(response, result, "userinfo")
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	return result, nil
}
