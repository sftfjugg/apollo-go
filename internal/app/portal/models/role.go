package models

type Role struct {
	AppId     string  `json:"app_id" form:"app_id"`
	Namespace string  `json:"namespace" form:"namespace"`
	Cluster   string  `json:"cluster" form:"cluster"`
	Env       string  `json:"env" form:"env"`
	Write     []*User `json:"write" form:"write"`
	Release   []*User `json:"release" form:"release"`
	Operator  string  `json:"operator" form:"operator"`
}

type User struct {
	UserId   string `json:"user_id" form:"user_id"`
	UserName string `json:"user_name" form:"user_name"`
}

type NamespaceRole struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type Auth struct {
	IsOwner bool             `json:"is_owner"`
	Role    []*NamespaceRole `json:"role"`
}
