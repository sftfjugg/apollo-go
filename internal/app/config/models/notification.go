package models

type Notification struct {
	NamespaceName  string    `json:"namespaceName"`
	NotificationId int       `json:"notificationId"`
	Messages       *Messages `json:"messages"`
}

type Messages struct {
	Details map[string]int `json:"details"`
}

type Version struct {
	Max   float64
	Index int
}
