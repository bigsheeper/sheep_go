package main

/*

#cgo CFLAGS: -I./

#cgo LDFLAGS: -L./build -ltest -Wl,-rpath=./build

#include <stdio.h>
#include <stdlib.h>

#include "test.h"

void myprint() {
    printf("Hello world\n");
}

*/
import "C"
import "fmt"
import "unsafe"
import "reflect"

const length int = 10

func test_print() {
	C.myprint()
}

func test_sum() {
	fmt.Println("Hello from go!")
	fmt.Println(C.sum(1, 1))
}

func test_go_ptr() {
	a := make([]float64, length)
	C.assign((*C.double)(&a[0]))
	fmt.Println(a[5])
}

func test_c_ptr() {
	var b unsafe.Pointer = C.c_malloc(C.int(8 * length))
	C.cptr_test(b)

	fmt.Println("Pointer type: ", reflect.TypeOf(b))
	if (b == nil) {
		fmt.Println("nullptr error")
		return
	}

	// pointer to array
	b_arr := *(*[length]float64)(b)
	fmt.Println("Array type: ", reflect.TypeOf(b_arr))
	fmt.Println(b_arr)

	C.c_free(b)
}

func main() {
	fmt.Println("Test print:")
	test_print()
	fmt.Println("\n\n\nTest sum:")
	test_sum()
	fmt.Println("\n\n\nTest go ptr:")
	test_go_ptr()
	fmt.Println("\n\n\nTest c ptr:")
	test_c_ptr()
}
