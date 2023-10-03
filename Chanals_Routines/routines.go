package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Random() []int {
	ch := make(chan []int)
	for i := 0; i < 6; i++ {
		randomNums := rand.Perm(10)
		time.Sleep(time.Second)
		ch <- randomNums
	}
	return <-ch
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func Average() int {
	s := Random()
	c := make(chan int)
	go sum(s[:], c)
	x := <-c
	return x / len(s)
}
func ShowAverage() {
	num := Average()
	time.Sleep(100 * time.Millisecond)
	go fmt.Println(num)
}

func main() {
	ShowAverage()
}
