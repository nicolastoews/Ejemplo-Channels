package main

import (
	"fmt"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan bool)

	go enviar(par, impar, salir)

	recibir(par, impar, salir)

	fmt.Println("Finalizar.")
}

// send channel
func enviar(par, impar chan<- int, salir chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(salir)//cerrar el canal!
}

// receive channel
func recibir(par, impar <-chan int, salir <-chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("El valor recibido del canal par:", v)
		case v := <-impar:
			fmt.Println("El valor recibido del canal impar:", v)
		case i, ok := <-salir: //en i vamos a almacenar lo del canal salir y el ok nos va a dar verdadero o falso dependiendo de si hay o no hay valor en el canal salir
			if !ok {//!ok es igual a if ok == false
				fmt.Println("desde comma ok", i)
				return
			} else {
				fmt.Println("desde comma ok", i)
			}
		}
	}
}
