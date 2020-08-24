package main

/*

#cgo CFLAGS: -I./

#cgo LDFLAGS: -L./build -lcwrap -Wl,-rpath=./build

#include "cwrap.h"

*/
import "C"
import "fmt"

func test_person() {
	var person = C.PersonInit(5)
	var age = C.age(person)
	fmt.Println(age)
}

func main() {
	fmt.Println("Test person:")
	test_person()
}
