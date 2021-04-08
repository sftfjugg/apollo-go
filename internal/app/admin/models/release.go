package models

type ReleaseRequest struct {
	Name        string   `json:"name"`
	Comment     string   `json:"comment"`
	AppId       string   `json:"app_id"`
	ClusterName string   `json:"cluster_name"`
	LaneName    string   `json:"lane_name"`
	NamespaceId uint64   `json:"namespace_id"`
	Keys        []string `json:"keys"`
	Values      []string `json:"values"`
	Operator    string   `json:"operator"`
	DeptName    string   `json:"dept_name"`
}
