package bootstrap

import (
	"by/user/container"
	"by/user/controller"
	"by/user/controller/shop"
	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine, app *container.Application) {
	//e.GET("/user", controller.User)
	//e.GET("/userinfo", controller.UserInfo)

	e.POST("/shop", shop.Add)
	e.DELETE("/shop/:id", shop.Del)
	e.PUT("/shop/:id", shop.Save)
	e.GET("/shop/:id", shop.GetById)
	e.GET("/shop", shop.GetListById)

	e.GET("/test/config", controller.Config)
	e.GET("/test/redis", controller.Redis)
	//e.GET("/", func(ctx *gin.Context) {
	//	db := container.Get("db").(*gorm.DB)
	//	defer container.Put("db", db)
	//	var user struct {
	//		Id uint
	//		Nickname string
	//	}
	//	db.Table("ex_users").First(&user, "id = 13")
	//	ctx.JSON(200, user)
	//})
	//r.Use(func(ctx *gin.Context) {
	//	log.Default().Println(ctx.ClientIP())
	//	log.Default().Println(ctx.Cookie("laravel_session"))
	//	log.Default().Println(ctx.DefaultQuery("b", "b"))
	//	log.Default().Println(ctx.GetQuery("a"))
	//	log.Default().Println(ctx.GetQuery("b"))
	//	log.Default().Println(ctx.GetHeader("User-Agent"))
	//
	//	ctx.String(200, "%s", "middleware")
	//})
	//infoGroup := r.Group("/info", func(ctx *gin.Context) {
	//	log.Default().Println(ctx.GetQuery("a"))
	//	log.Default().Println(ctx.GetQuery("b"))
	//	log.Default().Println(ctx.DefaultQuery("c", "c"))
	//
	//	ctx.String(200, "%s", "group")
	//})
	//{
	//	infoGroup.GET("/get", func(ctx *gin.Context) {
	//		ctx.String(200, "%s", "get")
	//	})
	//}
}
