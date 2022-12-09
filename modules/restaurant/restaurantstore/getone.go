package restaurantstore

import (
	"context"
	"go-simple-service/common"
	"go-simple-service/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) GetOne(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	db := s.db
	var result *restaurantmodel.Restaurant

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
