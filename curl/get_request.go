package curl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/kassy11/mycurl/utils"
)

func Get(client *http.Client, addr string, header bool, filename string) {

	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", "Bearer access-token")
	req.Header.Add("If-None-Match", `W/"wyzzy"`)

	// リクエストを送信
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// -vオプションがあるときリクエスト内容を表示
	if header {
		utils.DumpRequest(req, resp)
	}

	// レスポンスを受信して表示
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// TODO: ここのエラー処理直したい
	// -oオプションしかなくファイル名が指定されていない時はエラー表示
	if utils.Contains(os.Args, "-o") && filename == "" {
		fmt.Printf("%s: option -o: requires parameter\n", os.Args[0])
		fmt.Printf("%s: try '%s --help' or '%s --manual' for more information\n", os.Args[0], os.Args[0], os.Args[0])
		os.Exit(1)
	}

	// -oオプションがあってファイル名が指定されているときのみファイル書き込み
	if utils.Contains(os.Args, "-o") && filename != "" {
		fp, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer fp.Close()

		fp.WriteString(string(responseBody))
	}

	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))
}
