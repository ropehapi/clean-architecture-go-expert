package service

import (
	"context"
	"github.com/ropehapi/clean-architecture-go-expert/internal/infra/grpc/pb"
	"github.com/ropehapi/clean-architecture-go-expert/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUsecase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrderUseCase usecase.ListOrderUsecase) *OrderService {
	return &OrderService{CreateOrderUseCase: createOrderUseCase, ListOrderUseCase: listOrderUseCase}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    req.Id,
		Price: float64(req.Price),
		Tax:   float64(req.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

//func (s *OrderService) ListOrder(ctx context.Context, in *pb.Blank) (*pb.OrderList, error) {
//	var list *pb.OrderList
//	output, err := s.ListOrderUseCase.Execute()
//	if err != nil {
//		return nil, err
//	}
//	for _, order := range output {
//
//	}
//}
