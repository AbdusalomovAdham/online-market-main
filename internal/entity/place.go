package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Place struct {
	bun.BaseModel `bun:"table:places"`

	Id        int       `json:"id" bun:"id,pk,autoincrement"`
	Price     float64   `json:"price" bun:"price"`
	CreatedBy int       `json:"created_by" bun:"created_by,default:null"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
	UpdatedBy int       `json:"updated_by" bun:"updated_by,default:null"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,default:null"`
}

// id
//1 - front seat
//2 - back-rigth seat
//3 - back-left seat
