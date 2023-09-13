package service

import (
	"context"

	"github.com/caiocp/clean-arch-go/internal/infra/grpc/pb"
	"github.com/caiocp/clean-arch-go/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, request *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    request.Id,
		Price: float64(request.Price),
		Tax:   float64(request.Tax),
	}

	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	response := &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}

	return response, nil
}
