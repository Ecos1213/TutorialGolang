// los paquetes terceros son paquetes que no estan incluidos en las libreria estandard de go como el paquete fmt, estos paquetes terceros han sido desarrollado y mantenidos por las comunidades de programadores de go o por otros desarrolladores de go
// para trabajar con paquetes de terceros (paquetes externos) tenemos que inicializar un manejador de modulos para nuestra aplicacion, entonces para inicializar un manejador de modulos para nuestra aplicacion vamos a ejcutar el siguiente comando dentro de nuestro proyecto, el comando es go mod, para el manejador de modulos y con init vamos a  inicializar el manejador de modulos despues de init sigue el nombre con el cual vamos a manejar el modulo de nuestro proyecto
// go mod init nombredemanejadordemodulos
// este comando creara un archivo llamado go.mod, el cual contiene el nombre del manejador del modulo (module holamundo) y una version de go (go 1.23.0)
// el archivo go.mod es un archivo que se utiliza para definir y gestionar los modulos y las dependencias de nuestro proyecto en go, tambien cuando empezemos a descargar y utilizar los archivos externos tambien se va a crear un archivo llamado go.sum, este archivo go.sum es un archivo se utiliza para registrar la suma de verificacion de los modulos y las dependencias de nuestro proyecto
// ahora vamos a utilizar el paquete cod, es un paquete externo de tercero para el lenguaje de programacion go que proporciona una serie de citas famosas como el hola mundo
// con el comando go get rutaderepositoriodelpaquete descargamos el paquete y agregarlo a nuestro proyecto, este puede estar en github como tambien puede estar en otra parte el paquete quite se descarga de la siguiente manera
// go get rsc.io/quote
// al hacer esto descargara el paquete y simeplemente agregara el modulo en el archivo go.mod y crear el archivo go.sum con el respectivo modulo

package main

import (
	"fmt"

	"rsc.io/quote"
) // en el import para agregar varios paquetes tenemos que usar el parentesis y dentro del parentesis agregamos todos los paquetes que vamos a usar, para agregar el paquete externo simplemente colocamos la ruta del paquete que es rsc.io/quote

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Hello()) // ahora podemos usar el objeto quote y llamar la funcion anidada Hello que tiene quote esto nos imprimira Hello world
}

//con go run hola.go ejecutamos nuestro programa en consola
