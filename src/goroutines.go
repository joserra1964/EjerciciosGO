package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done() // esta función es la que resta 1 al indicar
	// goroutine hecha
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hola")
	// say("Hola")
	wg.Add(1) // el numero 1 indica cuantas goroutines debe esperar
	go say("Mundo", &wg)
	wg.Wait() // cada vez que una goroutine termina, disminuye
	// el contador en 1, hasta que llega a 0 y continua el programa

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)

}
