package main

import (
	"fmt"

	"rsc.io/quote"
)

var firstName2, lastName2, age2 = "alex", "roel", 27 // las variables podemos declararlas tanto fuera como adentro de la funcion, incluso como es una variable con scope global podemos usarlos en varias funciones del mismo archivo

func main() {
	fmt.Println(firstName2, lastName2, age2)
	fmt.Println(quote.Hello())
	// declaracion de variables
	//var firstName string // una forma de declarar una variable es usando primero la palabra var para declarar la variable luego sigue el nombre de la variable y despues el tipo de dato en este caso es una cadena por eso string, en go si declaramos una variable sin usarla golang nos dara un error al compilar
	//var firstName, lastName string // si necesitamos declarar varias variables del mismo tipo simplemente usamos var para declarar las variables y despues los variables separadas con coma y al final el tipo de dato en este caso seria string por que todas las variables son cadenas
	//var age int                    // como vemos esta variable esta declarada como integer y es otro tipo de dato por eso tenemos que declaralo en otra linea de codigo

	// hay una manera de declarar varias variables de diferentes tipos de la siguiente forma
	/*var (
		firstName, lastName string
		age                 int
	)*/ // como vemos simplemente usamos la palabra clave var y entre parentesis colocamos las variables, la variable tiene que ir con sus respectivos tipos de datos separado con comas y si tiene un diferente tipo de dato se deja un salto de linea para definir el siguiente tipo de dato con sus respectivas varaibles y tipos de datos

	// para definir las variables tenemos que hacerlo aparte de la declaracion y simplemente la igualamos con el simbolo =
	/*firstName = "Alex"
	lastName = "Roel"
	age = 27*/
	// aunque asignemos variables aun asi si no usamos las variables para alguna operacion, al compilar o correr golang nos dara un error

	// si queremos declarar las variables y definirlas tambien podemos hacerlo de la siguiente manera
	/*var (
		firstName string = "Alex"
		lastName  string = "Roel"
		age       int    = 27
	)*/ // como vemos para declarar y definir simplemente usamos la misma nomenclatura que usamos para declarar varias variables de varios tipos pero la diferencia es que todas por obligacion tiene que tener salto de linea con su respectivo tipo de dato y dato definiendo la variable

	/*var (
		firstName = "Alex"
		lastName  = "Roel"
		age       = 27
	)*/ // una cosa a tener en cuenta es que en golang cuando inicializamos no es necesdario colocar el tipo de dato por que automaticamente lo interpreta, se vuelve un tipo dinamico pero recuerda que a un siendo tipo dinamico golang lo interpreta como con su respectivo tipo de dato si es string sera string, si solo la declaramos ahi si necesitamos colocarle el tipo de dato

	// Si tenemos todas las variables definida con datos podemos incluso simplificar mas el codigo y dejarlo en una sola linea sin parentesis de la siguiente manera
	var firstName, lastName, age = "Alex", "Roel", 27 // como vemos usamos aun var y dividimos cada variable con coma despues le damos igual y cada dato pertenece a una variable segun su orden, el primer dato va a ir a la primera variable y asi sucesivamente para cada dato y variable, ademas asigna tipo de dato como vemos los dos primeros son de tipo string por que su dato es un string osea que estas variables son de tipo de dato string y seran string y su modificacion sera string a menos que hagamos un cast esto sucede igual tambien con la variable int

	// para que no nos de error por no usar las variables simplemente podemos imprimirla para que no nos de error
	fmt.Println(firstName, lastName, age) // como vemos en println podemos imprimir varias variables simplemente colocando varias variables en el parametro del metodo println separados con comas

	//tambien hay una manera de inicializar las variables sin la palabra reservada var y ademas tambien podemos casi todo lo que vimos y es usando el simbolo :=
	firstNameDaniel, laslastNameDaniel, ageDaniel := "Daniel", "Viera", 34 // con el simbolo := estamos declarando variables nuevas e inicializandolas, por lo tanto no se puede usar el simbolo := para modificar y estas tipo de declaracion e inicializacion solo se puede usar dentro de las funciones no fuera, para hacerlo fuera toca usar var
	ageDiana := 35
	ageDiana = 34
	fmt.Println(firstNameDaniel, laslastNameDaniel, ageDaniel)
	fmt.Println(ageDiana)
}
