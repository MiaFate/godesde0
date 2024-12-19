package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 21, 12, 33, 1, 44}

func MuestroSlice() {
	fmt.Println(tablaS)
	porcion := arreglo[3:]   // Slice creado desde un vector desde la posicion 3 hasta el final
	porcion2 := arreglo[:5]  // Slice creado desde un vector desde la posicion 0 hasta la posicion 5
	porcion3 := arreglo[2:4] // Slice creado desde un vector desde la posicion 2 hasta la posicion 4
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
