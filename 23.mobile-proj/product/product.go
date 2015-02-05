package product

import (
	"appengine"
	"appengine/datastore"
)

type Product struct {
	PaymentKey *datastore.Key
	ContentKeys []*datastore.Key
	Value float64
}

func NewProduct () (*Product, error) {
	var createdProduct = new(Product)
	createdProduct.Value = 0.0
	return createdProduct, nil
}

func (this *Product) Save(context appengine.Context) (*datastore.Key, error) {
	key, err := datastore.Put(context, datastore.NewIncompleteKey(context, "Product", nil), this)
    if err != nil {
        return nil, err
    }
    
	return key, nil
}

func Load(context appengine.Context, encodedProductKey string) (*Product, error) {
	productKey, _ := datastore.DecodeKey(encodedProductKey)
	var loadedProduct = new(Product)
	err := datastore.Get(context, productKey, loadedProduct);
	if err != nil {
		return nil, err
	}
	return loadedProduct, nil
}