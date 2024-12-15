package ejercicios

import (
	"bufio"
	"fmt"
	"os"
  "strconv"
)

var numero int
var err error

func Ejercicio2()  {

  inputOK:= false

  for inputOK == false {
	  scanner := bufio.NewScanner(os.Stdin)
	  fmt.Println("Ingrese un numero")
    scanner.Scan()
		numero, err = strconv.Atoi(scanner.Text())
    if err != nil {
    continue
    }
    inputOK = true
  }

  for i := 1; i <= 9; i++{
  resultado := numero * i
    fmt.Println(numero, " x ", i," = ", resultado)
  }

}


