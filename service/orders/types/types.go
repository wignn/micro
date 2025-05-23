package types

import (
	"context"

	"github.com/wignn/micro/service/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context,  *orders.Order) error
}