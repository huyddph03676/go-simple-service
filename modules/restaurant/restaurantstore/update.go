package restaurantstore

import (
	"context"
	"go-simple-service/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	db := s.db
	db = db.Table(restaurantmodel.RestaurantUpdate{}.TableName()).Where("id = ?", id)

	if err := db.Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
