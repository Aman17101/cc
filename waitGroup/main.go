package main

import (
	"fmt"
	"sync"
)

func a(g *sync.WaitGroup) {
	defer g.Done()
	fmt.Println(" a is runn ing ...")

}
func main() {
	wg := sync.WaitGroup{}
	//var wg sync.WaitGroup
	wg.Add(1)
	go a(&wg)
	fmt.Println("-------------")
	//fmt.Println("**********")
	wg.Wait()
	fmt.Println("**********")

}
