### curl by Golang

### Usage
1. `go get github.com/kassy11/mycurl`
2. `cd $GOPATH/src/github.com/kassy11/mycurl`
3. `make`
4. GET `./mycurl -v -o test.txt "http://httpbin.org/get"`
5. POST `./mycurl -X POST -v -d 'key=value&key2=value2' -o test.txt "http://httpbin.org/post"`
