package main

import "fmt"

func mensajef(text string, c chan string) {
	c <- text
}

func main() {

	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Println(len(c), cap(c))

	// Range y close
	close(c) // cierra el canal y no va a recibir ningún dato más

	// c <- "Mensaje3"
	for mensaje := range c {
		fmt.Println(mensaje)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)
	go mensajef("mensaje1", email1)
	go mensajef("mensaje2", email2)
	go mensajef("mensaje3", email1)
	go mensajef("mensaje4", email2)
	go mensajef("mensaje5", email1)
	go mensajef("mensaje6", email2)
	for i := 0; i < 6; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}
