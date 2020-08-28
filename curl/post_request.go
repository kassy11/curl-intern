package curl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"os"
	"github.com/kassy11/mycurl/utils"
)

func Post(client *http.Client, addr string, header bool, values url.Values) {

	req, err := http.NewRequest("POST", addr, strings.NewReader(values.Encode()))
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))
}
