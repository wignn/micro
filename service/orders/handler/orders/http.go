package handler

import (
	"net/http"
	"github.com/wignn/micro/service/common/genproto/orders"
	"github.com/wignn/micro/service/common/util"
	"github.com/wignn/micro/service/orders/types"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersService(ordersService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		ordersService: ordersService,
	}

	return handler
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	util.WriteJSON(w, http.StatusOK, res)
}