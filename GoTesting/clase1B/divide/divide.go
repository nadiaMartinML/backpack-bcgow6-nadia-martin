package divide

import "fmt"

func Divide(numA, numB int) (int, error) {
	if numB == 0 {
		return 0, fmt.Errorf("El denominador no puede ser 0")
	}
	return numA / numB, nil
}
