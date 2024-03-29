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

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = requester.GetUserId()

		// appCtx.GetMainDBConnection() return db cursor pointer => store = sqlStore{db: db}
		store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())

		// => store have "Create" method so that its interface is also { Create: func() }
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		data.GenUID(common.DbTypeRestaurant)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
