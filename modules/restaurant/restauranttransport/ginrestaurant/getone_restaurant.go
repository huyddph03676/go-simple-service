package ginrestaurant

import (
	"go-simple-service/common"
	"go-simple-service/component"
	"go-simple-service/modules/restaurant/restaurantbiz"
	"go-simple-service/modules/restaurant/restaurantstore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOneRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetOneRestaurantBiz(store)

		result, err := biz.GetOneRestaurant(ctx.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		result.Mask(false)

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))
	}
}
