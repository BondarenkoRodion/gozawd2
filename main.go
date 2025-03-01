package main

import (
	"fmt"
	"math"
)

func integral(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := f(a) + f(b)
	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		if i%2 == 0 {
			sum += 2 * f(x)
		} else {
			sum += 4 * f(x)
		}
	}

	return sum * (h / 3)
}

func main() {
	var Pc, s1, s2, B float64

	fmt.Print("Введіть Pc: ")
	fmt.Scan(&Pc)

	fmt.Print("Введіть значення до покращення (s1): ")
	fmt.Scan(&s1)

	fmt.Print("Введіть значення після покращення (s2): ")
	fmt.Scan(&s2)

	fmt.Print("Введіть B: ")
	fmt.Scan(&B)

	dW1 := integral(func(p float64) float64 {
		return (1 / (s1 * math.Sqrt(2*math.Pi))) * math.Exp(-math.Pow(p-Pc, 2)/(2*math.Pow(s1, 2)))
	}, 4.75, 5.25, 1000000)

	W1 := Pc * 24 * dW1
	Pryb1 := W1 * B
	W21 := Pc * 24 * (1 - dW1)
	Sch1 := W21 * B
	Riznytsia1 := Pryb1 - Sch1

	dW2 := integral(func(p float64) float64 {
		return (1 / (s2 * math.Sqrt(2*math.Pi))) * math.Exp(-math.Pow(p-Pc, 2)/(2*math.Pow(s2, 2)))
	}, 4.75, 5.25, 1000000)

	W12 := Pc * 24 * dW2
	Pryb2 := W12 * B
	W22 := Pc * 24 * (1 - dW2)
	Sch2 := W22 * B
	Riznytsia2 := Pryb2 - Sch2

	fmt.Printf("Прибуток до покращення: %.2f, після покращення: %.2f\n", Riznytsia1, Riznytsia2)
}
