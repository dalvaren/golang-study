package mobile

import (
	"fmt"
	"net/http"

	"account"
	"product"
	"payment"
	// "content"

	"appengine"
	"appengine/datastore"
)

func init() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/test", handleTwo)
	http.HandleFunc("/customer", handleListCustomers)
	http.HandleFunc("/customer/add", handleAddNewCustomer)
	http.HandleFunc("/delete", handleRemoveAll)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "GET requests only", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "<html><body>Hello, World! Dude! 세상아 안녕!</body></html>")
}

func handleTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body>lalalão!</body></html>")
}

func handleListCustomers(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)
	query := datastore.NewQuery("Account")
	var customers []account.Account
    if _, err := query.GetAll(context, &customers); err == nil {
    	fmt.Fprint(w, "<div>Customers:</div>")
    	for index, customer := range customers {
    		fmt.Fprintf(w, "<div>%d. ---</div>", index)
    		fmt.Fprintf(w, "<div>%+v</div>", customer)

    		queryProducts := datastore.NewQuery("Product").Filter("AccountId =", customer.Id)
    		var products []product.Product
    		queryProducts.GetAll(context, &products);
    		for _, product := range products {
    			fmt.Fprintf(w, "<div>%+v</div>", product)

    			queryPayments := datastore.NewQuery("Payment").Filter("ProductId =", product.Id)
	    		var payments []payment.Payment
	    		queryPayments.GetAll(context, &payments);
	    		for _, payment := range payments {
	    			fmt.Fprintf(w, "<div>%+v</div>", payment)
	    		}

    		}
	    }
    }
    fmt.Fprint(w, "<div>Done!</div>")
}

func handleRemoveAll(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)
	query := datastore.NewQuery("Account")
	var customers []account.Account
    if keys, err := query.GetAll(context, &customers); err == nil {
    	datastore.DeleteMulti(context, keys)
    }

    queryProducts := datastore.NewQuery("Product")
	var products []product.Product
    if keys, err := queryProducts.GetAll(context, &products); err == nil {
    	datastore.DeleteMulti(context, keys)
    }

    queryPayments := datastore.NewQuery("Payment")
	var payments []payment.Payment
    if keys, err := queryPayments.GetAll(context, &payments); err == nil {
    	datastore.DeleteMulti(context, keys)
    }

    fmt.Fprint(w, "<div>Done!</div>")
}

func handleAddNewCustomer(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)

	customer, product, payment, err := account.CreateCustomer(context, "John Doe", "danielalvarengacampos@gmail.com")
	if err != nil {
		fmt.Fprint(w, "<html><body>Error Creating Customer!</body></html>")
	}   

    if err := SaveToDatastore(context, "Account", customer); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := SaveToDatastore(context, "Product", product); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := SaveToDatastore(context, "Payment", payment); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	fmt.Fprintf(w, "<div>Customer created: %+v</div>", customer)
	fmt.Fprintf(w, "<div>Product created: %+v</div>", product)
	fmt.Fprintf(w, "<div>Payment created: %+v</div>", payment)
}

func SaveToDatastore(context appengine.Context, key string, entity interface{}) error {
	_, err := datastore.Put(context, datastore.NewIncompleteKey(context, key, nil), entity)
    if err != nil {
        return err
    }

	return nil
}
