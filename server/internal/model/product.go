package model

type Products struct {
	Name         string `gorm:"not null" json:"name"`
	ProductPhoto string `gorm:"not null" json:"product_photo"`
	Price        int    `gorm:"not null" json:"price"`
}
