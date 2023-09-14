package service

import (
	"context"

	"github.com/caiocp/clean-arch-go/internal/infra/grpc/pb"
	"github.com/caiocp/clean-arch-go/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	FindOrdersUseCase  usecase.FindOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, findOrdersUseCase usecase.FindOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		FindOrdersUseCase:  findOrdersUseCase,
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

func (s *OrderService) ListOrders(ctx context.Context, request *pb.Blank) (*pb.OrdersList, error) {
	output, err := s.FindOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	orders := []*pb.Order{}

	for _, order := range output {
		orders = append(orders, &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}

	response := &pb.OrdersList{
		Orders: orders,
	}

	return response, nil
}
