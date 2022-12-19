package ginrestaurant

import (
	"go-simple-service/common"
	"go-simple-service/component"
	"go-simple-service/modules/restaurant/restaurantbiz"
	"go-simple-service/modules/restaurant/restaurantmodel"
	"go-simple-service/modules/restaurant/restaurantstore"
	restaurantlikestore "go-simple-service/modules/restaurantlike/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter restaurantmodel.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())
		likeStore := restaurantlikestore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store, likeStore)

		result, err := biz.ListRestaurant(ctx.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
