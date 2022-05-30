package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go say("hola", &wg)
	go say("adios", &wg)
	go say("UNO", &wg)
	go say("DOS", &wg)
	go say("TRES", &wg)
	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios sss")
	time.Sleep(time.Second * 1)

}
