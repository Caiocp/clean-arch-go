package web

import (
	"encoding/json"
	"net/http"

	"github.com/caiocp/clean-arch-go/internal/entity"
	"github.com/caiocp/clean-arch-go/internal/usecase"
	"github.com/caiocp/clean-arch-go/pkg/events"
)

type WebOrderHandler struct {
	OrderRepository   entity.OrderRepositoryInterface
	EventDispatcher   events.EventDispatcherInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	orderRepository entity.OrderRepositoryInterface,
	eventDispatcher events.EventDispatcherInterface,
	orderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		OrderRepository:   orderRepository,
		EventDispatcher:   eventDispatcher,
		OrderCreatedEvent: orderCreatedEvent,
	}
}

func (h *WebOrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var input usecase.OrderInputDTO

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrderUseCase := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrderUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	listOrdersUseCase := usecase.NewListOrdersUseCase(h.OrderRepository)
	output, err := listOrdersUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
