package main

import "fmt"

func main() {

	ca := make(chan int) ////El int indica el tipo de dato que el canal va a estar enviando o recibiendo
	//ca es una variable

	go func() { //Necesitamos crear una goroutine que envie el valor mediante el canal asi main (que es quien lo recibe) lo recibe, por ser unbuffered channels
		ca <- 42 //Con esta linea decimos que le mandamos al canal llamado ca, el valor 42

	}()
	fmt.Println(<-ca) //acá indicamos que donde se pone <-ca es donde se deberia RECIBIR el valor

}
