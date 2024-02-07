package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/xALEGORx/go-expression-calculator/docs"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/handler"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/middlewares"
	"net/http"
)

func InitRouter(router *gin.Engine) *gin.Engine {
	v1 := router.Group("/api/v1")
	v1.Use(middlewares.CORS())
	{
		// task
		{
			task := &handler.Task{Route: v1.Group("/task")}
			task.Route.GET("", task.Index)
			task.Route.POST("", task.Store)
			task.Route.GET("/:id", task.Show)

			// why? idk
			task.Route.OPTIONS("", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	return router
}
