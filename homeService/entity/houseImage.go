package entity

type HouseImage struct {
	Id int `json:"house_image_id"`
	Url string `gorm:"size:256" json:"url"`
	HouseId uint `json:"house_id"`
}
