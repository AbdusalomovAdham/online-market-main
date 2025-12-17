package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel `bun:"table:products"`

	Id              int64      `json:"id" bun:"id,pk,autoincrement"`
	Description     string     `json:"description" bun:"description" default:""`
	Price           float64    `json:"price" bun:"price" default:"0"`
	StockQuantity   int64      `json:"stock_quantity" bun:"stock_quantity" default:"0"`
	Rating          int8       `json:"rating" bun:"rating" default:"0"`
	Status          bool       `json:"status" bun:"status" default:"true"`
	SellerId        int64      `json:"seller_id" bun:"seller_id, notnull"`
	CategoryId      int64      `json:"category_id" bun:"category_id" default:"0"`
	ViewsCount      int64      `json:"views_count" bun:"views_count" default:"0"`
	DiscountPercent int8       `json:"discount_percent" bun:"discount_percent" default:"0"`
	Images          []File     `json:"images" bun:"images" default:"[]"`
	CreatedAt       time.Time  `json:"created_at" bun:"created_at"`
	CreatedBy       *string    `json:"created_by" bun:"created_by"`
	UpdatedAt       *time.Time `json:"updated_at" bun:"updated_at"`
	UpdatedBy       *string    `json:"updated_by" bun:"updated_by"`
	DeletedAt       *time.Time `json:"deleted_at" bun:"deleted_at"`
	DeletedBy       *string    `json:"deleted_by" bun:"deleted_by"`
}
