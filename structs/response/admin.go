package request

type Admin struct {
	ID int `json:"id,omitempty"`
	// 用户名
	Username string `json:"username"`
	// 名字
	Name string `json:"name"`
}
