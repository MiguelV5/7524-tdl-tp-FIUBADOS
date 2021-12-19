package main

import (
	"fmt"
	"sync"
)

func printtest(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	wg.Add(3)
	go printtest("a", &wg)
	go printtest("b", &wg)
	go printtest("c", &wg)
	wg.Wait()

}
