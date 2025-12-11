package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Region struct {
	bun.BaseModel `bun:"table:regions"`

	Id   int    `json:"id" bun:"id,pk,autoincrement"`
	Name string `json:"name" bun:"name"`

	CreatedBy int       `json:"created_by" bun:"created_by,default:null"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
	UpdatedBy int       `json:"updated_by" bun:"updated_by,default:null"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,default:null"`
}
