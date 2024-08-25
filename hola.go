// con el comando en la consola go env nos permite saber donde esta el path principal de golang donde se tendra que colocar el proyecto para correrlo sin problemas, se puede dejar en otra carpeta pero no es recomendado ya que en el path principal tiene carpetas que usa golang en esto esta la carpeta bin que es donde dejara nuestra compilacion y el pkg que son paquetes de proyectos terceros que usaremos en nuestro proyecto
// la extension de archivo de go es .go asi tenemos que crear archivos de golang
// con go version podemos saber que version de golang tenemos
// podemos trabajar en playground de golang pero ten en cuenta que tiene limitaciones por ejemplo no podemos acceder al sistema operativo, tampoco podemos usar paquetes de terceros y no podemos crear un software dinamico esto quiere decir que no podemos crear aplicaciones donde el usuario necesite colocar datos como por ejemplo tomar datos de la consola
// con el simbolo // se crea comentarios en golang
package main // esto hace referencia a que este archivo o paquete se llamara main todos los paquetes main tienen una funcion principal llamada main que esta declarada mas adelante y al ejecutarse automaticamente ejecutara primero esa funcion

import "fmt" // con import podemos importar paquetes que necesite golang, el paquete fmt es un paquete principal de golang (no es un paquete tercero) y este paquete principalmente lo que hace es tomar e imprimir datos osea input y output de datos

func main() { // como todos los archivos main tiene una funcion llamada main, que lo que hace es ejecutar esta funcion ya que es la principal, para crear una funcion usamos la palabra reservada func
	fmt.Println("Hola Mundo") // cuando importamos el pauquete podemos usar el objeto fmt para llamar funciones en este caso usamos Println, esta funcion tiene como parametro un string que lo que hace es imprimir en consola el texto que pasamos en
}

// si queremos compilar internamente y correr el archivo, en la terminal podemos dar el comando go run nombredelarchivo.go, lo que hace esto es compilar (lo compila en la carpeta principal y en bin donde se instalo golang coloca la compilacion) y a la vez ejecutara el archivo mostrandonos hola mundo, pero no nos dejara la compilacion
// si queremos compilar y dejar el archivo para produccion podemos usar el comando go build nombredelarchivo.go, esto compilara el archivo y nos lo dejara en la carpeta del archivo compilado, si queremos ejecutar simplemente escribimos en consola hola.exe (esto solo sucede en windows) en linux seria ./hola
