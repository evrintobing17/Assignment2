package order

type OrderUsecase interface {
	CreateOrder(interface{}) (interface{}, error)
	DeleteOrder(interface{}) (interface{}, error)
	GetOrderByID(interface{}) (interface{}, error)
	UpdateOrderByID(interface{}, map[string]interface{}) (interface{}, error)
	GetAllOrder() (interface{}, error)
}
