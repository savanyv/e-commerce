package models

import (
	"time"
)

type Brand struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(255); not null"`
	Products []Product `json:"products,omitempty" gorm:"foreignKey:IDBrand"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
