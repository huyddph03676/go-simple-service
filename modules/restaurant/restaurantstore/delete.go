package restaurantstore

import (
	"context"
	"go-simple-service/common"
	"go-simple-service/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	db := s.db
	db = db.Table(restaurantmodel.RestaurantCreate{}.TableName()).Where("id = ?", id)

	if err := db.Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
