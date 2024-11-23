package ejercicios

import "strconv"

func ConvertirAInt(entrada string) (int, string) {
	numero, err := strconv.Atoi(entrada)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}
}
