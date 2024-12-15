package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ObtenerTablas() string {
	var err error
	var numero int
	var texto string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese un numero: ")
		scanner.Scan()
		numero, err = strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("eso no es un numero...")
			continue
		}
		break
	}
	texto += fmt.Sprintln("TABLA DEL ", numero)
	texto += fmt.Sprintln("---------------")
	for i := 1; i <= 10; i++ {
		resultado := numero * i
		// fmt.Println(numero, " x ", i, " = ", resultado)
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, resultado)
	}
	return texto
}
