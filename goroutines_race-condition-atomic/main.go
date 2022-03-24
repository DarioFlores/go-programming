package main

import (
	"fmt"
	"sync"
)

func main() {
	// TODO ===> Correr con ambos
	// $ go run -race main.go
	// $ go run main.go
	var wg sync.WaitGroup
	var mutex sync.Mutex
	incremento := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock()
			v := incremento
			v++
			fmt.Println(fmt.Sprintf("incremento: %d -- v: %d", incremento, v))
			incremento = v
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(fmt.Sprintf("Valor final incremento: %d", incremento))
}
