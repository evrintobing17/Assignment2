package model

type Items struct {
	ItemID      int    `gorm:"column:item_id;not null;AUTO_INCREMENT" json:"item_id" example:"1"`
	ItemCode    string `gorm:"column:item_code" json:"item_code" example:"123"`
	Description string `gorm:"column:description" json:"description" example:"Iphone X"`
	Quantity    string `gorm:"column:quantity" json:"quantity" example:"5"`
	OrderID     int    `gorm:"column:orders_id" json:"order_id" example:"1"`
}

type UpdateItems struct {
	ItemID      int    `gorm:"column:item_id"`
	ItemCode    string `gorm:"column:item_code"`
	Description string `gorm:"column:description"`
	Quantity    string `gorm:"column:quantity"`
	OrderID     int    `gorm:"column:orders_id"`
}

func (m *UpdateItems) TableName() string {
	return "orders.items"
}
