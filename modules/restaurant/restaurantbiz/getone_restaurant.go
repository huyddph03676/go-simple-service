package restaurantbiz

import (
	"context"
	"go-simple-service/modules/restaurant/restaurantmodel"
)

type GetOneRestaurantStore interface {
	GetOne(ctx context.Context, id int) (restaurantmodel.RestaurantCreate, error)
}

type getOneRestaurantBiz struct {
	store GetOneRestaurantStore
}

func NewGetOneRestaurantBiz(store GetOneRestaurantStore) *getOneRestaurantBiz {
	return &getOneRestaurantBiz{store: store}
}

func (biz *getOneRestaurantBiz) GetOneRestaurant(ctx context.Context, id int) (restaurantmodel.RestaurantCreate, error) {
	result, err := biz.store.GetOne(ctx, id)

	return result, err
}
