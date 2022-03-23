package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	fanin := make(chan int)

	go enviar(par, impar)

	go recibir(par, impar, fanin)

	for v := range fanin {
		fmt.Println(v)//va a imprimir todos los valores de par e impar
	}

	fmt.Println("Finalizando.")
}

// send channel
func enviar(par, impar chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {//asigna valores a par e impar
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)//cierra el canal
	close(impar)
}

// receive channel
func recibir(par, impar <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)//definimos que habrán 2 goroutines

	go func() {
		for v := range par {
			fanin <- v//le pasamos a fanin todos los valores pares
		}
		wg.Done()//le decimos al contador que esta goroutine ya terminó
	}()

	go func() {
		for v := range impar {
			fanin <- v//le pasamos a fanin todos los valores impares
		}
		wg.Done()//le decimos al contador que esta goroutine ya terminó
	}()

	wg.Wait()
	close(fanin)//cerramos el canal fanin
}
