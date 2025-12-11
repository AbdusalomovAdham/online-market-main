package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	Id        int       `json:"id" bun:"id,pk,autoincrement"`
	Username  string    `json:"username" bun:"username"`
	Birthdate string    `json:"birth_date" bun:"birth_date"`
	Email     string    `json:"email" bun:"email"`
	Status    bool      `json:"status" bun:"status,default:false"`
	Role      int       `json:"role" bun:"role,default:3"`
	Avatar    File      `json:"avatar" bun:"avatar"`
	CreatedBy int       `json:"created_by" bun:"created_by,default:null"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
	UpdatedBy int       `json:"updated_by" bun:"updated_by,default:null"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,default:null"`
}

//Role

//1-super admin
//2-admin
//3-user
