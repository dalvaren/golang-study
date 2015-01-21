package entity

import (
	"fmt"
)

type entity interface {
	Walk()
}

type Any interface {}

type dog struct {
	name string
}

func NewEntity(animalType string) (entity){
	if len(animalType) < 1 {
		return nil
	}
	if(animalType == "dog") {
		return NewDog("Generic Dog")
	}
	return NewCat("Generic Cat")
}

func NewDog(name string) *dog {
	if len(name) < 1 {
		return nil
	}
	dog := new(dog)
	return dog
}

type cat struct {
	name string
}

func NewCat(name string) *cat {
	if len(name) < 1 {
		return nil
	}
	cat := new(cat)
	return cat
}

func (this *dog) Walk() {
	fmt.Println("I'm a dog and I'm walking.")
}

func (this *cat) Walk() {
	fmt.Println("I'm a cat and I'm walking.")
}

func (this *cat) Run() {
	fmt.Println("I'm a cat and I'm running.")
}

func WalkWithAnimal(animal entity) {
	animal.Walk()
}
