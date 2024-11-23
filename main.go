package main

import (
	"fmt"

	"github.com/miafate/hellogo/variables"
)

func main() {
	estado, texto := variables.ConvertirATexto(14)
	fmt.Println(estado)
	fmt.Println(texto)
}
