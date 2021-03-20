package goarea

import "math"

//Pi é doido
const Pi = 3.1416

//Circ é responsável por calcular a área de um Círculo
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect é responsável por calular a área de um Retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}

// chorume
