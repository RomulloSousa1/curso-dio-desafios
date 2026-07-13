package main

import "fmt"

const ebulicaoK float64 = 373.15

func main() {
	var tempK float64 = ebulicaoK
	var tempC float64 = (tempK - 273)
	fmt.Println(tempC)
}
