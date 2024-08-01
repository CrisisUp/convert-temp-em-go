package main

import "fmt"

const ebulicaoakelvin = 373.0

func main() {
	tempK := ebulicaoakelvin
	tempC := tempK - 273.0
	fmt.Printf("A temperatura de ebulição da água em °K é %.1f, mas em °C é %.1f\n", tempK, tempC)
}
