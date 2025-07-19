package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32

func RestoVariables() {
	fmt.Println("Resto")
	Nombre = "Dixon"
	Estado = true
	Sueldo = 1500.50

	fmt.Println("Nombre", Nombre)
	fmt.Println("Estado", Estado)
	fmt.Println("Sueldo", Sueldo)
	fmt.Println("Fecha", time.Now())
}

func ConviertoaTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
