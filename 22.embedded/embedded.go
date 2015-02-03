package embedded

import (
	"fmt"
	"net/http"
	// "strings"

	"appengine"
	"appengine/datastore"
)

type A struct {
	X int
	Y int
}

type C struct {
	Z string
}

type D struct {
	Contents []string
}

type B struct {
	I int
	InnerA A
	ListC []C
	ListD []D
}


func (x *B) Load(c <-chan datastore.Property) error {
    // Load I and J as usual.
    if err := datastore.LoadStruct(x, c); err != nil {
        return err
    }
    // Derive the Sum field.
    // x.Sum = x.I + x.J
    return nil
}

func (x *B) Save(c chan<- datastore.Property) error {
    defer close(c)

    c <- datastore.Property{
        Name:  "I",
        Value: int64(x.I),
    }

    c <- datastore.Property{
        Name:  "InnerA.X",
        Value: int64(x.InnerA.X),
    }

    c <- datastore.Property{
        Name:  "InnerA.Y",
        Value: int64(x.InnerA.Y),
    }

    for _, v := range x.ListC {
        c <- datastore.Property{Name: "ListC", Value: string(v.Z), Multiple: true}
    }

    for _, v := range x.ListD {
    	for _, content := range v.Contents {
    		c <- datastore.Property{Name: "ListD.Contents", Value: content, Multiple: true}
    	}
    }

    // Save I and J as usual. The code below is equivalent to calling
    // "return datastore.SaveStruct(x, c)", but is done manually for
    // demonstration purposes.
    // c <- datastore.Property{
    //     Name:  "I",
    //     Value: int64(x.I),
    // }
    // c <- datastore.Property{
    //     Name:  "J",
    //     Value: int64(x.J),
    // }
    return nil
    // return datastore.SaveStruct(x, c)
}

func init() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/search", handleSearch)
	http.HandleFunc("/delete", handleRemoveAll)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)

    query := datastore.NewQuery("Test")
    			// Filter("I =", 1)
    			// Filter("InnerA.X =", 10)
    			// Filter("ListC.Z =", "Mary")
    			// Filter("ListC.Z =", "Mary")
    var items []B
    query.GetAll(context, &items);
    for _, item := range items {
		fmt.Fprintf(w, "<div>%+v</div>", item)
		fmt.Fprintf(w, "<div>---</div>", item)
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)

	entityA := new(A)
	entityA.X = 10
	entityA.Y = 5

	entityCa := new(C)
	entityCa.Z = "Mary"
	
	entityCb := new(C)
	entityCb.Z = "Kay"

	entityDa := new(D)
	entityDa.Contents = append(entityDa.Contents, "New")
	entityDa.Contents = append(entityDa.Contents, "York")

	entityDb := new(D)
	entityDb.Contents = append(entityDb.Contents, "São")
	entityDb.Contents = append(entityDb.Contents, "Paulo")

	entityB := new(B)
	entityB.I = 1
	entityB.InnerA = *entityA
	entityB.ListC = append(entityB.ListC, *entityCa)
	entityB.ListC = append(entityB.ListC, *entityCb)
	entityB.ListD = append(entityB.ListD, *entityDa)
	entityB.ListD = append(entityB.ListD, *entityDb)

	if err := SaveToDatastore(context, "Test", entityB); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	fmt.Fprintf(w, "<div>Retrieved:</div>")
    query := datastore.NewQuery("Test")
    var items []B
    query.GetAll(context, &items);
    for _, item := range items {
		fmt.Fprintf(w, "<div>%+v</div>", item)
		fmt.Fprintf(w, "<div>---</div>", item)
	}
}

func SaveToDatastore(context appengine.Context, key string, entity interface{}) error {
	_, err := datastore.Put(context, datastore.NewIncompleteKey(context, key, nil), entity)
    if err != nil {
        return err
    }

	return nil
}

func handleRemoveAll(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)
	query := datastore.NewQuery("Test")
	var items []B
    if keys, err := query.GetAll(context, &items); err == nil {
    	datastore.DeleteMulti(context, keys)
    }

    fmt.Fprint(w, "<div>Done!</div>")
}