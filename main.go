package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	var requestHeader bool
	flag.BoolVar(&requestHeader, "v", false, " Make the operation more talkative")
	var outputFile string
	flag.StringVar(&outputFile, "o", "", "--output <file> Write to file instead of stdout")
	var responseType string
	flag.StringVar(&responseType, "X", "GET", "--request <command> Specify request command to use")
	flag.Parse()
	
	if responseType == "GET"{
		get(flag.Arg(0))
	}else{
		post(flag.Arg(0))
	}

}

func get(url string){
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: tr,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	fmt.Println(string(body))
}

func post(url string){

}

