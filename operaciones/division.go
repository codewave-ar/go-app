package operaciones

import "fmt"

func Division(dividendo int, divisor int) (float32, error) {
	if esCero := (divisor == 0); esCero {
		return 0, fmt.Errorf("No se puede dividir por 0")
	} else {
		return float32(dividendo) / float32(divisor), nil
	}
}
