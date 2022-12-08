package restaurantstore

import (
	"context"
	"go-simple-service/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) GetOne(ctx context.Context, id int) (restaurantmodel.RestaurantCreate, error) {
	db := s.db
	var result restaurantmodel.RestaurantCreate

	if err := db.Table(restaurantmodel.RestaurantCreate{}.TableName()).Where("id = ?", id).First(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}
