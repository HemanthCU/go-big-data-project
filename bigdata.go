package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	getapi "github.com/HemanthCU/go-big-data-project/getapi"
)

func getData(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	res := string(body)
	return res
}

func main() {
	var val float64
	var wg sync.WaitGroup
	var url string
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

	url = "http://api.worldbank.org/v2/countries/BGD/indicators/NY.GDP.MKTP.KD.ZG?per_page=11&date=2000:2010"
	data := getData(url)
	fmt.Println(data)

	fmt.Println("all code complete")
}
