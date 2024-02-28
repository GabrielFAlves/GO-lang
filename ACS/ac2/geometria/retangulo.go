package geometria

func CalcularArea(largura, altura float64) float64 {
    area := largura * altura
    return area
}

func CalcularPerimetro(largura, altura float64) float64 {
    perimetro := 2 * (largura + altura)
    return perimetro
}
