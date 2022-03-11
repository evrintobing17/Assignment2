package orderdelivery

type Orders struct {
	CustomerName string  `json:"customerName"`
	OrderedAt    string  `json:"orderedAt"`
	Items        []Items `json:"items"`
}

type InsertOrders struct {
	CustomerName string  `json:"customerName"`
	OrderedAt    string  `json:"orderedAt"`
	Items        []Items `json:"items"`
}

type Items struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
}

type GetAllOrderResponse []Orders

type UpdateOrder struct {
	CustomerName string        `json:"customerName"`
	OrderedAt    string        `json:"orderedAt"`
	Items        []UpdateItems `json:"items"`
}

type UpdateItems struct {
	ItemID      string `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
}
