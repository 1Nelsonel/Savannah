package routes

import (
	"github.com/1Nelsonel/Savannah/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine) {
	// customers
	customer := app.Group("/customers")
	customer.GET("/", controllers.GetCustomers)
	customer.POST("/create/", controllers.CreateCustomer)
	customer.PUT("/update/:id/", controllers.UpdateCustomer)
	customer.DELETE("/delete/:id/", controllers.DeleteCustomer)

	// orders
	order := app.Group("/orders")
	order.GET("/", controllers.GetOrders)
	order.POST("/create", controllers.CreateOrder)
	order.PUT("/update/:id", controllers.UpdateOrder)
	order.DELETE("/delete/:id",controllers. DeleteOrder)
}