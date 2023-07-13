package order

import (
	"github.com/kataras/iris/v12/x/errors"
	"gorm.io/gorm"
	"iris-app/app/model"
	"iris-app/app/model/dto"
	"iris-app/app/rpc/client"
)

type orderService struct{}

var Service = orderService{}

const (
	StatusNew int = 0
	StatusEnd int = 10
)

func (us *orderService) Get(id int) (dto.Order, error) {
	var order dto.Order
	result := model.Instance().Find(&order, id)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return order, errors.New("订单不存在")
	}
	//调用用户模块查询用户信息
	err := client.NewUserClient().Get(order.UserID, &order.User)
	if err != nil {
		return order, errors.New("用户不存在")
	}
	return order, nil
}
