package main

import (
	"fmt"
	"sync"

	getapi "github.com/HemanthCU/go-big-data-project/getapi"
)

func main() {
	var val float64
	var wg sync.WaitGroup
	var url string
	fmt.Println("Enter a value greater than 4 to see more")
	fmt.Scanln(&val)
	if val > 4 {
		wg.Add(1)
		go func() {
			url = "http://api.worldbank.org/v2/countries/BGD/indicators/NY.GDP.MKTP.KD.ZG?per_page=11&date=2000:2010&format=json"
			data := getapi.Getapi(url)
			fmt.Println(data)
			wg.Done()
		}()
		fmt.Println("main")
		wg.Wait()
	}

	fmt.Println("all code complete")
}
