package customers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/kaualimamartins/gRPC-Go-Gorm/customers/helpers"
	pbCustomer "github.com/kaualimamartins/gRPC-Go-Gorm/customers/proto/pb"
	"github.com/kaualimamartins/gRPC-Go-Gorm/customers/repositories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pbCustomer.UnimplementedCustomerServiceServer
}

func NewCustomerServer() *Server {
	return &Server{}
}

func (s *Server) GetAllCustomers(ctx context.Context, req *pbCustomer.GetAllCustomersReq) (*pbCustomer.GetAllCustomersRes, error) {
	customers, err := repositories.List()

	if err != nil {
		return nil, status.Errorf(
			codes.DataLoss,
			err.Error(),
		)
	}

	convertedCustomers := make([]*pbCustomer.Customer, len(customers))

	for i, r := range customers {
		convertedCustomers[i] = helpers.StructToResData(r)
	}

	return &pbCustomer.GetAllCustomersRes{
		Customers: convertedCustomers,
	}, nil
}

func (s *Server) GetCustomerById(ctx context.Context, req *pbCustomer.GetCustomerByIdReq) (*pbCustomer.GetCustomerByIdRes, error) {
	if len(req.GetCustomerId()) == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"CustomerId is required",
		)
	}

	customerId, err := strconv.ParseUint(req.GetCustomerId(), 10, 32)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			err.Error(),
		)
	}

	customer, err := repositories.FindById(customerId)

	if err != nil {
		return nil, status.Errorf(
			codes.DataLoss,
			err.Error(),
		)
	}

	if customer.ID == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"customer not found",
		)
	}

	convertedCustomer := helpers.StructToResData(customer)

	return &pbCustomer.GetCustomerByIdRes{
		Customer: convertedCustomer,
	}, nil
}

func (s *Server) CreateCustomer(ctx context.Context, req *pbCustomer.CreateCustomerReq) (*pbCustomer.CreateCustomerRes, error) {
	data, err := helpers.ReqDataToStruct(req.GetCustomer())

	if err != nil {
		return nil, status.Errorf(
			codes.DataLoss,
			err.Error(),
		)
	}

	err = repositories.Create(&data)

	if err != nil {
		return nil, status.Errorf(
			codes.DataLoss,
			err.Error(),
		)
	}

	return &pbCustomer.CreateCustomerRes{
		Customer: helpers.StructToResData(data),
		Status:   http.StatusOK,
	}, nil
}

func (s *Server) UpdateCustomer(ctx context.Context, req *pbCustomer.UpdateCustomerReq) (*pbCustomer.UpdateCustomerRes, error) {
	data, err := helpers.ReqDataToStruct(req.GetCustomer())

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			err.Error(),
		)
	}

	currentCustomer, err := repositories.FindById(data.ID)

	if err != nil {
		return nil, status.Errorf(
			codes.DataLoss,
			err.Error(),
		)
	}

	if currentCustomer.ID == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"customer not found",
		)
	}

	err = repositories.Update(&data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			err.Error(),
		)
	}

	return &pbCustomer.UpdateCustomerRes{
		Status:   http.StatusOK,
		Customer: helpers.StructToResData(data),
	}, nil
}

func (s *Server) DeleteCustomer(ctx context.Context, req *pbCustomer.DeleteCustomerReq) (*pbCustomer.DeleteCustomerRes, error) {
	ids := req.GetCustomersIds()

	convertedIds := make([]uint64, len(ids))

	for i, r := range ids {
		id, err := strconv.ParseUint(r, 10, 32)

		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				err.Error(),
			)
		}

		convertedIds[i] = id
	}

	err := repositories.Delete(convertedIds)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			err.Error(),
		)
	}

	return &pbCustomer.DeleteCustomerRes{
		Status: http.StatusOK,
	}, nil
}
