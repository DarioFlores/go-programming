package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// TODO ===> Correr con ambos
	// $ go run -race main.go
	// $ go run main.go
	var wg sync.WaitGroup
	var incremento int64
	gs := 100

	for i := 0; i < gs; i++ {
		go func() {
			wg.Add(1)
			atomic.AddInt64(&incremento, 1)
			fmt.Println(fmt.Sprintf("incremento-1: %d", atomic.LoadInt64(&incremento)))
			wg.Done()
		}()
		go func() {
			wg.Add(1)
			atomic.AddInt64(&incremento, 1)
			fmt.Println(fmt.Sprintf("incremento-2: %d", atomic.LoadInt64(&incremento)))
			wg.Done()
		}()
	}
	fmt.Println(fmt.Sprintf("CASI FINALIZA: %d", atomic.LoadInt64(&incremento)))
	fmt.Println(fmt.Sprintf("GOROUTINES: %d", runtime.NumGoroutine()))

	wg.Wait()
	fmt.Println(fmt.Sprintf("Valor final incremento: %d", incremento))
}
