package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	//enviar valores mediante goroutine y función
	go enviar(par, impar, salir)

	//recibir, como main cuenta como una routine más, no necesitamos mandarla a una aparte
	recibir(par, impar, salir)

	fmt.Println("Finalizando")

}

func enviar(p, i, s chan<- int) { //canales de send only tipo entero
	for j := 0; j < 100; j++ { //le mandamos del 0 al 99
		if j%2 == 0 { //compara si el valor de j es par o impar
			p <- j //envío el valor par de j al canal par (p)
		} else {
			i <- j //envío el valor impar de j al canal impar (i)
		}
	}
	s <- 0 //al canal salir (s) le mando el valor 0 para usarlo de salida en el select
}

func recibir(p, i, s <-chan int) {
	for { //bucle infinito
		select { //select tipo switch pero para canales
		case v := <-p: //asigna a v los valores del canal par (p)
			fmt.Println("Desde el canal par:", v)
		case v := <-i: //asigna a v los valores del canal impar (i)
			fmt.Println("Desde el canal impar:", v)
		case v := <-s: //asigna a v el valor del canal salir(s) el cual es 0
			fmt.Println("Desde el canal salir:", v)
			return //finaliza el bucle infinito
		}
	}
}
