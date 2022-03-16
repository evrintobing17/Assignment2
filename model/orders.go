package model

type Orders struct {
	OrderID      int     `gorm:"column:orders_id;not null;AUTO_INCREMENT" json:"order_id" example:"1"`
	CustomerName string  `gorm:"column:customer_name" json:"customer_name" example:"Evrin"`
	OrderedAt    string  `gorm:"column:ordered_at" json:"ordered_at" example:"2019-11-09T21:21:46Z"`
	Items        []Items `gorm:"foreignKey:OrderID;association_foreignkey:OrderID;constraint:OnDelete:CASCADE;" json:"items"`
}
