package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex := sync.Mutex{}
	accBalance := 0

	wg := sync.WaitGroup{}
//wg.Add(1000)
	for i := 1; i <= 1000; i++ {
		wg.Add(1)

		go paisaBadhao(&accBalance, &wg, &mutex)
	}
	wg.Wait()
	fmt.Println(accBalance)

}

func paisaBadhao(balance *int, wg *sync.WaitGroup, tala *sync.Mutex) {

	defer wg.Done()
	tala.Lock()
	*balance++
	tala.Unlock()

}
