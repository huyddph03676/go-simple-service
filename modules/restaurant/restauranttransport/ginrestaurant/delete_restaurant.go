package ginrestaurant

import (
	"go-simple-service/common"
	"go-simple-service/component"
	"go-simple-service/modules/restaurant/restaurantbiz"
	"go-simple-service/modules/restaurant/restaurantstore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			ctx.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())

		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(ctx.Request.Context(), int(uid.GetLocalID())); err != nil {
			ctx.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	}
}
