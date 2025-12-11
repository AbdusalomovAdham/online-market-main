package order

type Tariff struct {
	Id         *int    `json:"id"`
	TariffName *string `json:"tariff_name"`
	Level      *int    `json:"level"`
}

type Create struct {
	ID             string  `json:"id"`
	ClientId       int     `json:"client_id"`
	DriverId       *int    `json:"driver_id"`
	FromAddress    string  `json:"from_address"`
	ToAddress      string  `json:"to_address"`
	FromDistrictId int     `json:"from_district_id"`
	ToDistrictId   int     `json:"to_district_id"`
	TariffId       int     `json:"tariff_id"`
	Place          int     `json:"places"`
	PlaceCount     int     `json:"place_count"`
	Price          float64 `json:"price"`
	Distance       float64 `json:"distance"`
	DepartureDate  string  `json:"departure_date"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
}

type Get struct {
	Id       int `json:"id"`
	ClientId int `json:"client_id"`

	FromRegionId     int    `json:"from_region_id"`
	FromRegionName   string `json:"from_region_name"`
	ToRegionId       int    `json:"to_region_id"`
	ToRegionName     string `json:"to_region_name"`
	FromDistrictId   int    `json:"from_district_id"`
	FromDistrictName string `json:"from_district_name"`
	ToDistrictId     int    `json:"to_district_id"`
	ToDistrictName   string `json:"to_district_name"`

	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`

	TariffId      int     `json:"tariff_id"`
	Tariff        Tariff  `json:"tariff"`
	Places        int     `json:"places"`
	PlaceId       int     `json:"place_id"`
	PlaceCount    int     `json:"place_count"`
	Price         float64 `json:"price"`
	Distance      float64 `json:"distance"`
	Status        *Status `json:"status"`
	Driver        *Driver `json:"driver"`
	Car           *Car    `json:"car"`
	DepartureDate string  `json:"departure_date"`
}

type Update struct {
	DriverId *int `json:"driver_id"`
	StatusId *int `json:"status_id"`
	OrderId  *int `json:"order_id"`
}

type Driver struct {
	Id       *int    `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Avatar   *string `json:"avatar"`
}

type Car struct {
	Id     *int    `json:"id"`
	Color  *string `json:"color"`
	Number *string `json:"number"`
	Type   *string `json:"type"`
	Level  *int    `json:"level"`
}

type Status struct {
	Id   *int    `json:"id"`
	Name *string `json:"name"`
}

type UpdateDriverLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	DriverId  int     `json:"driver_id"`
}

type OrderUUID struct {
	OrderUUID string `json:"order_uuid"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type OrderDetail struct {
	UUID           string  `json:"uuid"`
	ToDistrictId   int     `json:"to_district_id"`
	FromDistrictId int     `json:"from_district_id"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
}
