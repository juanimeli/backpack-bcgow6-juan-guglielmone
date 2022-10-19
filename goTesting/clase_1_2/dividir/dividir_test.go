package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivision(t *testing.T) {

	//Datos
	num1 := 10
	den := 2
	esperado := 5

	//test

	resultado, _ := Dividir(num1, den)

	//validate

	if resultado != esperado {
		t.Errorf("El resultado de Dividir() fue %d y lo esperado es %d", resultado, esperado)
	}

}
func TestDivisionbyZero(t *testing.T) {

	//Datos
	num1 := 10
	den := 0

	//test

	_, err := Dividir(num1, den)

	//validate
	assert.NotNil(t, err)

}
