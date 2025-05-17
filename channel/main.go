package main

import (
	"fmt"
	"sync"
)

func main() {

	//Channel buffering
	ch2 := make(chan int, 5)
	go func() {
		ch2 <- 1
		ch2 <- 2
		ch2 <- 3
		ch2 <- 4
		ch2 <- 5
		close(ch2)    //close to avoid hanging
	}()
	i := 0
	for val := range ch2 {
		fmt.Println("buffering example wala output hai, val =", val)
		if i == 4 {
			break
		} else {
			i++
		}
	}
	//partial sync
	var ch chan int = make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go sendNumber(ch)
	go recieveNumber(ch, &wg)
	wg.Wait()

}

func sendNumber(ch chan int) {

	for i := 0; i <= 10; i++ {
		ch <- i

	}
	close(ch)

}
func recieveNumber(ph chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	//for i := 0; i <= 10; i++ {

	for val := range ph {
		fmt.Println(val)
	}
}

//}
