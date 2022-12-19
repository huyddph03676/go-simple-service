package ginrestaurant

import (
	"go-simple-service/common"
	"go-simple-service/component"
	"go-simple-service/modules/restaurant/restaurantbiz"
	"go-simple-service/modules/restaurant/restaurantmodel"
	"go-simple-service/modules/restaurant/restaurantstore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data restaurantmodel.RestaurantUpdate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())

		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(ctx.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	}
}
