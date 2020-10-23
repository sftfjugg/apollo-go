package models

import "apollo-adminserivce/internal/pkg/models"

type AppNamespace struct {
	Name       string       `json:"name"`
	AppId      string       `json:"app_id"`
	AppName    string       `json:"app_name"`
	Namespaces []*Namespace `json:"namespaces"`
}

type Namespace struct {
	Id          uint64         `json:"id"`
	ClusterName string         `json:"cluster_name"` //灰度使用
	LaneName    string         `json:"lane_name"`
	Items       []*models.Item `json:"items"`
}
