package ejer_interfaces

import (
	"fmt"

	"github.com/miafate/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy humano %s y respiro \n", hu.Genero())
	fmt.Printf("Estoy vivo? %t \n", hu.EstaVivo())
}
