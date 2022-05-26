package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

type car struct {
	año    int
	marca  string
	precio int
}

// estamos personalizando el método string sobre-escribiéndolo asociado al struct pc
func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	myPC := pc{ram: 16, brand: "msi", disk: 100}
	fmt.Println(myPC)
	otroPC := pc{ram: 64, brand: "otra Marca", disk: 200}
	fmt.Println(otroPC)

	micoche := car{año: 20100, marca: "Hyundai", precio: 20000}
	fmt.Println(micoche)
}
