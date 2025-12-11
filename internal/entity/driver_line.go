package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type DriverLine struct {
	bun.BaseModel  `bun:"table:driver_lines"`
	ID             int       `json:"id"`
	DriverID       int       `json:"driver_id"`
	FromDistrictId int       `json:"from_district_id"`
	ToDistrictId   int       `json:"to_district_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
