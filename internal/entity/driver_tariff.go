package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type DriverTariff struct {
	bun.BaseModel `bun:"table:driver_tariffs"`

	Id       int `json:"id" bun:"id,pk,autoincrement"`
	DriverId int `json:"driver_id" bun:"driver_id"`
	TarifId  int `json:"tarif_id" bun:"tarif_id"`

	CreatedAt time.Time `json:"created_at" bun:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at"`
}
