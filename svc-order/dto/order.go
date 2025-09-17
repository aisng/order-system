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
