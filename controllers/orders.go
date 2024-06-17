package controllers

import (
	"log"
	"net/http"

	"github.com/1Nelsonel/Savannah/database"
	"github.com/1Nelsonel/Savannah/models"
	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	db := database.DBConn

	// list all orders here
	orders := []models.Order{}

	if err := db.Find(&orders).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "No orders found",
		})
		return
	}

	context := gin.H{
		"orders": orders,
	}
	
	c.JSON(http.StatusOK, context)
}

// CreateOrder handles the creation of a new order.
func CreateOrder(c *gin.Context) {
	db := database.DBConn

	var order models.Order

	// Bind JSON input to the order struct
	if err := c.ShouldBindJSON(&order); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create the order
	if err := db.Create(&order).Error; err != nil {
		log.Printf("Error creating order: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}

func UpdateOrder(c *gin.Context) {
	db := database.DBConn

	// update an order here
	id := c.Param("id")

	// find order by id
	var order models.Order
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Order not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Save(&order).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Unable to update an order",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Update an order",
	})
}

func DeleteOrder(c *gin.Context) {
	db := database.DBConn

	// delete an order here
	id := c.Param("id")

	if err := db.Delete(&models.Order{}, id).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Unable to delete an order",
		})
		return
	}

	if err := db.Delete(&models.Order{}, id).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Unable to delete an order",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Delete an order",
	})
}