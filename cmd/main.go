package main

import (
	"book-api/cmd/docs"
	"book-api/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Book API
// @version 1.0
// @description This is a sample API for managing books.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func SetupRouter() *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		bookRoutes := v1.Group("/books")
		{
			bookRoutes.GET("/:id", handlers.GetBookByID)
			bookRoutes.GET("", handlers.GetAllBooks)
			bookRoutes.POST("", handlers.CreateBook)
			bookRoutes.PUT("/:id", handlers.UpdateBook)
			bookRoutes.DELETE("/:id", handlers.DeleteBook)
		}

		authorsRoutes := v1.Group("/authors")
		{
			authorsRoutes.GET("/:id", handlers.GetAuthorByID)
			authorsRoutes.GET("", handlers.GetAllAuthors)
			authorsRoutes.POST("", handlers.CreateAuthor)
			authorsRoutes.PUT("/:id", handlers.UpdateAuthor)
			authorsRoutes.DELETE("/:id", handlers.DeleteAuthor)
		}
	}

	return router
}

func main() {
	router := SetupRouter()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run()
}
