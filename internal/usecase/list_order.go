package usecase

import (
	"github.com/ropehapi/clean-architecture-go-expert/internal/entity"
	"github.com/ropehapi/clean-architecture-go-expert/pkg/events"
)

type ListOrderUsecase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUsecase(
	orderRepository entity.OrderRepositoryInterface,
	orderCreatedEvent events.EventInterface,
	eventDispatcher events.EventDispatcherInterface) *ListOrderUsecase {
	return &ListOrderUsecase{
		OrderRepository: orderRepository,
		OrderCreated:    orderCreatedEvent,
		EventDispatcher: eventDispatcher,
	}
}

func (uc *ListOrderUsecase) Execute() ([]OrderOutputDTO, error) {
	orders, err := uc.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	var ordersDTO []OrderOutputDTO
	for _, order := range orders {
		orderDTO := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		ordersDTO = append(ordersDTO, orderDTO)
	}

	return ordersDTO, nil
}
