package model

type Items struct {
	ItemID      int    `gorm:"column:item_id;not null;AUTO_INCREMENT"`
	ItemCode    string `gorm:"column:item_code"`
	Description string `gorm:"column:description"`
	Quantity    string `gorm:"column:quantity"`
	OrderID     int    `gorm:"column:orders_id"`
}

type UpdateItems struct {
	ItemID      int    `gorm:"column:item_id;`
	ItemCode    string `gorm:"column:item_code"`
	Description string `gorm:"column:description"`
	Quantity    string `gorm:"column:quantity"`
	OrderID     int    `gorm:"column:orders_id"`
}

func (m *UpdateItems) TableName() string {
	return "orders.items"
}
