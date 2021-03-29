package main

import (
	"math"
)

func funcion(x float64, n float64) float64 {
	if x == 0 {
		return n
	}
	return (1 - (math.Pow((1 + x), (-1 * n)))) / x
}

func error(A float64, B float64) float64 {
	return math.Abs(B-A) / math.Abs(B)
}
func bisec(y float64, n float64) float64 {
	var parada float64 = math.Pow(10, -6)
	var a float64 = 0
	var b float64 = 1
	var Fa float64 = funcion(a, n)
	var Fb float64 = funcion(b, n)
	if error(Fa, y) < parada {
		return a
	}
	if error(Fb, y) < parada {
		return b
	}
	var m float64 = 0.5
	var Fm float64 = funcion(m, n)
	for error(Fm, y) > parada {
		if Fm > y {
			a = m
		} else {
			b = m
		}
		m = (a + b) / 2
		Fm = funcion(m, n)
	}
	return m
}

func Anualizar(m float64) float64 {
	return math.Pow(1+m, 12) - 1
}

func cae(total int, cuota int, periodo int) float64 {
	m := bisec(float64(total/cuota), float64(periodo))
	return Anualizar(m)
}
