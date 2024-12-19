package arreglos_slices

import "fmt"

// arreglo
var tabla [10]int = [10]int{10, 0, 59}
var matriz [4][4]int

func MuestroArreglos() {

	tabla[7] = 33
	tabla[2] = 54

	tabla2 := [10]string{"Pablo", "Juan"}

	fmt.Println("vector: ", tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla[i])
	}

	matriz[3][3] = 15
	fmt.Println(matriz)
}
