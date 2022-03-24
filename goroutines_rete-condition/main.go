package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// TODO ===> Correr con ambos
	// $ go run -race main.go
	// $ go run main.go
	var wg sync.WaitGroup
	incremento := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incremento
			runtime.Gosched()
			v++
			fmt.Println(fmt.Sprintf("incremento: %d -- v: %d", incremento, v))
			incremento = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(fmt.Sprintf("Valor final incremento: %d", incremento))
}
