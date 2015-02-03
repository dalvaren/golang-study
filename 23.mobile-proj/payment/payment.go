package payment

import (
	"time"
)

type Payment struct {
	ProductId string
	CreatedAt time.Time
	ValidUntil time.Time
}

func NewPayment (productId string) (*Payment, error) {
	var createdPayment = new(Payment)
	createdPayment.ProductId = productId
	createdPayment.CreatedAt = time.Now()
	createdPayment.ValidUntil = time.Now()
	return createdPayment, nil
}