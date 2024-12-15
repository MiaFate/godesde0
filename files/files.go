package files

import (
	"bufio"
	"fmt"
	"github.com/miafate/godesde0/ejercicios"
	"os"
)

var filename string = "./files/txt/tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.ObtenerTablas()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func SumarTabla() {
	var texto string = ejercicios.ObtenerTablas()
	texto += fmt.Sprintln("")
	if !Append(filename, texto) {
		fmt.Println("Error al concatenar contenido")
		fmt.Println("se creara el archivo")
		GrabarTabla()
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el append " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("Error durante el writestring " + err.Error())
		return false
	}

	arch.Close()
	return true
}

// lee todo el archivo
func LeerArchivo() {
	archivo, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}
	fmt.Println(string(archivo))
}

// lee linea por linea
func LeerArchivo2() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
