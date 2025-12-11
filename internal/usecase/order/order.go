package order

import (
	"context"
	"errors"
	"fmt"
	"main/internal/services/location"
	"main/internal/services/order"

	"github.com/google/uuid"
)

type UseCase struct {
	order Order
	auth  Auth
	ws    Websocket
	cache Cache
}

func NewUseCase(order Order, auth Auth, ws Websocket, cache Cache) *UseCase {
	return &UseCase{
		order: order,
		auth:  auth,
		ws:    ws,
		cache: cache,
	}
}

func (uc UseCase) GetList(ctx context.Context, authHeader string) ([]order.Get, error) {
	tokenClaims, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return nil, err
	}
	return uc.order.GetList(ctx, tokenClaims.Id)
}

func (uc UseCase) CreateOrder(ctx context.Context, req order.Create, authHeader string) error {
	token, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return err
	}

	uuid := uuid.New().String()
	req.ClientId = token.Id

	var order order.OrderDetail
	order.FromDistrictId = req.FromDistrictId
	order.ToDistrictId = req.ToDistrictId
	order.UUID = uuid
	order.Latitude = req.Latitude
	order.Longitude = req.Longitude

	uc.ws.Broadcast("order_created", order)

	if err = uc.cache.Set(ctx, uuid, req); err != nil {
		return err
	}
	return nil
}

func (uc UseCase) OrderAccept(ctx context.Context, authHeader string, orderUUID string) (string, error) {

	tokenClaims, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return "", err
	}

	var dest order.Create
	if err = uc.cache.Get(ctx, orderUUID, &dest); err != nil {
		return "", err
	}

	if dest.DriverId == nil {
		dest.DriverId = &tokenClaims.Id
	}

	orderDetail, err := uc.order.Create(ctx, dest)
	if err != nil {
		return "", err
	}

	statusId := 2
	var update order.Update
	update.OrderId = &orderDetail.Id
	update.StatusId = &statusId
	update.DriverId = &tokenClaims.Id

	var driverLocation order.Location
	key := fmt.Sprintf("driver:%d:location", *dest.DriverId)
	if err = uc.cache.Get(ctx, key, &driverLocation); err != nil {
		return "", err
	}

	if err = uc.order.Update(ctx, update); err != nil {
		return "", err
	}

	locationDetail, err := location.Location(driverLocation.Latitude, driverLocation.Longitude, dest.Latitude, dest.Longitude)
	if err != nil {
		return "", err
	}

	// if locationDetail.features[0].

	uc.ws.Broadcast("order_accepted", orderUUID)

	uc.ws.SendToClient(dest.ClientId, "driver accepted order", dest)

	if err := uc.cache.Delete(ctx, orderUUID); err != nil {
		return "", err
	}
	return locationDetail, nil
}

func (uc UseCase) StartTrip(ctx context.Context, orderId int, authHeader string) error {
	_, err := uc.auth.IsValidToken(ctx, authHeader)
	if err != nil {
		return err
	}

	if orderId == 0 {
		return errors.New("order id is required")
	}

	uc.ws.Broadcast("trip_started", orderId)
	return nil
}
