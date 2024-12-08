package main

/*
Bibliotecas usadas
fmt: para entrada y salida de datos (similar a System.out.println y Scanner en Java)
math/rand: para generar números aleatorios (equivalente a Random en Java)
strings: para manipulación de cadenas (similar a métodos de String en Java)
time: para inicializar el generador de números aleatorios con una semilla basada en el tiempo actual
*/

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main(){
	fmt.Println("*** Sistema Generador de ID Único ***")

    // Crear un generador de números aleatorios con una semilla basada en el tiempo actual
	/*
	rand.New(): crea un generador de números pseudoaleatorios basado en esa fuente
	rand.NewSource(): utiliza ese valor como semilla para inicializar la fuente del generador de números aleatorios
	time.Now().UnixNano(): obtiene la cantidad de nanosegundos desde el Epoch Time
	*/
	generador := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Pedir datos al usuario
	var nombre, apellido, añoNacimiento string
	fmt.Print("Cúal es tu nombre? ")
	fmt.Scanln(&nombre)
	fmt.Print("Cúal es tu apellido? ")
	fmt.Scanln(&apellido)
	fmt.Print("Cúal es tu año de nacimiento? ")
	fmt.Scanln(&añoNacimiento)

    // Normalizar los valores
	nombre2 := strings.ToUpper(strings.TrimSpace(nombre))[:2]
	apellido2 := strings.ToUpper(strings.TrimSpace(apellido))[:2]
	añoNacimiento2 := strings.ToUpper(strings.TrimSpace(añoNacimiento))[2:]

	// Generar el valor aleatorio entrre 1 y 9999
	valorAleatorio := generador.Intn(9999) + 1

	// Formato de 4 dígitos
	valorAleatorioFormato := fmt.Sprintf("%04d", valorAleatorio)

	// Generar el ID Único
	idUnico := nombre2 + apellido2 + añoNacimiento2 + valorAleatorioFormato

	// Imprimir el ID único
	fmt.Printf(`
	Hola %s,
	Tu nuevo número de identificación (ID) generado por el sistema es:
	%s
	Felicidades!
	`, nombre, idUnico)
}
