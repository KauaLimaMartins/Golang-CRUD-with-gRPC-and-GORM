package repositories

import (
	"github.com/kaualimamartins/gRPC-Go-Gorm/customers/models"
	"github.com/kaualimamartins/gRPC-Go-Gorm/database"
	"golang.org/x/crypto/bcrypt"
)

func List() ([]models.Customer, error) {
	var allCustomers []models.Customer

	result := database.DB.Find(&allCustomers)

	return allCustomers, result.Error
}

func FindById(customerId uint64) (models.Customer, error) {
	var customer models.Customer

	result := database.DB.Find(&customer, customerId)

	return customer, result.Error
}

func Create(customer *models.Customer) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	customer.Password = string(bytes)

	result := database.DB.Create(&customer)

	return result.Error
}

func Update(newCustomer *models.Customer) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(newCustomer.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	newCustomer.Password = string(bytes)

	result := database.DB.Updates(&newCustomer)

	return result.Error
}

func Delete(ids []uint64) error {
	result := database.DB.Delete(&models.Customer{}, ids)

	return result.Error
}
