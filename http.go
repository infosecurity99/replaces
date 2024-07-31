package main

import "fmt"


var URL string

type Protocal interface {
	request(URLEndpoint string)
	response() string
}

type HTTPProtocol struct {
	URL string
}

func (HTTP *HTTPProtocol) request(URLEndpoint string) {
	URL = fmt.Sprintf("%s%s", HTTP.URL, URLEndpoint)
	HTTP.URL = URL
}

func (http HTTPProtocol) response() string {
	return URL
}

func main() {

	HTTPP := HTTPProtocol{URL: "https:/www.golinuxcloud.com"}
	HTTPP.request("go-methods")
	res := HTTPP.response()

	fmt.Println(res)
	fmt.Println(HTTPP)
}
