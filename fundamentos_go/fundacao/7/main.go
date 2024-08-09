package main

type numero interface {
	int | float64
}

func Soma[T numero](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	salarios := map[string]int{"victor": 2000, "Julia": 8443, "Pandora": 64521}
	a := Soma(salarios)
	println(a)
}
