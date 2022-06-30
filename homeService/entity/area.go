package entity

type Area struct {
	Id    int      `json:"aid"`
	Name  string   `gorm:"size:32" json:"aname"`
	Houses []*House `json:"houses"`
}
