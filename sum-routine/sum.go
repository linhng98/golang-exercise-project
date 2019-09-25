package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sumArr(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var arr [1000000]int
	var sum int
	c := make(chan int)

	for i := range arr {
		arr[i] = rand.Intn(1000)
	}

	start1 := time.Now()
	sum = 0
	for _, v := range arr {
		sum += v
	}
	t1 := time.Since(start1) * time.Nanosecond
	fmt.Printf("%d    %v\n", sum, t1)

	for i := 0; i < len(arr); i += 10 {
		go sumArr(arr[i:i+10], c)
	}
	start2 := time.Now()
	sum = 0
	for i := 0; i < len(arr); i += 10 {
		sum += <-c
	}
	t2 := time.Since(start2) * time.Nanosecond
	fmt.Printf("%d    %v\n", sum, t2)
}
