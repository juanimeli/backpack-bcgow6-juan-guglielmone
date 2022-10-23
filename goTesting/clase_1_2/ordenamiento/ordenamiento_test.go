package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {

	// Datos para test ARRANGE
	unorderSlice := []int{3, 2, 4, 1}
	esperado := []int{1, 2, 3, 4}

	// Act  ACT

	resultado := OrderSlice(unorderSlice)

	// Test ASSERT

	assert.Equal(t, esperado, resultado) // CON LIBRERIA ASSERT

	/* ESTA OPCION ES SIN LIBRERIA ASSERT (NOSOTROS HACEMOS LA LOGICA)
	for i := range resultado {
		if resultado[i] != esperado[i] {
			t.Errorf("la funcion no ordeno correctamente el slice")
		}
	}
	*/
	/* ESTA OPCION ES MEDIA RARA PORQUE EN REALIDAD NO ESTOY COMPARANDO CON EL ESPERADO DE ARRANGE
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
