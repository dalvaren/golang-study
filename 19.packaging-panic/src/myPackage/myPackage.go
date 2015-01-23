package myPackage

import (
	"fmt"
)

func Test() (number int, err error){
	defer func(){
		if r := recover(); r != nil {
			var isOk bool
			err, isOk = r.(error)
			if !isOk {
				err = fmt.Errorf("Panicking in myPackage: %v\n", r)
			}
		}
	}()

	number = 1
	BadCall()
	fmt.Println("After bad call...")
	return
}

func BadCall() {
	panic("bad end")
}
