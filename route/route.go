package route

import (
	"auth-services/controller"
	"auth-services/middleware"
	"auth-services/repository"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// SetupRoutes : all the routes are defined here
func SetupRoutes() {
	httpRouter := gin.Default()

	httpRouter.Use(middleware.CORSMiddleware())
	

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	customerRepo := repository.NewCustomerRepo()
	customerController := controller.NewCustomerController(customerRepo)
	apiRoutes := httpRouter.Group("/api")
	{
		apiRoutes.POST("/signin", customerController.Login)
	}

	

	httpRouter.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	// httpRouter.Run(":8089")
}
