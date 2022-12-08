package main

import (
	"go-simple-service/component"
	"go-simple-service/modules/restaurant/restauranttransport/ginrestaurant"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "200lab:200lab@tcp(127.0.0.1:3306)/200lab?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// Migrate the schema
	db.AutoMigrate(&Restaurant{})

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// CRUD
	// => return cursor pointer
	appCtx := component.NewAppContext(db)

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))

		restaurants.GET("/:id", ginrestaurant.GetOneRestaurant(appCtx))

		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))

		restaurants.PATCH("/:id", func(ctx *gin.Context) {
			id, err := strconv.Atoi(ctx.Param("id"))

			if err != nil {
				ctx.JSON(401, gin.H{
					"error": err.Error(),
				})

				return
			}

			var data RestaurantUpdate

			if err := ctx.ShouldBind(&data); err != nil {
				ctx.JSON(401, gin.H{
					"error": err.Error(),
				})

				return
			}

			if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
				ctx.JSON(401, gin.H{
					"error": err.Error(),
				})

				return
			}

			ctx.JSON(http.StatusOK, gin.H{"ok": 1})
		})

		restaurants.DELETE("/:id", func(ctx *gin.Context) {
			id, err := strconv.Atoi(ctx.Param("id"))

			if err != nil {
				ctx.JSON(401, gin.H{
					"error": err.Error(),
				})

				return
			}

			if err := db.Where("id = ?", id).Delete(&Restaurant{}).Error; err != nil {
				ctx.JSON(401, gin.H{
					"error": err.Error(),
				})

				return
			}

			ctx.JSON(http.StatusOK, gin.H{"ok": 1})
		})

	}

	return r.Run()
}

type Restaurant struct {
	Id      int       `json:"id" gorm:"column:id;"`
	Name    *string   `json:"name" gorm:"column:name;"`
	Addr    *string   `json:"addr" gorm:"column:addr;"`
	Created time.Time `json:"created_at" gorm:"autoCreateTime; column:created_at"`
	Updated time.Time `json:"updated_at" gorm:"autoCreateTime; column:updated_at"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name    *string    `json:"name" gorm:"column:name"`
	Addr    *string    `json:"addr" gorm:"column:addr"`
	Updated *time.Time `gorm:"autoUpdateTime;column:updated_at"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
