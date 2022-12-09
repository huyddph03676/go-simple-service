package restaurantbiz

import (
	"context"
	"errors"
	"go-simple-service/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	GetOne(ctx context.Context, id int) (*restaurantmodel.Restaurant, error)
	Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	oldData, err := biz.store.GetOne(ctx, id)

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
