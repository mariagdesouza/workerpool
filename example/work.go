package main

import (
	"fmt"
	"sync"

	"github.com/mariadesouza/workerpool"
)

func findandStoreFibonacci(n int, fibonacci *int64, wg *sync.WaitGroup) {
	var x, y int64
	x, y = 1, 1
	for i := 1; i < n; i++ {
		x, y = y, x+y
	}
	*fibonacci = y
	wg.Done()
}

func main() {
	s := 20
	var fibonacci [20]int64
	wp := workerpool.New(3)
	defer wp.Close()
	var wg sync.WaitGroup
	for i := 0; i < s; i++ {
		wg.Add(1)
		index := i
		wp.AddWorkToPool(
			func() {
				findandStoreFibonacci(index, &fibonacci[index], &wg)
			})
		//fmt.Printf("+")
	}
	wg.Wait()
	fmt.Println(fibonacci)
}
