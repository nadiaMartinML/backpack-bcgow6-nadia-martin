package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortingGod(t *testing.T) {
	//
	num := []int{6, 2, 8, 3}
	expectResult := []int{2, 3, 6, 8}
	//
	realResult := SortingAsc(num)
	//
	// if expectResult[0] != realResult[0] || expectResult[1] != realResult[1] || expectResult[2] != realResult[2] {
	// 	t.Errorf("El resultado esperado es: %v y el resultado real es: %v. Son distintos, test falló", expectResult, realResult)
	// }
	assert.Equal(t, expectResult, realResult)

}

func TestSortingBad(t *testing.T) {
	//
	num := []int{6, 2, 8, 3}
	expectResult := []int{8, 6, 3, 2}
	//
	realResult := SortingAsc(num)
	//
	// if expectResult[0] == realResult[0] || expectResult[1] == realResult[1] || expectResult[2] == realResult[2] {
	// 	t.Errorf("El resultado esperado es: %v y el resultado real es: %v. Son iguales, test falló", expectResult, realResult)
	// }
	assert.NotEqual(t, expectResult, realResult)
}
