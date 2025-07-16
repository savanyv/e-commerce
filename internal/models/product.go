package models

import (
	"time"
)

type Product struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(255); not null"`
	Price float64 `json:"price" gorm:"type:float; not null"`
	Quantity int `json:"quantity" gorm:"type:int; not null"`
	IDBrand uint `json:"id_brand" gorm:"not null"`
	Brand Brand `json:"brand" gorm:"foreignKey:IDBrand"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
