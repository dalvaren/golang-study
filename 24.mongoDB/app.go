package main

import (
        "fmt"
		"log"
		"errors"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type Person struct {
        Name string
        Phone string
}

func main() {
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("people")

        if err := savePerson(c, "Danilo", "+55 53 8402 8513"); err != nil {
                log.Fatal(err)
        }

        result := Person{}
        err = c.Find(bson.M{"name": "Daniel"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Phone:", result.Phone)
}

func savePerson(c *mgo.Collection, name string, phone string) error {
	if checkExistentPhone(c, phone) == true {
		return errors.New("Existent Phone!")
	}
	if checkExistentName(c, name) == true {
		return errors.New("Existent Name!")
	}
	err := c.Insert(&Person{name, phone})
	return err
}

func checkExistentPhone(c *mgo.Collection, phone string) bool{
	result := Person{}
    err := c.Find(bson.M{"phone": phone}).One(&result)
    if err != nil {
    	return false
    }
    return true
}

func checkExistentName(c *mgo.Collection, name string) bool{
	result := Person{}
    err := c.Find(bson.M{"name": name}).One(&result)
    if err != nil {
    	return false
    }
    return true
}