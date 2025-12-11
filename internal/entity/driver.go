package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Driver struct {
	bun.BaseModel `bun:"table:drivers"`

	Id        int    `json:"id" bun:"id,pk,autoincrement"`
	Username  string `json:"username" bun:"username"`
	Birthdate string `json:"birth_date" bun:"birth_date"`
	Email     string `json:"email" bun:"email"`
	Status    bool   `json:"status" bun:"status,default:false"`
	Avatar    File   `json:"avatar" bun:"avatar"`
	CarId     int    `json:"car_id" bun:"car_id,default:null"`
	TariffId  int    `json:"tariff_id" bun:"tariff_id,default:null"`

	CreatedBy int       `json:"created_by" bun:"created_by,default:null"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
	UpdatedBy int       `json:"updated_by" bun:"updated_by,default:null"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,default:null"`
}
