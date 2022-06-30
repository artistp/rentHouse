package entity

import "github.com/jinzhu/gorm"

type House struct {
	gorm.Model
	UserID          uint
	AreaID          uint
	Title           string        `gorm:"size:64"`
	Address         string        `gorm:"size:512"`
	Room_count      int           `gorm:"default:1"`
	Acreage         int           `gorm:"default:0" json:"acreage"`
	Price           int           `json:"price"`
	Unit            string        `gorm:"size:32;default:''" json:"unit"`
	Capacity        int           `gorm:"default:1" json:"capacity"`
	Beds            string        `gorm:"size:64;default:''" json:"beds"`
	Deposit         int           `gorm:"default:0" json:"deposit"`
	Min_days        int           `gorm:"default:1" json:"min_days"`
	Max_days        int           `gorm:"default:0" json:"max_days"`
	Order_count     int           `gorm:"default:0" json:"order_count"`
	Index_image_url string        `gorm:"size:256;default:''"`
	Facilities      []*Facility   `gorm:"many2many:house_facilities"`
	Image           []*HouseImage `json:"img_urls"`
	Order           []*OrderHouse `json:"orders"`
}
