package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go agregar(c1)
	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("Finalizando")
}

func agregar(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 { //range bloquea canales si no se cierran!
		wg.Add(1)
		go func(v2 int) { //por cada valor que tengo en el c1, estoy lanzando una goroutine
			c2 <- trabajoConsumeTiempo(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func trabajoConsumeTiempo(n int) int {//Nada, consume tiempo y devuelve un dato aleatorio
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
