package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func generateOrder(customerID int, products []Product,
	orderChannel chan<- map[string]int) {
	for {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		quantity := rand.Intn(5) + 1
		order := map[string]int{
			"CustomerID": customerID,
			"ProductID":  rand.Intn(len(products)),
			"Quantity":   quantity,
		}
		orderChannel <- order
	}
}

func processOrders(products []Product, orderChannel <-chan map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	customerOrders := make(map[int]map[int]int)
	for {
		order := <-orderChannel
		customerID := order["CustomerID"]
		productID := order["ProductID"]
		quantity := order["Quantity"]
		product := products[productID]
		orderValue := product.Price * float64(quantity)

		if _, exists := customerOrders[customerID]; !exists {
			customerOrders[customerID] = make(map[int]int)
		}
		if _, exists := customerOrders[customerID][productID]; !exists {
			customerOrders[customerID][productID] = 0
		}
		customerOrders[customerID][productID] += quantity
		fmt.Printf("Замовлення від покупця %d: %s x%d\n", customerID, product.Name, quantity)
		fmt.Printf("Загальна вартість замовлення: %.2f\n", orderValue)
	}

}

func main() {
	numCustomers := flag.Int("customers", 3, "Кількість покупців")
	flag.Parse()

	products := []Product{
		{Name: "Товар 1", Price: 10.0},
		{Name: "Товар 2", Price: 20.0},
		{Name: "Товар 3", Price: 30.0},
	}
	orderChannel := make(chan map[string]int)
	var wg sync.WaitGroup
	for i := 1; i <= *numCustomers; i++ {
		wg.Add(1)
		go generateOrder(i, products, orderChannel)
		go processOrders(products, orderChannel, &wg)
	}

	wg.Wait()
}
