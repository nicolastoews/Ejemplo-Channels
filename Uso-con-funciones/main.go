package main

import "fmt"

func main() {
	c := make(chan int)

	//enviar
	go enviar(c)

	//recibir
	recibir(c)

	fmt.Println("Finalizando")
}

func enviar(c chan<- int) { //canal SEND ONLY como parámetro en la función
	c <- 42
}

func recibir(c <-chan int) { //canal RECIEVE ONLY como parámetro en la función
	fmt.Println(<-c)
}
