package routes

import (
	"goer/docs"
	"goer/global"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// MapSwagRoutes
// @title                       API Docs
// @version                     1.0
// @BasePath                    /
// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
func MapSwagRoutes(r *gin.Engine) {
	// swagger info
	docs.SwaggerInfo.Title = "API Docs"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = global.Config.App.ApiUrl
	// docs.SwaggerInfo.BasePath = "/"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// use ginSwagger middleware to serve the API docs
	// DefaultModelsExpandDepth: set -1 to hide models below
	r.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))
}
