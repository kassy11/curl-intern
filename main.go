package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"github.com/kassy11/mycurl/curl"
	"github.com/kassy11/mycurl/utils"
)

func main() {
	// -hオプション用文言
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s [options...] <url>\n", os.Args[0])
		flag.PrintDefaults()
	}

	// オプションの設定
	var showHeader bool
	flag.BoolVar(&showHeader, "v", false, "-v, --verbose  Make the operation more talkative")
	var outputFile string
	flag.StringVar(&outputFile, "o", "", "-o, --output <file>  Write to file instead of stdout")
	var postValues string
	flag.StringVar(&postValues, "d", "", "-d, --data <data>  HTTP POST data")
	var requestType string
	flag.StringVar(&requestType, "X", "GET", "-X, --request <command>  Specify request command to use")
	flag.Parse()

	// URLの指定がない時
	if len(flag.Args()) <= 0 {
		fmt.Printf("%s: no URL specified!\n", os.Args[0])
		fmt.Printf("%s: try '%s --help' or '%s --manual' for more information\n", os.Args[0], os.Args[0], os.Args[0])
		os.Exit(1)
	}
	addr := flag.Arg(0)

	// -dオプションのみでqueryがないとき
	if utils.Contains(os.Args, "-d") && postValues == "" {
		fmt.Printf("%s: option -d: requires parameter\n", os.Args[0])
		fmt.Printf("%s: try '%s --help' or '%s --manual' for more information\n", os.Args[0], os.Args[0], os.Args[0])
		os.Exit(1)
	}

	// クライアントを作成
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

	// postValuesをsplitしてurl.Values{}に格納
	values := url.Values{}
	if postValues != "" {
		splitEach := strings.Split(postValues, "&")
		for _, v := range splitEach {
			splitKeyVaue := strings.Split(v, "=")
			if len(splitKeyVaue) == 2 {
				values.Add(splitKeyVaue[0], splitKeyVaue[1])
			} else {
				values.Add(splitKeyVaue[0], "")
			}
		}
		fmt.Println(values.Encode())
	}

	// GETかPOSTで分岐
	if requestType == "GET" {
		curl.Get(client, addr, showHeader, outputFile)
	} else if requestType == "POST" {
		curl.Post(client, addr, showHeader, values)
	} else {
		fmt.Printf("%s: requestType is not correct!\n", os.Args[0])
		fmt.Printf("%s: try 'kcurl --help' or 'kcurl --manual' for more information\n", os.Args[0])
		os.Exit(1)
	}
}
