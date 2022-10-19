package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtractGood(t *testing.T) {
	//
	numA := 9
	numB := 3
	expectResult := 6
	//
	realResult := Subtract(numA, numB)
	//

	//if expectResult != realResult {
	//	t.Errorf("El resultado esperado es: %v y el resultado real es: %v. Son distintos, test falló", expectResult, realResult)
	//}

	assert.Equal(t, expectResult, realResult, "El resultado esperado es: %v y el resultado real es: %v. Son distintos, test falló", expectResult, realResult)
}

func TestSubtractBad(t *testing.T) {
	//
	numA := 9
	numB := 3
	expectResult := 4
	//
	realResult := Subtract(numA, numB)
	//

	//if expectResult == realResult {
	//	t.Errorf("El resultado esperado es: %v y el resultado real es: %v. Son iguales, test falló", expectResult, realResult)
	//}

	assert.NotEqual(t, expectResult, realResult, "El resultado esperado es: %v y el resultado real es: %v. Son iguales, test falló", expectResult, realResult)
}
