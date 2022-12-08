package ginrestaurant

import (
	"go-simple-service/common"
	"go-simple-service/component"
	"go-simple-service/modules/restaurant/restaurantbiz"
	"go-simple-service/modules/restaurant/restaurantstore"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOneRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetOneRestaurantBiz(store)

		result, err := biz.GetOneRestaurant(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))
	}
}
