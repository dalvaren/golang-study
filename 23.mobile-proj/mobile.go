package mobile

import (
	"fmt"
	"net/http"
	"html/template"
	"time"

	"account"
	"product"
	"payment"
	"content"
	"helper"

	"appengine"
    "appengine/user"
	"appengine/datastore"

	"github.com/gorilla/mux"
)

func init() {
	var router = mux.NewRouter()
	router.HandleFunc("/", handleMain).Methods("GET")
	router.HandleFunc("/dashboard", handleDashboard).Methods("GET")
	router.HandleFunc("/customer", handleListCustomers).Methods("GET")
	router.HandleFunc("/customer/new", handleNewCustomer).Methods("GET")
	router.HandleFunc("/customer/edit/{key}", handleEditCustomer).Methods("GET")
	router.HandleFunc("/customer/add", handleAddNewCustomer).Methods("POST")
	router.HandleFunc("/customer/{key}", handleListSpecificCustomer).Methods("GET")
	router.HandleFunc("/product/{key}", handleListSpecificProduct).Methods("GET")
	router.HandleFunc("/product/{key}/content/add", handleAddContent).Methods("GET")
	router.HandleFunc("/content/{productKey}/{name}", handleContent).Methods("GET")
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

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
    c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
        url, _ := user.LoginURL(c, "/")
        fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
        return
    }
    url, _ := user.LogoutURL(c, "/")
    fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}

func handleNewCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

    data := struct {
		Name  string
		Email string
		Key   string
	}{
		Name: "",
		Email: "",
		Key: "",
	}
    
    t, _ := template.ParseFiles("templates/customers-form.html")
    t.Execute(w, data)
}

func handleEditCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	
	vars := mux.Vars(r)
	key := vars["key"]
	if key == "" {
		http.Error(w, "key param error", http.StatusInternalServerError)
		return
	}

	context := helper.SetContext(r)
	accountData, err := account.Load(context, key)
	if helper.HasError(w, err, fmt.Sprintf("Error: %+v", err)) { return }

	var productKeys []string
	for _,productKey := range accountData.ProductKeys {
		productKeys = append(productKeys, productKey.Encode())
	}

    data := struct {
		Name  string
		Email string
		Key   string
		Products []string
	}{
		Name: accountData.Name,
		Email: accountData.Email,
		Key: key,
		Products: productKeys,
	}
    
    t, _ := template.ParseFiles("templates/customers-form.html")
    t.Execute(w, data)
}

func handleListCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	context := helper.SetContext(r)
	query := datastore.NewQuery("Account")
	var customers []account.Account
	customersKeys, err := query.GetAll(context, &customers);
    if helper.HasError(w, err, fmt.Sprintf("Error: %+v", err)) { return }

    type FlatAccount struct {
    	Name string
    	Email string
    	CreatedAt time.Time
    	Key string
    }

    var customersList []FlatAccount
    for index,customer := range customers {
    	singleCustomer := new(FlatAccount)
    	singleCustomer.Name = customer.Name
    	singleCustomer.Email = customer.Email
    	singleCustomer.CreatedAt = customer.CreatedAt
    	singleCustomer.Key = customersKeys[index].Encode()
    	customersList = append(customersList, *singleCustomer)
    }

    data := struct {
		List            []FlatAccount
		Title string
	}{
		List: customersList,
		Title: "Customers List",
	}
    
    t, _ := template.ParseFiles("templates/customers.html")
    t.Execute(w, data)
}

// http://localhost:8080/product/agpkZXZ-bW9iaWxlchQLEgdQcm9kdWN0GICAgICAwL8IDA
func handleListSpecificProduct(w http.ResponseWriter, r *http.Request) {
	context := helper.SetContext(r)

	vars := mux.Vars(r)
	key := vars["key"]
	if key == "" {
		http.Error(w, "key param error", http.StatusInternalServerError)
		return
	}

	product, err := product.Load(context, key)
	if helper.HasError(w, err, fmt.Sprintf("Error: %+v", err)) { return }

	fmt.Fprintf(w, "<div>product: %+v</div>", product)
}

