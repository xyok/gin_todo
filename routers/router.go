package routers

import (
	"gin_todo/controllers"
	"gin_todo/middleware"
	"gin_todo/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func  InitRouter() *gin.Engine{
	r := gin.Default()
	//r.Use(gin.Logger())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	r.POST("/login", controllers.LoginHandler)
	r.POST("/signup", controllers.SignUp)

	authorized := r.Group("/")
	authorized.Use(middleware.LoginRequired())
	{
		authorized.GET("/valide", controllers.ValidateToken)

		authorized.GET("/todo",controllers.AllTodo)
		authorized.POST("/todo",controllers.AddTodoHandler)

		authorized.GET("/todo/:id/",controllers.GetTodo)
		authorized.PUT("/todo/:id/",controllers.UpdateTodo)
		authorized.DELETE("/todo/:id/",controllers.DeleteTodo)
	}

	return r
}