package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {

	// Datos para test
	num1 := 20
	num2 := 5
	esperado := 15

	// Act

	resultado := Restar(num1, num2)

	// Test

	assert.Equal(t, esperado, resultado, "Funcion Resta() dio resultado: %d, y se esperaba %d", resultado, esperado)
	/*if resultado != esperado {

		t.Errorf("Funcion Resta() dio resultado: %d, y se esperaba %d", resultado, esperado)

	}
	*/
}
