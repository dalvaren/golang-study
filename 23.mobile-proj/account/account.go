package account

import (
	"time"
	"errors"

	"product"
	"payment"
	// "content"

	"appengine"
	"appengine/datastore"
)

type Account struct {
	Name  string
	Email string
	CreatedAt time.Time
	ProductKeys []*datastore.Key
}

func NewAccount (name string, email string) (*Account, error) {
	if len(name) == 0 || len(email) == 0 {
		return nil, errors.New("Empty Name or Email")
	}

	var createdAccount = new(Account)
	createdAccount.Name = name
	createdAccount.Email = email
	createdAccount.CreatedAt = time.Now()
	return createdAccount, nil
}

func (this *Account) Save(context appengine.Context) (*datastore.Key, error) {
	key, err := datastore.Put(context, datastore.NewIncompleteKey(context, "Account", nil), this)
    if err != nil {
        return nil, err
    }
    
	return key, nil
}

func Load(context appengine.Context, encodedKey string) (*Account, error) {
	key, _ := datastore.DecodeKey(encodedKey)
	var loadedEntity = new(Account)
	err := datastore.Get(context, key, loadedEntity);
	if err != nil {
		return nil, err
	}
	return loadedEntity, nil
}

func CheckExistentAccount(context appengine.Context, email string) bool {
	queryAccount := datastore.NewQuery("Account").
		Filter("Email =", email)
	var accounts []Account
	_, _ = queryAccount.GetAll(context, &accounts)
	if len(accounts) > 0 {
		return true
	}
	return false
}

func CreateAccount(context appengine.Context ,name string, email string) (*Account, *product.Product, *payment.Payment, error){

	if existentAccount := CheckExistentAccount(context, email); existentAccount == true {
		return nil, nil, nil, errors.New("An user with this email already exists")
	}

	payment, err := payment.NewPayment()
	if err != nil { return nil, nil, nil, err }
	paymentKey, err := payment.Save(context)
	if err != nil { return nil, nil, nil, err }

	product, err := product.NewProduct()
	if err != nil { return nil, nil, nil, err }
	product.PaymentKey = paymentKey
	productKey, err := product.Save(context)
	if err != nil { return nil, nil, nil, err }

	account, err := NewAccount(name, email)
	if err != nil { return nil, nil, nil, err }
	account.ProductKeys = append(account.ProductKeys, productKey)
	_, err = account.Save(context)
	if err != nil { return nil, nil, nil, err }

	return account, product, payment, nil
}
