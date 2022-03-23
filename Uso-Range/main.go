package main

import "fmt"

func main() {
	c := make(chan int)

	//enviar
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) // Es necesario cerrar el canal para que luego no tire error el range
	}()

	//recibir
	for v := range c { //OJO acá porque el range va a querer seguir sacando datos del canal y va a tirar error porque necesitamos CERRAR EL CANAL
		fmt.Println(v)
	}

	fmt.Println("Finalizando")
}
