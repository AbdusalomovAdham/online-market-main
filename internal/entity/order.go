package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Address struct {
	Longitude float64 `json:"longitude" bun:"longitude"`
	Latitude  float64 `json:"latitude" bun:"latitude"`
}

type Places struct {
	Id    int     `json:"id" bun:"id,pk,autoincrement"`
	Price float64 `json:"price" bun:"price"`
}

type Order struct {
	bun.BaseModel `bun:"table:orders"`

	Id       int  `json:"id" bun:"id,pk,autoincrement"`
	ClientId int  `json:"client_id" bun:"client_id"`
	DriverId *int `json:"driver_id" bun:"driver_id"`
	TariffId int  `json:"tariff_id" bun:"tariff_id"`

	Price    float64 `json:"price" bun:"price"`
	Distance float64 `json:"distance" bun:"distance"`
	StatusId int     `json:"status_id" bun:"status_id"`

	FromRegionId   int     `json:"from_region_id" bun:"from_region_id"`
	ToRegionId     int     `json:"to_region_id" bun:"to_region_id"`
	FromDistrictId int     `json:"from_district_id" bun:"from_district_id"`
	ToDistrictId   int     `json:"to_district_id" bun:"to_district_id"`
	FromAddress    Address `json:"from_address" bun:"from_address"`
	ToAddress      Address `json:"to_address" bun:"to_address"`

	Places     []int `json:"places" bun:"places"`
	PlaceCount int   `json:"place_count" bun:"place_count"`

	StartTime time.Time `json:"start_time" bun:"start_time"`
	EndTime   time.Time `json:"end_time" bun:"end_time"`
	TimeSpent int       `json:"time_spent" bun:"time_spent"`

	CreatedBy int       `json:"created_by" bun:"created_by,default:null"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
	UpdatedBy int       `json:"updated_by" bun:"updated_by,default:null"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,default:null"`

	DepartureDate string `json:"departure_date" bun:"departure_date"`
}
