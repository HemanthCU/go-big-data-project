package main

import (
	"fmt"
	"sync"

	getapi "github.com/HemanthCU/go-big-data-project/getapi"
)

func main() {
	var val float64
	var wg sync.WaitGroup
	fmt.Println("Enter a value greater than 4 to see more")
	fmt.Scanln(&val)
	if val > 4 {
		wg.Add(1)
		go func() {
			getapi.Getapi(val)
			fmt.Println("goroutine")
			wg.Done()
		}()
		fmt.Println("main")
		wg.Wait()
	}
	fmt.Println("all code complete")
}
