package router

import (
	"github.com/basic-go/internal/book"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *book.Handler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		books := api.Group("/books")
		{
			books.POST("", handler.Create)
			books.GET("", handler.FindAll)
			books.GET("/:id", handler.FindByID)
			books.PUT("/:id", handler.Update)
			books.DELETE("/:id", handler.Delete)
		}
	}

	return r
}
