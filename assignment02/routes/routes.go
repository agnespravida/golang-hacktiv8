package routes

import (
	"assignment02/controllers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.POST("/orders", controllers.CreateOrder)

	// router.PUT("/cars/:carID", controllers.UpdateCar)
	router.GET("/orders/:orderID", controllers.GetOrderByID)

	router.GET("/orders", controllers.GetAllOrders)
	router.DELETE("/orders/:orderID", controllers.DeleteOrder)
	router.PUT("/orders/:orderID", controllers.UpdateOrder)

	// router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
