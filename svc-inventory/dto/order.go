package dto

type Order struct {
	ID           string `json:"id"`
	BuyerAddress string `json:"buyer_address"`
	ItemID       string `json:"item_id"`
	Status       string `json:"status"`
}

type OrderCreatedEvent struct {
	OrderID      string `json:"order_id"`
	BuyerAddress string `json:"buyer_address"`
	ItemID       string `json:"item_id"`
}

type OrderUpdatedEvent struct {
	OrderID    string `json:"id"`
	StatusFrom string `json:"status_from"`
	StatusTo   string `json:"status_to"`
}
