package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func main() {
	// -hオプション用文言
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s [options...] <url>\n", os.Args[0])
		flag.PrintDefaults()
	}

	var requestHeader bool
	flag.BoolVar(&requestHeader, "v", false, " Make the operation more talkative")
	var outputFile string
	flag.StringVar(&outputFile, "o", "default", "--output <file> Write to file instead of stdout")
	var requestType string
	flag.StringVar(&requestType, "X", "GET", "--request <command> Specify request command to use")
	flag.Parse()

	fmt.Println(flag.Args(), requestHeader, outputFile, requestType)

	// URLの指定がない時
	if len(flag.Args())<=0{
		fmt.Printf("%s: no URL specified!\n", os.Args[0])
		fmt.Printf("%s: try '%s --help' or '%s --manual' for more information\n", os.Args[0], os.Args[0], os.Args[0])
		os.Exit(1)
	}

	if requestType == "GET"{
		get(flag.Arg(0), requestHeader, outputFile)
	}else if requestType == "POST"{
		fmt.Println("POSTリクエスト")
	}else{
		fmt.Printf("%s: requestType is not correct!\n", os.Args[0])
		fmt.Printf("%s: try 'kcurl --help' or 'kcurl --manual' for more information\n", os.Args[0])
		os.Exit(1)
	}

}

func get(url string, requestHeader bool, filename string){
	// レスポンスを作成
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
		os.Exit(1)
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)

	// リクエストを送信
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// レスポンスを受信して表示
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// -vオプションでファイル名を指定した時
	// TODO: ここの条件が気持ち悪いので直したい, ここのエラー処理
	if filename == "default"{
		fmt.Printf("%s: option -o: requires parameter\n", os.Args[0])
		fmt.Printf("%s: try '%s --help' or '%s --manual' for more information\n", os.Args[0], os.Args[0], os.Args[0])
		os.Exit(1)
	}else{
		fp, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer fp.Close()

		fp.WriteString(string(responseBody))
	}

	// -vオプションがあるときリクエスト内容を表示
	if requestHeader{
		dump, _ := httputil.DumpRequest(req, true)
		fmt.Println(string(dump))
	}

	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))
}

func post(url string){

}

