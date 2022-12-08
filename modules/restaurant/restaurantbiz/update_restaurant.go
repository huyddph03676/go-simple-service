package restaurantbiz

import (
	"context"
	"go-simple-service/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	err := biz.store.Update(ctx, id, data)

	return err
}
