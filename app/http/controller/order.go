package controller

import (
	"github.com/kataras/iris/v12"
	"order/app/http/response"
	"order/app/service/order"
)

type OrderController struct {
	Service *order.Service
}

func NewOrderController() *OrderController {
	return &OrderController{
		Service: order.NewService(),
	}
}

func (oc *OrderController) Info(context iris.Context) {
	orderId, err := context.Params().GetInt("id")
	if err != nil {
		response.Fail(context, "参数错误")
	}
	orderInfo, err := oc.Service.Get(orderId)
	response.Success(context, iris.Map{"message": "this is order service!", "order info": orderInfo})
}
