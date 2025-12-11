package order

import (
	"context"
	"fmt"
	"main/internal/services/order"

	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) Create(ctx context.Context, data order.Create) (order.Get, error) {
	var detailOrder order.Get
	var statusId int
	query := `
	INSERT INTO orders (
		client_id, tariff_id, price, distance,
		 from_district_id, to_district_id,
		places, place_count, status_id, departure_date
	) VALUES (
		?, ?, ?, ?,
		?, ?, ?, ?,
		?, ?
	)
	RETURNING
		id, client_id, tariff_id, price, distance, from_district_id, to_district_id,
		places, place_count, status_id, departure_date
	`

	err := r.DB.QueryRowContext(ctx, query,
		data.ClientId,
		data.TariffId,
		data.Price,
		data.Distance,
		data.FromDistrictId,
		data.ToDistrictId,
		data.Place,
		data.PlaceCount,
		1,
		data.DepartureDate,
	).Scan(
		&detailOrder.Id,
		&detailOrder.ClientId,
		&detailOrder.TariffId,
		&detailOrder.Price,
		&detailOrder.Distance,
		&detailOrder.FromDistrictId,
		&detailOrder.ToDistrictId,
		&detailOrder.PlaceId,
		&detailOrder.PlaceCount,
		&statusId,
		&detailOrder.DepartureDate,
	)

	detailOrder.Status = &order.Status{
		Id: &statusId,
	}

	if err != nil {
		return order.Get{}, err
	}

	query = `SELECT name FROM districts WHERE id = ?`
	err = r.DB.QueryRowContext(ctx, query, detailOrder.FromDistrictId).Scan(&detailOrder.FromDistrictName)
	if err != nil {
		return order.Get{}, err
	}

	err = r.DB.QueryRowContext(ctx, query, detailOrder.ToDistrictId).Scan(&detailOrder.ToDistrictName)
	if err != nil {
		return order.Get{}, err
	}

	query = `SELECT
		r.id,
		r.name
		FROM districts d
		LEFT JOIN regions r ON d.region_id = r.id
		WHERE d.id = ?`
	err = r.DB.QueryRowContext(ctx, query, detailOrder.FromDistrictId).Scan(&detailOrder.FromRegionId, &detailOrder.FromRegionName)
	if err != nil {
		return order.Get{}, err
	}

	err = r.DB.QueryRowContext(ctx, query, detailOrder.ToDistrictId).Scan(&detailOrder.ToRegionId, &detailOrder.ToRegionName)
	if err != nil {
		return order.Get{}, err
	}

	var tariff order.Tariff

	query = `SELECT id, tariff_name, level FROM tariffs WHERE id = ?`
	err = r.DB.QueryRowContext(ctx, query, detailOrder.TariffId).Scan(&tariff.Id, &tariff.TariffName, &tariff.Level)
	if err != nil {
		return order.Get{}, err
	}

	detailOrder.Tariff = tariff

	return detailOrder, nil
}

func (r Repository) GetList(ctx context.Context, clientId int) ([]order.Get, error) {
	var orders []order.Get

	query := fmt.Sprintf(`
		SELECT
			o.id,
			o.client_id,

			d.id AS driver_id,
			d.avatar->>'path' AS avatar,
			d.username AS driver_username,
			d.email AS driver_email,

			c.id AS car_id,
			c.color AS car_color,
			c.number AS car_number,
			c.type AS car_type,
			c.level AS car_level,

			t.id AS tariff_id,
			t.tariff_name AS tariff_name,
			t.level AS tariff_level,


			fr.id AS from_region_id,
			tr.id AS to_region_id,
			o.from_district_id,
			o.to_district_id,
			fr.name AS from_region_name,
			tr.name AS to_region_name,
			fd.name AS from_district_name,
			td.name AS to_district_name,

			o.place_count,
			o.price,
			o.places,
			o.distance,
			o.departure_date,

			s.id AS status_id,
			s.name AS status_name

			FROM orders o
				LEFT JOIN districts fd ON o.from_district_id = fd.id
				LEFT JOIN districts td ON o.to_district_id = td.id
				LEFT JOIN regions fr ON fd.region_id = fr.id
				LEFT JOIN regions tr ON td.region_id = tr.id
				LEFT JOIN tariffs t ON o.tariff_id = t.id
				LEFT JOIN drivers d ON d.id = o.driver_id
				LEFT JOIN cars c ON d.car_id = c.id
				LEFT JOIN statuses s ON o.status_id = s.id
			WHERE o.client_id = ?

	`)

	rows, err := r.DB.QueryContext(ctx, query, clientId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o order.Get
		var d order.Driver
		var c order.Car
		var t order.Tariff
		var s order.Status

		if err := rows.Scan(
			&o.Id,
			&o.ClientId,

			&d.Id,
			&d.Avatar,
			&d.Username,
			&d.Email,

			&c.Id,
			&c.Color,
			&c.Number,
			&c.Type,
			&c.Level,

			&t.Id,
			&t.TariffName,
			&t.Level,

			&o.FromRegionId,
			&o.ToRegionId,
			&o.FromDistrictId,
			&o.ToDistrictId,
			&o.FromRegionName,
			&o.ToRegionName,
			&o.FromDistrictName,
			&o.ToDistrictName,

			&o.PlaceCount,
			&o.Price,
			&o.Places,
			&o.Distance,
			&o.DepartureDate,

			&s.Id,
			&s.Name,
		); err != nil {
			return nil, err
		}

		o.Driver = &d
		o.Car = &c
		o.Status = &s
		o.Tariff = t

		orders = append(orders, o)
	}
	return orders, nil

}

func (r Repository) Update(ctx context.Context, data order.Update) error {
	query := `
	UPDATE orders SET status_id = ?, updated_at = NOW(), driver_id = ? WHERE id = ?
`
	_, err := r.DB.ExecContext(ctx, query, data.StatusId, data.DriverId, data.OrderId)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) UpdateDriverLocation(ctx context.Context, data order.UpdateDriverLocation) error {

	query := `UPDATE drivers SET longitude = ?, latitude = ? WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, data.Longitude, data.Latitude, data.DriverId)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetDriverListByDistrictId(ctx context.Context, fromDistrictId, toDistrictId int) ([]int, error) {
	query := `
	 	SELECT id FROM drivers WHERE from_district_id = ? AND to_district_id = ?
	`
	var driverIds []int

	rows, err := r.DB.QueryContext(ctx, query, fromDistrictId, toDistrictId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int

		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		driverIds = append(driverIds, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return driverIds, nil
}
