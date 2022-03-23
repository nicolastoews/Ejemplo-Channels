package main

import "fmt"

func main() {

	c := make(chan int)    //Canal bidireccional
	cs := make(chan<- int) //chan<- indica que es SEND ONLY
	cr := make(<-chan int) //chan<- indica que es RECIEVE ONLY

	fmt.Println("--------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("c\t%T\n", cs)
	fmt.Printf("c\t%T\n", cr)

	// c = canal bidireccional | cs = canal send only | cr = canal read only

	fmt.Println("--------")
	//No asigna, no funciona	
	c = cs  //específico a general
	c = cr  //específico a general
	cs = cr //específico a específico

	fmt.Println("--------")
	//Asigna, si funciona
	cr = c //general a específico
	cs = c //general a específico

	fmt.Println("--------")
	//No funciona la conversión
	fmt.Printf("c\t%T\n", (chan int)(cr)) //específico a general
	fmt.Printf("c\t%T\n", (chan int)(cs)) //específico a general

	fmt.Println("--------")
	//Si funciona la conversión
	fmt.Printf("c\t%T\n", (chan int)(c)) //general a específico
	fmt.Printf("c\t%T\n", (chan int)(c)) //general a específico

}
