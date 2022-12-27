package helpers

import (
	"strconv"

	"github.com/kaualimamartins/gRPC-Go-Gorm/customers/models"
	pbCustomer "github.com/kaualimamartins/gRPC-Go-Gorm/customers/proto/pb"
)

func ReqDataToStruct(reqData *pbCustomer.Customer) (models.Customer, error) {
	id, err := strconv.ParseUint(reqData.Id, 10, 32)

	if err != nil {
		id = 0
	}

	customer := models.Customer{
		ID:        id,
		FirstName: reqData.FirstName,
		LastName:  reqData.LastName,
		Email:     reqData.Email,
		Password:  reqData.Password,
	}

	return customer, nil
}

func StructToResData(structData models.Customer) *pbCustomer.Customer {
	customer := &pbCustomer.Customer{
		Id:        strconv.FormatUint(uint64(structData.ID), 10),
		FirstName: structData.FirstName,
		LastName:  structData.LastName,
		Email:     structData.Email,
		Password:  "",
		CreatedAt: structData.CreatedAt.String(),
		UpdatedAt: structData.UpdatedAt.String(),
	}

	return customer
}
