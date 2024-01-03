package request

type Admin struct {
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 用户名
	Username string `json:"username" validate:"required"`
	// 密码
	Password string `json:"password" validate:""`
	// 名字
	Name        string `json:"name" validate:"required"`
	AdminRoleID int    `json:"admin_role_id" validate:"required"`
}

type AdminLogin struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	PersistDays int    `json:"persist_days"`
}
