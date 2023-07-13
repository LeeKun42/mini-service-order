// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dto

import (
	"time"
)

const TableNameOrder = "orders"

// Order mapped from table <orders>
type Order struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID    int       `gorm:"column:user_id;not null;comment:用户id" json:"user_id"`
	Status    int       `gorm:"column:status;not null;comment:订单状态 0：生成状态 10：完成状态" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	User 	  User	`json:"user"`
}

// TableName Order's table name
func (*Order) TableName() string {
	return TableNameOrder
}
