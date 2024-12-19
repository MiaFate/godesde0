package main

import (
	"github.com/miafate/godesde0/ejer_interfaces"
	"github.com/miafate/godesde0/modelos"
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
	// numero, texto := ejercicios.ConvertirAInt("10")
	// fmt.Println(numero, texto)

	//teclado.IngresoNumeros()

	//iteraciones.Iterar()
	// contar_letras.ContarLetras()
	// prueba.Prueba1Funcion()
	// fmt.Println(ejercicios.ObtenerTablas())
	// files.SumarTabla()
	// files.LeerArchivo2()
	// funciones.LlamarClosure()
	// funciones.Exponencia(2)
	// arreglos_slices.Capacidad()
	// mapas.MostrarMapas()
	//users.AltaUsuario()
	Mia := new(modelos.Mujer)
	ejer_interfaces.HumanosRespirando(Mia)

	Pedro := new(modelos.Hombre)
	ejer_interfaces.HumanosRespirando(Pedro)
}
