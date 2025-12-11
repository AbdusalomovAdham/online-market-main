package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Tariff struct {
	bun.BaseModel `bun:"table:tariffs"`

	Id        int       `json:"id" bun:"id,pk,autoincrement"`
	Name      string    `json:"tariff_name" bun:"tariff_name"`
	Status    string    `json:"status" bun:"status"`
	Level     string    `json:"level" bun:"level"`
	CreatedBy int       `json:"created_by" bun:"created_by,default:null"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
	UpdatedBy int       `json:"updated_by" bun:"updated_by,default:null"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,default:null"`
}
