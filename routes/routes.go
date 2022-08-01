package routes

import (
	"golang_cms/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func BannerRoute() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/banner", controllers.CreateBanner)
	r.GET("/banner/:bannerId", controllers.GetABanner)
	r.PUT("/banner/:bannerId", controllers.EditABanner)
	r.DELETE("/banner/:bannerId", controllers.DeleteBanner)
	r.GET("/banners", controllers.GetAllBanner)
	r.Run(":9999")
	return r
}
