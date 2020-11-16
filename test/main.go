package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"go.uber.org/atomic"
)

type Counter atomic.Int32

func produce(integers chan int) {
	counter := 0
	for {
		fmt.Println("produce", counter)
		integers <- 1
		counter++
		time.Sleep(20 * time.Millisecond)
	}
}

func consume(integers chan int) {
	counter := 0
	for {
		fmt.Println("consume", counter)
		<-integers
		counter++
		time.Sleep(10 * time.Millisecond)
	}
}

func TestMap() {
	mm := make(map[int]int, 0)
	fmt.Println(mm)
	if _, ok := mm[0]; !ok {
		fmt.Println("Not exists")
	}
	fmt.Println(mm)
}

func TestChannel() {
	var integers = make(chan int, 100)

	go produce(integers)
	consume(integers)
}

func TestCopy() {
	const N = 200000
	var arr1 = make([]int, N)
	var offset = 0
	for i := 0; i < N; i++ {
		arrTmp := make([]int, 1)
		arrTmp[0] = i
		copy(arr1[offset:], arrTmp)
		offset += len(arrTmp)
	}

	fmt.Println(arr1)
}

func TestMarshal() {
	type InsertLog struct {
		MsgLength              int
		DurationInMilliseconds int64
		NumSince               int64
		Speed                  float32
		String                 string
	}

	insertLog := InsertLog{123, 456, 51354, 0.43453, "fdgdfg"}

	fmt.Println(insertLog)
	insertLogJson, err := json.Marshal(&insertLog)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("insertLogJson", string(insertLogJson))
}

func TestWrite() {
	f, err := os.Create("/tmp/aaa.txt")
	if err != nil {
		log.Fatal(err)
	}

	// write logs
	for i := 0; i < 10; i++ {
		writeString := "54523436134362341353464552453423542\n"

		_, err2 := f.WriteString(writeString)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	fmt.Println("write log done")
}

func readMap(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	wg.Done()
}

func TestMapInGoroutine() {
	const N = 100
	testM1 := make(map[int]int)
	for i := 0; i < N; i++ {
		testM1[i] = i * 100
	}

	wg := sync.WaitGroup{}

	for k, _ := range testM1 {
		wg.Add(1)
		go readMap(k, &wg)
	}

	wg.Wait()
}

func TestFmtSprint() {
	const N = 10
	test := make([]int, 0)
	for i := 0; i < N; i++ {
		test = append(test, i)
	}

	s := fmt.Sprint(test)
	fmt.Println(s)
}

func TestGoroutine() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func TestScan() {
	var num float64
	_, err := fmt.Scanf("%f", &num)
	if err != nil {
		panic(err)
	}

	fmt.Printf("You have entered : %f \n", num)

	fmt.Println("Alternative output ", strconv.FormatFloat(num, 'f', 6, 64))
}

func sendMsgFromCmd() {
	for {
		var num float64
		_, err := fmt.Scanf("%f", &num)
		fmt.Println(num)
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
}

func testChannelBlock(ctx context.Context, c chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-c:
			fmt.Println(msg)
		}
	}
}

func testEndlessLoop(ctx context.Context, c chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if len(c) > 0 {
				fmt.Println(<-c)
			} else {
				continue
			}
		}
	}
}

func sendMsgTmp(c chan int) {
	for {
		c <- 10
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}

func testCpuUsage() {
	ctx := context.Background()
	c := make(chan int, 100)
	go sendMsgTmp(c)
	go testEndlessLoop(ctx, c)
	for {
		select {}
	}
}

type Animal interface {
	Speak()
}

type Dog struct{}

func (d *Dog) Speak() { fmt.Println("wang wang wang") }

type ChineseDog struct {
	Dog
	name string
}

func TestInterface() {
	var d Animal
	_, ok := d.(*ChineseDog)
	fmt.Println(ok)
}

func main() {
	testSort()
}
