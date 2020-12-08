package models

type Role struct {
	AppId    string  `json:"app_id" form:"app_id"`
	Write    []*User `json:"write" form:"write"`
	Release  []*User `json:"release" form:"release"`
	Operator string  `json:"operator" form:"operator"`
}

type User struct {
	UserId   string `json:"user_id" form:"user_id"`
	UserName string `json:"user_name" form:"user_name"`
}
