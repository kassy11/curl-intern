package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func Contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

// -vオプションでリクエスト・レスポンスのヘッダーを表示
// TODO: できればHTTPSのときのSSL証明書の表示の追加
func DumpRequest(req *http.Request, resp *http.Response) {
	reqDump, _ := httputil.DumpRequest(req, false)
	respDump, _ := httputil.DumpResponse(resp, false)
	fmt.Printf("%s", string(reqDump))
	fmt.Printf("%s", string(respDump))
}
