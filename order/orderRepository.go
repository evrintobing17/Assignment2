package order

type OrderRepository interface {
	Insert(interface{}) (interface{}, error)
	GetAll() (interface{}, error)
	Update(interface{}, map[string]interface{}) (interface{}, error)
	UpdateItem(req interface{}) (interface{}, error)
	Delete(interface{}) (interface{}, error)
	GetByID(interface{}) (interface{}, error)
}
