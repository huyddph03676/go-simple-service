package main

import (
	"go-simple-service/component"
	"go-simple-service/middleware"
	"go-simple-service/modules/restaurant/restaurantmodel"
	"go-simple-service/modules/restaurant/restauranttransport/ginrestaurant"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured with env file. Err: %s", err)
	}

	dsn := os.Getenv("DB_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// Migrate the schema
	db.AutoMigrate(&restaurantmodel.RestaurantCreate{})

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	// => return cursor pointer
	appCtx := component.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	// CRUD
	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetOneRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run()
}
