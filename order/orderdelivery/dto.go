package orderdelivery

type Orders struct {
	CustomerName string  `json:"customerName" example:"George"`
	OrderedAt    string  `json:"orderedAt" example:"2019-11-09T21:21:46Z"`
	Items        []Items `json:"items"`
}

type InsertOrders struct {
	CustomerName string  `json:"customerName"`
	OrderedAt    string  `json:"orderedAt"`
	Items        []Items `json:"items"`
}

type Items struct {
	ItemCode    string `json:"itemCode" example:"123"`
	Description string `json:"description" example:"Car"`
	Quantity    string `json:"quantity" example:"2"`
}

type GetAllOrderResponse []Orders

type UpdateOrder struct {
	CustomerName string        `json:"customerName" example:"George"`
	OrderedAt    string        `json:"orderedAt" example:"2019-11-09T21:21:46Z"`
	Items        []UpdateItems `json:"items"`
}

type UpdateItems struct {
	ItemID      string `json:"lineItemId" example:"12"`
	ItemCode    string `json:"itemCode" example:"123"`
	Description string `json:"description" example:"Car"`
	Quantity    string `json:"quantity" example:"2"`
}
