package properties

type ResponseCode int

const (
	Success ResponseCode = iota + 200
)
const (
	RequestError ResponseCode = iota + 900
	NotLoggedIn
	Unauthorized
)
const (
	ServerError ResponseCode = iota + 1000
)

func (p ResponseCode) Message() string {
	switch p {
	case Success:
		return "成功"
	case RequestError:
		return "请求数据错误, 请检查"
	case NotLoggedIn:
		return "未登录"
	case Unauthorized:
		return "未授权"
	case ServerError:
		return "服务器发生错误, 请稍后再试"
	default:
		return "Unknown Error"
	}
}
