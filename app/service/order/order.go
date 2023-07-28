package order

import (
	"github.com/kataras/iris/v12/x/errors"
	"gorm.io/gorm"
	"order/app/model/dto"
	"order/app/rpc/client"
)

type Service struct {
	DB *gorm.DB
}

func NewService() *Service {
	return &Service{}
}

const (
	StatusNew int = 0
	StatusEnd int = 10
)

func (service *Service) Get(id int) (dto.Order, error) {
	var order dto.Order
	result := service.DB.Find(&order, id)
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
