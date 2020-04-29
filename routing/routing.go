package routing

import (
	"playground/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// SetupRouter Creates all routing configurations
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS Configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200", "http://localhost:4545"}
	r.Use(cors.New(config))

	// Swagger Configuration
	c := controllers.NewController()

	v1 := r.Group("/api/v1")
	{
		examples := v1.Group("/examples")
		{
			examples.GET("ping", c.PingExample)
		}
		lyrics := v1.Group("/lyrics")
		{
			lyrics.GET(":artist/:song", c.GetLyrics)
		}
	}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
