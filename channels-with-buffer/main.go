package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	c <- 49

	fmt.Println(<-c)
}