// http://localhost:8080/customer/agpkZXZ-bW9iaWxlchQLEgdBY2NvdW50GICAgICA4JcLDA
func handleListSpecificCustomer(w http.ResponseWriter, r *http.Request) {
	context := helper.SetContext(r)

	vars := mux.Vars(r)
	key := vars["key"]
	if key == "" {
		http.Error(w, "key param error", http.StatusInternalServerError)
		return
	}

	account, err := account.Load(context, key)
	if helper.HasError(w, err, fmt.Sprintf("Error: %+v", err)) { return }

	fmt.Fprintf(w, "<div>account: %+v</div>", account)
}

func handleRemoveAll(w http.ResponseWriter, r *http.Request) {
	context := helper.SetContext(r)

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
	r.ParseForm()
	name := r.FormValue("name")
	email := r.FormValue("email")
	encodedKey := r.FormValue("key")

	context := helper.SetContext(r)
	if encodedKey != "" {
		accountData, err := account.Load(context, encodedKey)
		if helper.HasError(w, err, fmt.Sprintf("Error: %+v", err)) { return }
		accountData.Name = name
		accountData.Email = email
		key, _ := datastore.DecodeKey(encodedKey)
		_, err = datastore.Put(context, key, accountData)
	    if helper.HasError(w, err, fmt.Sprintf("Error: %+v", err)) { return }
	    fmt.Fprintf(w, "<div>Customer updated: %+v</div>", accountData)
	} else {
		customer, product, payment, err := account.CreateAccount(context, name, email)
		if helper.HasError(w, err, fmt.Sprintf("Error Creating Customer: %+v", err)) { return }
		fmt.Fprintf(w, "<div>Customer saved: %+v</div>", customer)
		fmt.Fprintf(w, "<div>Product saved: %+v</div>", product)
		fmt.Fprintf(w, "<div>Payment saved: %+v</div>", payment)
	}	

	
}

// http://localhost:8080/product/agpkZXZ-bW9iaWxlchQLEgdQcm9kdWN0GICAgICAwL8IDA/content/add
func handleAddContent(w http.ResponseWriter, r *http.Request) {
	context := helper.SetContext(r)

	vars := mux.Vars(r)
	key := vars["key"]
	if key == "" {
		http.Error(w, "key param error", http.StatusInternalServerError)
	}

	product, err := product.Load(context, key)
	if helper.HasError(w, err, fmt.Sprintf("Error: %+v", err)) { return }
	fmt.Fprintf(w, "<div>product: %+v</div>", product)

	content, err := content.NewContent("page1", "{v:1,j:{}}")
	if helper.HasError(w, err, fmt.Sprintf("Error: %+v", err)) { return }

	content.ProductKey, _ = datastore.DecodeKey(key)
	contentKey, err := content.Save(context)
	if helper.HasError(w, err, fmt.Sprintf("Error: %+v", err)) { return }

	product.ContentKeys = append(product.ContentKeys, contentKey)

	fmt.Fprintf(w, "<div>Content created: %+v</div>", content)
	fmt.Fprintf(w, "<div>Product: %+v</div>", product)
}

// http://localhost:8080/content/agpkZXZ-bW9iaWxlchQLEgdQcm9kdWN0GICAgICAwL8IDA/page2
func handleContent(w http.ResponseWriter, r *http.Request) {
	context := helper.SetContext(r)

	w.Header().Add("Content-Type", "application/json")
	
	vars := mux.Vars(r)
	encodedProductKey := vars["productKey"]
	if encodedProductKey == "" {
		http.Error(w, "Parameter shall not be empty", http.StatusInternalServerError)
		return
	}
	contentName := vars["name"]
	if contentName == "" {
		http.Error(w, "Parameter shall not be empty", http.StatusInternalServerError)
		return
	}
	loadedProduct, err := product.Load(context, encodedProductKey)
	if helper.HasError(w, err, "That was not possible to load the product key") { return }

	if productIsPaid := loadedProduct.IsValid(context); productIsPaid == false {
		http.Error(w, "There is a problem with your account, please enter in contact with the support team.", http.StatusInternalServerError)
		return
	}

	productKey, _ := datastore.DecodeKey(encodedProductKey)
	queryContents := datastore.NewQuery("Content").
		Filter("ProductKey =", productKey).
		Filter("Name =", contentName).
		Order("-Version").
		Limit(1)
	var contents []content.Content
    if _, err := queryContents.GetAll(context, &contents); err == nil {
    	for _,content := range contents {
    		fmt.Fprint(w, content.JsonData)
    	}
    }

}
