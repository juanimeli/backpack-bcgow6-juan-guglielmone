package ordenamiento

import "testing"

func TestRestar(t *testing.T) {

	// Datos para test
	unorderSlice := []int{3, 2, 4, 1}
	esperado := []int{1, 2, 3, 4}

	// Act

	resultado := OrderSlice(unorderSlice)

	// Test

	for i, _ := range resultado {
		if resultado[i] != esperado[i] {
			t.Errorf("la funcion no ordeno correctamente el slice")
		}
	}
	/*
		for i, num := range(resultado) {

			if i == 0 {
				continue
			}
			if num < resultado[i-1] {
				t.Errorf("la funcion no ordeno correctamente el slice")

			}

		}
	*/
}
