package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Price struct {
	bun.BaseModel `bun:"table:prices"`

	Id             int     `json:"id" bun:"id,pk,autoincrement"`
	Tariff         string  `json:"tariff" bun:"tariff"`
	FromDistrictId int     `json:"from_district_id" bun:"from_district_id"`
	ToDistrictId   int     `json:"to_district_id" bun:"to_district_id"`
	Amount         float64 `json:"amount" bun:"amount"`

	CreatedBy int       `json:"created_by" bun:"created_by,default:null"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
	UpdatedBy int       `json:"updated_by" bun:"updated_by,default:null"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,default:null"`
}
