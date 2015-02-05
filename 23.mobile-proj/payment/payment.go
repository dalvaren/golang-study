package payment

import (
	"time"
	"appengine"
	"appengine/datastore"
)

type Payment struct {
	CreatedAt time.Time
	ValidUntil time.Time
}

func NewPayment () (*Payment, error) {
	var createdPayment = new(Payment)
	createdPayment.CreatedAt = time.Now()
	createdPayment.ValidUntil = time.Now()
	return createdPayment, nil
}

func (this *Payment) Save(context appengine.Context) (*datastore.Key, error) {
	key, err := datastore.Put(context, datastore.NewIncompleteKey(context, "Payment", nil), this)
    if err != nil {
        return nil, err
    }
    
	return key, nil
}