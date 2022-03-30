package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// c <- 49 // fatal error: all goroutines are asleep - deadlock!
	go func() {
		// time.Sleep(10 * time.Second) // Si lo pongo el Println espera hasta que llegue un resultado
		c <- 49
	}()

	fmt.Println(<-c)
}
