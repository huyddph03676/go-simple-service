package common

import "time"

type SQLModel struct {
	Id      int        `json:"id" gorm:"column:id;"`
	Status  int        `json:"status" gorm:"column:status;"`
	Created *time.Time `json:"created_at" gorm:"autoCreateTime; column:created_at"`
	Updated *time.Time `json:"updated_at" gorm:"autoCreateTime; column:updated_at"`
}
