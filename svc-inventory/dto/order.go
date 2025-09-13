package dto

type Order struct {
	ID           string  `json:"id"`
	BuyerAddress string  `json:"buyer_address"`
	Item         string  `json:"item"`
	Quantity     int     `json:"quantity"`
	Amount       float64 `json:"amount"`
	Status       string  `json:"status"`
}

type OrderCreatedEvent struct {
	OrderID      string  `json:"order_id"`
	BuyerAddress string  `json:"buyer_address"`
	Item         string  `json:"item"`
	Quantity     int     `json:"quantity"`
	Amount       float64 `json:"amount"`
}
