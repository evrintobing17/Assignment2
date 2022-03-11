package orderusecase

import (
	"fmt"
	"strconv"

	"github.com/evrintobing17/model"
	"github.com/evrintobing17/order"
	"github.com/evrintobing17/order/orderdelivery"
	"github.com/mitchellh/mapstructure"
)

type OrderUC struct {
	orderRepo order.OrderRepository
}

type Orders struct {
	CustomerName string
	OrderedAt    string
	Items        []Items
}

type Items struct {
	ItemID      int
	ItemCode    string
	Description string
	Quantity    string
	OrderID     int
}

type UpdateItems struct {
	ItemID      string `mapstructure:"lineItemId"`
	ItemCode    string `mapstructure:"itemCode"`
	Description string `mapstructure:"description"`
	Quantity    string `mapstructure:"quantity"`
	OrderID     int    `mapstructure:"orders_id"`
}

func NewOrderUsecase(orderRepo order.OrderRepository) order.OrderUsecase {
	return &OrderUC{
		orderRepo: orderRepo,
	}
}

func (o *OrderUC) CreateOrder(req interface{}) (interface{}, error) {
	order := req.(orderdelivery.InsertOrders)
	var itemData []model.Items

	orders := &model.Orders{
		CustomerName: order.CustomerName,
		OrderedAt:    order.OrderedAt,
	}
	for _, item := range order.Items {
		itemData = append(itemData, model.Items{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		})

	}
	orders.Items = itemData

	orderData, err := o.orderRepo.Insert(orders)
	if err != nil {
		return nil, err
	}
	return orderData, nil
}

func (o *OrderUC) DeleteOrder(req interface{}) (interface{}, error) {
	orders := req.(string)

	orderData, err := o.orderRepo.Delete(orders)
	if err != nil {
		return nil, err
	}

	return orderData, nil
}

func (o *OrderUC) GetAllOrder() (interface{}, error) {
	orderData, err := o.orderRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return orderData, nil
}

func (o *OrderUC) GetOrderByID(req interface{}) (interface{}, error) {
	orders := req.(int)

	orderData, err := o.orderRepo.GetByID(orders)
	if err != nil {
		return nil, err
	}

	return orderData, nil
}

func (o *OrderUC) UpdateOrderByID(req interface{}, updateData map[string]interface{}) (interface{}, error) {
	id := req.(string)
	orders, _ := strconv.Atoi(id)
	var request []UpdateItems
	err := mapstructure.Decode(updateData["items"], &request)
	if err != nil {
		return nil, err
	}

	fmt.Println(request,"data pasti")
	fmt.Println(request[0],"data pasti")
	// fmt.Println(request[1],"data pasti")
	for _, item := range request {
		fmt.Println(item,"data jelas")
		itemID, _ := strconv.Atoi(item.ItemID)
		items := model.Items{
			ItemID:      itemID,
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderID:     orders,
		}
		_, err := o.orderRepo.UpdateItem(items)
		if err != nil {
			return nil, err
		}
	}

	orderData, err := o.orderRepo.Update(orders, updateData)
	if err != nil {
		return nil, err
	}

	return orderData, nil
}
