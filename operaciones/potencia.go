package operaciones

func Potencia(base int, potencia int) int64 {
	resultado := int64(1)
	for rep := 0; rep < potencia; rep++ {
		resultado = resultado * int64(base)
	}
	return resultado
}
