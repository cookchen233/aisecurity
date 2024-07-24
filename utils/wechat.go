package utils

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/util"
	"go.uber.org/zap"
	"os"
)

var WechatOfficialAccount *officialaccount.OfficialAccount

func InitWechatOfficialAccount() {
	wc := wechat.NewWechat()
	memory, err := NewFileCache("wechat.cache")
	if err != nil {
		Logger.Error("NewFileCache error", zap.Error(err))
	}
	WechatOfficialAccount = wc.GetOfficialAccount(&config.Config{
		AppID:     os.Getenv("WECHAT_APP_ID"),
		AppSecret: os.Getenv("WECHAT_APP_SECRET"),
		Token:     os.Getenv("WECHAT_TOKEN"),
		//EncodingAESKey: "xxxx",
		Cache: memory,
	})

	//mu := WechatOfficialAccount.GetMenu()
	//buttons := []*menu.Button{
	//	{
	//		Type: "view",
	//		Name: "操作说明",
	//		URL:  "https://kunyunsafe.com/",
	//	},
	//	{
	//		Type: "view",
	//		Name: "主页",
	//		URL:  "https://demo.kunyunsafe.com/dashboard/",
	//	},
	//	{
	//		Type: "view",
	//		Name: "官网",
	//		URL:  "https://kunyunsafe.com/",
	//	},
	//}
	//err := mu.SetMenu(buttons)
	//if err != nil {
	//	fmt.Printf("err= %v", err)
	//	return
	//}

}

func GetQRCode(sceneStr string) (string, string, error) {
	accessToken, err := WechatOfficialAccount.GetAccessToken()
	if err != nil {
		return "", "", ErrorWithStack(err)
	}
	type Scene struct {
		SceneStr string `json:"scene_str"`
	}
	type ActionInfo struct {
		Scene Scene `json:"scene"`
	}

	msg := struct {
		ExpireSeconds int        `json:"expire_seconds"`
		ActionName    string     `json:"action_name"`
		ActionInfo    ActionInfo `json:"action_info"`
	}{
		ExpireSeconds: 604800,
		ActionName:    "QR_LIMIT_STR_SCENE",
		ActionInfo:    ActionInfo{Scene: Scene{SceneStr: sceneStr}},
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s", accessToken)
	var response []byte
	response, err = util.PostJSON(uri, msg)
	g.Dump("response-----------------------", response, err)
	if err != nil {
		return "", "", ErrorWithStack(err)
	}

	type respResult struct {
		util.CommonError
		Ticket string `json:"ticket"`
		Url    string `json:"url"`
	}
	var result respResult
	err = util.DecodeWithError(response, &result, "createQRCode")
	if err != nil {
		return "", "", ErrorWithStack(err)
	}
	return result.Ticket, result.Url, nil

}
