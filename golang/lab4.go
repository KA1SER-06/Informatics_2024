package lab4

import (
"fmt"
"math"
)

func calculateY(a, x float64) float64 {
xSquaredMinus1 := x*x - 1
part1 := math.Pow(a, xSquaredMinus1) 
part2 := math.Log(xSquaredMinus1) 
part3 := math.Cbrt(xSquaredMinus1) 

y := part1 - part2 + part3
return y
}

func SolveTaskA(a, b, xStart, xEnd, step float64) []float64 {
	results := []float64{}
	for x := xStart; x <= xEnd; x += step {
		y := CalculateY(a, b, x)
		results = append(results, y)
	}
	return results
}

func SolveTaskB(a, b float64, xValues []float64) []float64 {
	results := []float64{}
	for _, x := range xValues {
		y := CalculateY(a, b, x)
		results = append(results, y)
	}
	return results
}

func startLab4() {
a := 1.6

fmt.Println("Задача A:")
	xStartA := 1.2
	xEndA := 3.7
	step := 0.5

	for x := xStartA; x <= xEndA; x += step {
		y := calculateY(a, b, x)
		fmt.Printf("x = %.2f, y = %.5f\n", x, y)
	}


	fmt.Println("Задача B:")
	xValuesB := []float64{1.28, 1.36, 2.47, 3.68, 4.56}


	for _, x := range xValuesB {
		y := calculateY(a, b, x)
		fmt.Printf("x = %.2f, y = %.5f\n", x, y)
	}
}