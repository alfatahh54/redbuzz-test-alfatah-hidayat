package model

import "time"

type Transaction struct {
	ID              uint       `gorm:"column:id;primary_key" json:"id,omitempty"`
	ProductID       uint       `gorm:"column:product_id" json:"product_id"`
	CustomerName    string     `gorm:"column:customer_name" json:"customer_name"`
	TransactionCode string     `json:"transation_code"`
	Price           int        `json:"price"`
	Qty             int        `json:"qty"`
	Total           int        `json:"total"`
	CreatedAt       *time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	UpdatedAt       *time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP"`
	DeletedAt       *time.Time `gorm:"type:DATETIME;default:NULL"`
}

type CreateTransactionBody struct {
	CustomerName string            `json:"customer_name" binding:"required"`
	ProductList  []ProductListBody `json:"product_list" binding:"required"`
}

type ProductListBody struct {
	ProductID uint `json:"product_id" binding:"required"`
	Qty       int  `json:"qty" binding:"required"`
}
