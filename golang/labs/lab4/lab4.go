package lab4

import (
  "fmt"
  "math"
)

func Calculate(x, a float64) float64 {
  y := math.Tan(math.Pow(math.Log10(a+x),3))/math.Pow((a+x)*(a+x),1.0/7.0)
  return y
}

func TaskA(b, Xn, Xk, delX float64) []float64 {
  var y []float64
  for x := Xn; x <= Xk; x += delX {
    y = append(y, Calculate(b, x))
  }
  return y
}

func TaskB(b float64, x [5]float64) []float64 {
  var y []float64
  for _, value := range x {
    y = append(y, Calculate(b,value))
  }
  return y
}

func RunLab4() {
  a := 2.0
  fmt.Println(TaskA(a, 1.08, 1.88, 0.16))
  var s = [5]float64{1.16, 1.35, 1.48,1.52, 1.96}
  fmt.Println(TaskB(a, s))
}
