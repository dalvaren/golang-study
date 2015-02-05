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

	"github.com/gorilla/mux"
)

func init() {
	var router = mux.NewRouter()
	router.HandleFunc("/", handleMain).Methods("GET")
	router.HandleFunc("/customer", handleListCustomers).Methods("GET")
	router.HandleFunc("/customer/add", handleAddNewCustomer).Methods("GET")
	router.HandleFunc("/customer/{key}", handleListSpecificCustomer).Methods("GET")
	router.HandleFunc("/product/{key}", handleListSpecificProduct).Methods("GET")
	router.HandleFunc("/product/content/add", handleAddContent).Methods("GET")
	router.HandleFunc("/product/content/{key}", handleContent).Methods("GET")
	router.HandleFunc("/delete", handleRemoveAll).Methods("GET")
	http.Handle("/", router)
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

func handleListCustomers(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)
	query := datastore.NewQuery("Account")
	var customers []account.Account
    if _, err := query.GetAll(context, &customers); err == nil {
    	fmt.Fprint(w, "<div>Customers:</div>")
    	for index, customer := range customers {
    		fmt.Fprintf(w, "<div>%d. ---</div>", index)
    		fmt.Fprintf(w, "<div>%+v</div>", customer)

    		for _, productKey := range customer.ProductKeys {
    			var product product.Product
    			err := datastore.Get(context, productKey, &product);
	    		if err != nil {
	    			fmt.Fprintf(w, "<div>err: %+v</div>", err)
	    		}
    			fmt.Fprintf(w, "<div>product: %+v</div>", product)
    			encodedKey := productKey.Encode();
    			fmt.Fprintf(w, "<div>encodedKey: %+v</div>", encodedKey)

    			var payment payment.Payment
    			err = datastore.Get(context, product.PaymentKey, &payment);
	    		if err != nil {
	    			fmt.Fprintf(w, "<div>err: %+v</div>", err)
	    		}
    			fmt.Fprintf(w, "<div>payment: %+v</div>", payment)
    		}
	    }
    }
    fmt.Fprint(w, "<div>Done!</div>")
}

// http://localhost:8080/product/agpkZXZ-bW9iaWxlchQLEgdQcm9kdWN0GICAgICAwL8IDA
func handleListSpecificProduct(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)

	vars := mux.Vars(r)
	key := vars["key"]
	if key == "" {
		http.Error(w, "key param error", http.StatusInternalServerError)
	}

	product, err := product.Load(context, key)
	if err != nil {
		fmt.Fprintf(w, "<div>err: %+v</div>", err)
	}
	fmt.Fprintf(w, "<div>product: %+v</div>", product)
}

// http://localhost:8080/customer/agpkZXZ-bW9iaWxlchQLEgdBY2NvdW50GICAgICA4JcLDA
func handleListSpecificCustomer(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)

	vars := mux.Vars(r)
	key := vars["key"]
	if key == "" {
		http.Error(w, "key param error", http.StatusInternalServerError)
	}

	account, err := account.Load(context, key)
	if err != nil {
		fmt.Fprintf(w, "<div>err: %+v</div>", err)
	}
	fmt.Fprintf(w, "<div>product: %+v</div>", account)
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

	customer, product, payment, err := account.CreateAccount(context, "John Doe", "danielalvarengacampos@gmail.com")
	if err != nil {
		fmt.Fprintf(w, "<html><body>Error Creating Customer! %+v</body></html>", err)
	}

	fmt.Fprintf(w, "<div>Customer created: %+v</div>", customer)
	fmt.Fprintf(w, "<div>Product created: %+v</div>", product)
	fmt.Fprintf(w, "<div>Payment created: %+v</div>", payment)
}
