package restaurantmodel

import (
	"errors"
	"strings"
	"time"
)

type RestaurantCreate struct {
	Id      int       `json:"id" gorm:"column:id;"`
	Name    *string   `json:"name" gorm:"column:name;"`
	Addr    *string   `json:"addr" gorm:"column:addr;"`
	Created time.Time `json:"created_at" gorm:"autoCreateTime; column:created_at"`
	Updated time.Time `json:"updated_at" gorm:"autoCreateTime; column:updated_at"`
}

func (RestaurantCreate) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name    *string    `json:"name" gorm:"column:name"`
	Addr    *string    `json:"addr" gorm:"column:addr"`
	Updated *time.Time `gorm:"autoUpdateTime;column:updated_at"`
}

func (RestaurantUpdate) TableName() string {
	return RestaurantCreate{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	*res.Name = strings.TrimSpace(*res.Name)

	if len(*res.Name) == 0 {
		return errors.New("restaurant cannot be blank")
	}

	return nil

}
