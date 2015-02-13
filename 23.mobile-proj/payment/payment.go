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

func Load(context appengine.Context, key *datastore.Key) (*Payment, error) {
	var loadedEntity = new(Payment)
	err := datastore.Get(context, key, loadedEntity);
	if err != nil {
		return nil, err
	}
	return loadedEntity, nil
}


