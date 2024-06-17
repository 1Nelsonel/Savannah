package helpers

import (
	"strconv"

	"github.com/1Nelsonel/Savannah/models"
	"gorm.io/gorm"
)

// validateCustomerID checks if the customer exists in the database.
func ValidateCustomerID(db *gorm.DB, customerID uint) error {
	var customer models.Customer
	if err := db.First(&customer, customerID).Error; err != nil {
		return err
	}
	return nil
}

func StringToUint(s string) (uint, error) {
	id, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}