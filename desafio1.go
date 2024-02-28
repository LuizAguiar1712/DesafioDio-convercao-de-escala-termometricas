package main

import "fmt"

const ebulicaoK = 373

func main() {

	kelvin := ebulicaoK
	celsius := kelvin - 273

	fmt.Println("O ponto de ebulição da água em K é = ", kelvin, "K")
	fmt.Println("O ponto de ebulição da água em °C é = ", celsius, "°C")

}
