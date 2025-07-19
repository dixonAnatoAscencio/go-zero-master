package variables

import "fmt"

func MuestroEnteros() {
	//las variables no se definen como nulas, se definen con el valor minimo de su tipo de dato
	var intComun int       // declaracion de variable
	intDeOtroPaquete := 42 // declaracion de variable y asignacion de valor

	fmt.Println("intComun", intComun)
	fmt.Println("intDeOtroPaquete", intDeOtroPaquete)
}
