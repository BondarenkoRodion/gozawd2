package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
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

type Result struct {
	Riznytsia1 float64
	Riznytsia2 float64
}

var tmpl = template.Must(template.New("result").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Обчислення прибутку</title>
</head>
<body>
    <form action="/" method="POST">
        <label>Pc: <input type="text" name="Pc"></label><br>
        <label>s1 (до покращення): <input type="text" name="s1"></label><br>
        <label>s2 (після покращення): <input type="text" name="s2"></label><br>
        <label>B: <input type="text" name="B"></label><br>
        <input type="submit" value="Обчислити">
    </form>
    {{if .}}
    <h2>Результати</h2>
    <p>Прибуток до покращення: {{.Riznytsia1}}</p>
    <p>Прибуток після покращення: {{.Riznytsia2}}</p>
    {{end}}
</body>
</html>
`))

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		Pc := parseFloat(r.FormValue("Pc"))
		s1 := parseFloat(r.FormValue("s1"))
		s2 := parseFloat(r.FormValue("s2"))
		B := parseFloat(r.FormValue("B"))

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

		tmpl.Execute(w, Result{Riznytsia1, Riznytsia2})
		return
	}
	tmpl.Execute(w, nil)
}

func parseFloat(value string) float64 {
	var result float64
	fmt.Sscanf(value, "%f", &result)
	return result
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Сервер запущено на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
