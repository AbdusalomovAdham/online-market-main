package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Rete struct {
	bun.BaseModel `bun:"table:retes"`

	ID            int       `json:"id"`
	OrderClientId int       `json:"order_client_id"`
	Comment       string    `json:"comment"`
	CreatedAt     time.Time `json:"created_at"`
}
