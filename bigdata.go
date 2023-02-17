package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	//"time"

	getapi "github.com/HemanthCU/go-big-data-project/getapi"
)

type WorldBankStatInfo struct {
	Page        int    `json:"page"`
	Pages       int    `json:"pages"`
	Per_page    int    `json:"per_page"`
	Total       int    `json:"total"`
	Sourceid    string `json:"sourceid"`
	Sourcename  string `json:"sourcename"`
	Lastupdated string `json:"lastupdated"`
}

type WorldBankStatListItem struct {
	Indicator struct {
		Id    string `json:"id"`
		Value string `json:"value"`
	} `json:"indicator"`
	Country struct {
		Id    string `json:"id"`
		Value string `json:"value"`
	} `json:"country"`
	Countryiso3code string  `json:"countryiso3code"`
	Date            string  `json:"date"`
	Value           float64 `json:"value"`
	Unit            string  `json:"unit"`
	Obs_status      string  `json:"obs_status"`
	Deci            int     `json:"decimal"`
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
			url = "http://api.worldbank.org/v2/countries/BGD/indicators/NY.GDP.MKTP.KD.ZG?per_page=11&date=2000:2010&format=json"
			data := getapi.Getapi(url)
			var statinfo *WorldBankStatInfo
			var statlist []WorldBankStatListItem
			var tmp []json.RawMessage
			json.Unmarshal(data, &tmp)
			err := json.Unmarshal(tmp[0], &statinfo)
			if err != nil {
				log.Fatalln(err)
			}
			err = json.Unmarshal(tmp[1], &statlist)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(statinfo.Sourceid)
			fmt.Println(statlist[0].Country.Value)
			wg.Done()
		}()
		fmt.Println("main")
		wg.Wait()
	}

	/*fmt.Println("Starting infinite loop")
	url = "http://api.worldbank.org/v2/countries/BGD/indicators/NY.GDP.MKTP.KD.ZG?per_page=11&date=2000:2010&format=json"
	for {
		go func() {
			fmt.Println(getapi.Getapi(url))
		}()
		time.Sleep(time.Second * 10)
	}*/

	file, err := os.Open("nytkey.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	freader := bufio.NewReader(file)
	buf := make([]byte, 32)
	_, err = freader.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	nytkey := string(buf)
	nyturl := "https://api.nytimes.com/svc/topstories/v2/world.json?api-key=" + nytkey
	nytdata := getapi.Getapi(nyturl)
	if val > 4 {
		fmt.Println(string(nytdata))
	}
}
