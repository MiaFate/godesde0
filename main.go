package main

import (
	"fmt"
	"runtime"
)

func main() {
	// estado, texto := variables.ConvertirATexto(14)
	// fmt.Println(estado)
	// fmt.Println(texto)
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("esto no es windows, es", os)
	} else {
		fmt.Println("esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("esto es Linux")
	case "darwin":
		fmt.Println("esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}
