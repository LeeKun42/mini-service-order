package controller

import (
	"github.com/kataras/iris/v12"
	"order/app/http/response"
	"order/app/service/order"
)

type orderController struct{}

var UserController = orderController{}

func (uc *orderController) Info(context iris.Context) {
	orderId, err := context.Params().GetInt("id")
	if err != nil {
		response.Fail(context, "参数错误")
	}
	orderInfo, err := order.Service.Get(orderId)
	response.Success(context, iris.Map{"message": "this is order service!", "order info": orderInfo})
}
