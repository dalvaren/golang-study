package account

import (
	"time"
	"errors"
	"crypto/md5"
	"encoding/hex"

	"product"
	"payment"
	// "content"

	"appengine"
	"appengine/datastore"
)

type Account struct {
	Id int
	Name  string
	Email string
	CreatedAt time.Time
}

func CreateCustomer(context appengine.Context ,name string, email string) (*Account, *product.Product, *payment.Payment, error){
	appId := GetMD5Hash(time.Now().String())
	account, err := NewAccount(context, name, email)
	if err != nil { return nil, nil, nil, err }
	product, err := product.NewProduct(appId, account.Id)
	if err != nil { return nil, nil, nil, err }
	payment, err := payment.NewPayment(appId)
	if err != nil { return nil, nil, nil, err }

	return account, product, payment, nil
}

func NewAccount (context appengine.Context, name string, email string) (*Account, error) {
	if len(name) == 0 || len(email) == 0 {
		return nil, errors.New("Empty Name or Email")
	}

	query := datastore.NewQuery("Account")
	accountId, _ := query.Count(context)
	accountId = accountId + 1

	var createdAccount = new(Account)
	createdAccount.Id = accountId
	createdAccount.Name = name
	createdAccount.Email = email
	createdAccount.CreatedAt = time.Now()
	return createdAccount, nil
}

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
