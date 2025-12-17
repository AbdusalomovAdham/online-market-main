package wishlist

import "time"

type GetList struct {
	Id              int64     `json:"id"`
	ProductId       int64     `json:"product_id" bun:"product_id"`
	CreatedAt       time.Time `json:"created_at" bun:"created_at"`
	Name            string    `json:"name" bun:"name"`
	UserId          int64     `json:"user_id" bun:"user_id"`
	Price           float64   `json:"price" bun:"price"`
	Images          any       `json:"images" bun:"images"`
	ViewsCount      int64     `json:"views_count" bun:"views_count"`
	DiscountPercent int64     `json:"discount_percent" bun:"discount_percent"`
	CategoryId      int64     `json:"category_id" bun:"category_id"`
	Rating          int8      `json:"rating" bun:"rating"`
	Description     string    `json:"description" bun:"description"`
}

type Create struct {
	ProductId int64 `json:"product_id" bun:"product_id"`
}
