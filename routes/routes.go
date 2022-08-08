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
	//login user & generate token
	r.POST("/signup", controllers.UsersignUp)
	r.GET("/login")
	//banner
	r.POST("/banner", controllers.CreateBanner)
	r.GET("/banner/:bannerId", controllers.GetABanner)
	r.PUT("/banner/:bannerId", controllers.EditABanner)
	r.DELETE("/banner/:bannerId", controllers.DeleteBanner)
	r.GET("/banners", controllers.GetAllBanner)
	//meta
	r.POST("/meta", controllers.Createmeta)
	r.GET("/meta/:metaId", controllers.GetAmeta)
	r.PUT("/meta/:metaId", controllers.EditAmeta)
	r.DELETE("/meta/:metaId", controllers.Deletemeta)
	r.GET("/metas", controllers.GetAllmeta)
	r.Run(":8888")
	return r
}
