package types

type CartItem struct {
	ProductID int `json:"productId" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type CheckoutPayload struct {
	Items []CartItem `json:"items" validate:"required"`
}
