package entity

import "github.com/uptrace/bun"

type OrderClient struct {
	bun.BaseModel `bun:"table:order_clients"`

	ID           int    `json:"id"`
	OrderID      int    `json:"order_id"`
	ClientID     int    `json:"client_id"`
	StatusId     int    `json:"status_id"`
	PlaceId      int    `json:"place_id"`
	FromLocation string `json:"from_location"`
	FromAddress  string `json:"from_address"`
	ToLocation   string `json:"to_location"`
	ToAddress    string `json:"to_address"`
}
