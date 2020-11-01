package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
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
func DumpRequest(req *http.Request, resp *http.Response) {
	reqDump, _ := httputil.DumpRequest(req, false)
	respDump, _ := httputil.DumpResponse(resp, false)
	fmt.Printf("%s", string(reqDump))
	fmt.Printf("%s", string(respDump))
}

func ParseURL(postValues string) url.Values {
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
	}
	return values
}
