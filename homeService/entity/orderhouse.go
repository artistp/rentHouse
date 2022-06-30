package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type OrderHouse struct {
	gorm.Model
	UserId      uint      `json:"user_id"`
	HouseId     uint      `json:"house_id"`
	Begin_date  time.Time `gorm:"type:datetime"`
	End_date    time.Time `gorm:"type:datetime"`
	Days        int
	House_price int
	Amount      int
	Status      string `gorm:"default:'WAIT_ACCEPT'"`
	Comment     string `gorm:"size:512"`
	Credit      bool
}
