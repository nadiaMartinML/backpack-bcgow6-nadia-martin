package divide

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivideGood(t *testing.T) {
	//
	numA := 6
	numB := 2
	expectResult := 3
	//
	realResult, _ := Divide(numA, numB)
	//
	// if realResult != expectResult {
	// 	t.Error("El resultado esperado es:", expectResult, " y el resultado real es: ", realResult, ". Son distintos, test falló")
	// }
	assert.Equal(t, expectResult, realResult)
}

func TestDivideBad(t *testing.T) {
	//
	numA := 6
	numB := 0
	//expectResult := fmt.Errorf("El denominador no puede ser 0")
	//
	_, error := Divide(numA, numB)
	//
	// if error == expectResult {
	// 	t.Error("El denominador no puede ser 0")
	// }
	assert.NotNil(t, error)
}
