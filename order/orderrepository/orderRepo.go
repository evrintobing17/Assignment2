package orderrepository

import (
	"errors"
	"strconv"

	"github.com/evrintobing17/model"
	"github.com/evrintobing17/order"
	"github.com/jinzhu/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) order.OrderRepository {
	return &OrderRepo{
		db: db,
	}

}

func (o *OrderRepo) Insert(req interface{}) (interface{}, error) {
	order := req.(*model.Orders)
	db := o.db.Create(&order)
	if db.Error != nil {
		return nil, db.Error
	}
	return order, nil
}

func (o *OrderRepo) Delete(req interface{}) (interface{}, error) {
	var order model.Orders
	var items model.Items

	orderId := req.(string)
	orderIds, _ := strconv.Atoi(orderId)
	_ = o.db.Delete(&items, "orders_id=?", orderIds)
	db := o.db.Delete(&order, "orders_id=?", orderIds)
	if db.Error != nil {
		return nil, db.Error
	}

	return &order, nil
}

func (o *OrderRepo) GetAll() (interface{}, error) {
	var order []model.Orders

	db := o.db.Preload("Items").Find(&order)

	if db.Error != nil {
		return nil, db.Error
	}
	return &order, nil
}

func (o *OrderRepo) GetByID(req interface{}) (interface{}, error) {
	var order model.Orders
	orderId := req.(int)
	db := o.db.First(&order, "order_id=?", orderId)

	if db.Error != nil {
		return nil, db.Error
	}

	return &order, nil
}

func (o *OrderRepo) Update(req interface{}, updatedata map[string]interface{}) (interface{}, error) {
	if req == nil {
		return nil, errors.New("order id can't empty")
	}
	orders := req.(int)
	// orderId, _ := strconv.Atoi(orders)
	var order model.Orders

	db := o.db.First(&order, "orders_id=?", orders)
	if db.Error != nil {
		return nil, db.Error
	}
	db = o.db.Model(&order).Update(updatedata)
	if db.Error != nil {
		return nil, db.Error
	}
	return &order, nil
}

func (o *OrderRepo) UpdateItem(req interface{}) (interface{}, error) {
	if req == nil {
		return nil, errors.New("order id can't empty")
	}
	updateData := req.(model.Items)

	var itemModel model.Items

	itemModel.ItemCode = updateData.ItemCode
	itemModel.Description = updateData.Description
	itemModel.Quantity = updateData.Quantity

	db := o.db.Model(itemModel).Where("item_id = ?", updateData.ItemID).Updates(model.Items{
		ItemCode:    updateData.ItemCode,
		Description: updateData.Description,
		Quantity:    updateData.Quantity,
	})

	if db.Error != nil {
		return nil, db.Error
	}
	return &updateData, nil
}
