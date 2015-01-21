package main

import (
	"fmt"
	"entity"
	"reflect"
)

func main() {
	fmt.Println("Starting application...")
	fmt.Println()
	myDog := entity.NewDog("Rex")
	myCat := entity.NewCat("Sophia")
//	myDog.Walk()
//	myCat.Walk()
//	myCat.Run()
	entity.WalkWithAnimal(myDog)
	entity.WalkWithAnimal(myCat)
	fmt.Println()

	myNewAnimal := entity.NewEntity("cat")
	myNewAnimal.Walk()
	fmt.Println()

	fmt.Println(reflect.TypeOf(myDog))
	fmt.Println(reflect.TypeOf(myCat))
	fmt.Println(reflect.TypeOf(myNewAnimal))
	fmt.Println()

	fmt.Println(reflect.ValueOf(myCat).NumMethod())
	for i:=0;i<reflect.ValueOf(myCat).NumMethod();i++ {
		fmt.Println(reflect.ValueOf(myCat).Method(i))
	}
	fmt.Println()
}
