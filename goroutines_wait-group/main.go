package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println(fmt.Sprintf("Inicio - Numero de CPU: %d", runtime.NumCPU()))
	fmt.Println(fmt.Sprintf("Inicio - Numero de GoRoutines: %d", runtime.NumGoroutine()))

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("Hola que hace?")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hola que hace?x2")
		wg.Done()
	}()

	fmt.Println(fmt.Sprintf("Medio - Numero de CPU: %d", runtime.NumCPU()))
	fmt.Println(fmt.Sprintf("Medio - Numero de GoRoutines: %d", runtime.NumGoroutine()))

	wg.Wait()

	fmt.Println(fmt.Sprintf("Final - Numero de CPU: %d", runtime.NumCPU()))
	fmt.Println(fmt.Sprintf("Final - Numero de GoRoutines: %d", runtime.NumGoroutine()))
}
