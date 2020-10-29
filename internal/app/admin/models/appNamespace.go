package models

import "go.didapinche.com/foundation/apollo-plus/internal/pkg/models"

type AppNamespacePage struct {
	AppNamespaces []*AppNamespace `json:"app_namespaces"`
	Total         int             `json:"total"`
}

type AppNamespace struct {
	Name       string       `json:"name"`
	AppId      string       `json:"app_id"`
	Format     string       `gorm:"column:Format" json:"format" form:"format"` //类型
	Namespaces []*Namespace `json:"namespaces"`
}

type Namespace struct {
	Id          uint64         `json:"id"`
	ClusterName string         `json:"cluster_name"` //灰度使用
	LaneName    string         `json:"lane_name"`
	Items       []*models.Item `json:"items"`
}

//声明一个Hero结构体切片类型
type AppNamespaceSlice []*AppNamespace

//切片实现Interface 接口的三个方法
//1.Len() ：返回切片的大小
func (m AppNamespaceSlice) Len() int {
	return len(m)
}

//2.Less(i, j int) :决定使用什么规则进行排序
func (m AppNamespaceSlice) Less(i, j int) bool {
	return m[i].Name < m[j].Name
}

//3.Swap(i, j int) :Less(i, j int)返回true时进行交换
func (m AppNamespaceSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}