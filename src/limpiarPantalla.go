package main

import (
	"fmt"
	"os"
	"os/exec"
)

func limpiar() {
	cmd := exec.Command("clear")
	//cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	limpiar()

	fmt.Println("Hola")

}
