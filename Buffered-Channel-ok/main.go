package main

import "fmt"

func main() {

	ca := make(chan int, 1) 
	//ca es una variable
	//El int indica el tipo de dato que el canal va a estar enviando o recibiendo
	//El 1 indica cuantos valoes se pueden quedar adentro del channel sin ser recibidos por nadie, osea el buffer

	ca <- 42 //Con esta linea decimos que le mandamos al canal llamado ca, el valor 42
	fmt.Println(<-ca) //acá indicamos que donde se pone <-ca es donde se deberia RECIBIR el valor

}
