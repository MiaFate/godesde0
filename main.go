package main

import (
	"fmt"

	"github.com/miafate/hellogo/ejercicios"
)

func main() {
	// estado, texto := variables.ConvertirATexto(14)
	// fmt.Println(estado)
	// fmt.Println(texto)
	// if os := runtime.GOOS; os == "linux" || os == "OS X." {
	// 	fmt.Println("esto no es windows, es", os)
	// } else {
	// 	fmt.Println("esto es windows")
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("esto es Linux")
	// case "darwin":
	// 	fmt.Println("esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }
	numero, texto := ejercicios.ConvertirAInt("10")
	fmt.Println(numero, texto)
}
