package main

import (
	"fmt"
	"sync"
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
	fmt.Println("A  AA  DI OS ")

}
