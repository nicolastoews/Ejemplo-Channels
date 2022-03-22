package main

import "fmt"

func main() {

	ca := make(chan int)////El int indica el tipo de dato que el canal va a estar enviando o recibiendo
	//ca es una variable
	ca <- 42 //Con esta linea decimos que le mandamos al canal llamado ca, el valor 42

	fmt.Println(<-ca)

}
