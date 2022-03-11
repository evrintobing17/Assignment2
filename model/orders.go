package model

type Orders struct {
	OrderID      int     `gorm:"column:orders_id;not null;AUTO_INCREMENT"`
	CustomerName string  `gorm:"column:customer_name"`
	OrderedAt    string  `gorm:"column:ordered_at"`
	Items        []Items `gorm:"foreignKey:OrderID;association_foreignkey:OrderID;constraint:OnDelete:CASCADE;"`
}
