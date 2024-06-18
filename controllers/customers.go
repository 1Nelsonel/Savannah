package controllers

import (
	"net/http"

	"github.com/1Nelsonel/Savannah/database"
	"github.com/1Nelsonel/Savannah/models"
	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	db := database.DBConn
	// list all customers here
	var customers []models.Customer

	if err := db.Find(&customers).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "No customers found",
		})
		return 
	}

	context := gin.H{
		"customers": customers,
	}
	
	c.JSON(http.StatusOK, context)
}

func CreateCustomer(c *gin.Context) {
	db := database.DBConn

	customer := new(models.Customer)

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	
	
	// create a customer here
	if err := db.Create(&customer).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Unable to create a customer",
		})
		return
	}

	context := gin.H{
		"message": "Create a customer",
		"customer": customer,
	}

	c.JSON(http.StatusOK, context)
	
}

func UpdateCustomer(c *gin.Context) {
	db := database.DBConn

	// update a customer here
	id := c.Param("id")

	// find customer by id
	var customer models.Customer
	if err := db.First(&customer, id).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Customer not found",
		})
		return
	}

	// bind the request data to the customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// save the updated customer
	if err := db.Save(&customer).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Unable to update a customer",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Update a customer",
	})
}

func DeleteCustomer(c *gin.Context) {
	db := database.DBConn

	// delete a customer here
	id := c.Param("id")

	// find customer by id
	var customer models.Customer
	if err := db.First(&customer, id).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Customer not found",
		})
		return
	}

	// delete the customer
	if err := db.Delete(&customer).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Unable to delete a customer",
		})
		return
	}


	c.JSON(200, gin.H{
		"message": "Delete a customer",
	})
}