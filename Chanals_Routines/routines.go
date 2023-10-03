package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Random(ch chan int) {
	for i := 0; i < 6; i++ {
		randomNum := rand.Intn(100)
		ch <- randomNum
	}
	close(ch)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func Average(ch chan int) int {
	var s []int
	for num := range ch {
		s = append(s, num)
	}

	if len(s) == 0 {
		return 0
	}

	c := make(chan int)
	go sum(s, c)
	x := <-c
	return x / len(s)
}

func ShowAverage() {
	ch := make(chan int)
	go Random(ch)
	num := Average(ch)
	fmt.Println(num)
}

func main() {
	ShowAverage()
	time.Sleep(100 * time.Millisecond)
}
