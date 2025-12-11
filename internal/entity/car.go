package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Car struct {
	bun.BaseModel `bun:"table:cars"`

	Id        int       `json:"id" bun:"id,pk,autoincrement"`
	Color     string    `json:"color" bun:"color,default:null"`
	Number    string    `json:"number" bun:"number,default:null"`
	Type      string    `json:"type" bun:"type,default:null"`
	Level     int       `json:"level" bun:"level,default:null"`
	CreatedBy int       `json:"created_by" bun:"created_by,default:null"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
	UpdatedBy int       `json:"updated_by" bun:"updated_by,default:null"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,default:null"`
}
