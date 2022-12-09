package restaurantmodel

import (
	"errors"
	"go-simple-service/common"
	"strings"
	"time"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            *string `json:"name" gorm:"column:name;"`
	Addr            *string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            *string `json:"name" gorm:"column:name;"`
	Addr            *string `json:"addr" gorm:"column:addr;"`
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
