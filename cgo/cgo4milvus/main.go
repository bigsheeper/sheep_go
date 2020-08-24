package main

/*

#cgo CFLAGS: -I./

#cgo LDFLAGS: -L./build -lsegmentbase -Wl,-rpath=./build

#include "cwrap.h"

*/
import "C"
import "fmt"

func test_insert() {
	var segment = C.CSegmentBase();

}

func main() {
	fmt.Println("Test milvus segment base:")
	test_insert()
}
