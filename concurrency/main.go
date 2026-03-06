package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	TableNumber int
	Preptime    time.Duration
}

// func main() {
// 	orders := []Order{
// 		{TableNumber: 1, Preptime: time.Second * 2},
// 		{TableNumber: 2, Preptime: time.Second * 3},
// 		{TableNumber: 3, Preptime: time.Second * 4},
// 	}

// 	group01 := sync.WaitGroup{} //creating a wait group
// 	group01.Add(len(orders))    // and saying how many routines are there, its 3 the len of orders

// 	for _, order := range orders {
// 		go processOrder(order)
// 	}

// 	group01.Wait()
// }

func main() {
	orders := []Order{
		{TableNumber: 1, Preptime: time.Second * 2},
		{TableNumber: 2, Preptime: time.Second * 3},
		{TableNumber: 3, Preptime: time.Second * 4},
	}

	wg := sync.WaitGroup{}

	for _, order := range orders {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processOrder(order)
		}()
	}

	wg.Wait()
}

func processOrder(order Order) {
	//simulate order time
	fmt.Printf("Preparing Order for table number: %d\n", order.TableNumber)

	time.Sleep(order.Preptime)

	//order is ready
	fmt.Printf("The Order is ready. The table number is %d \n", order.TableNumber)
}
