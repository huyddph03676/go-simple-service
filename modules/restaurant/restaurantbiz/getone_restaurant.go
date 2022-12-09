package restaurantbiz

import (
	"context"
	"go-simple-service/common"
	"go-simple-service/modules/restaurant/restaurantmodel"
)

type GetOneRestaurantStore interface {
	GetOne(ctx context.Context, id int) (*restaurantmodel.Restaurant, error)
}

type getOneRestaurantBiz struct {
	store GetOneRestaurantStore
}

func NewGetOneRestaurantBiz(store GetOneRestaurantStore) *getOneRestaurantBiz {
	return &getOneRestaurantBiz{store: store}
}

func (biz *getOneRestaurantBiz) GetOneRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	result, err := biz.store.GetOne(ctx, id)

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	return result, err
}
