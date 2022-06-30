package entity

type Facility struct {
	Id     int      `json:"fid"`
	Name   string   `gorm:"size:32"`
	Houses []*House
}
