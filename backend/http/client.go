package http

import (	
	"github.com/valyala/fasthttp"
	"net/http"
	 "io/ioutil"
)

func Get(url string) (string, error) {
	if response, err := http.Get(url);	err != nil {
		return "", err
	} else {
		defer response.Body.Close()
		if content, err := ioutil.ReadAll(response.Body); err != nil {
			return "", err				
		} else {
			return string(content), nil
		}
	}
}

func Post(url string, body []byte) (string, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.SetBody(body)

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		return "", err
	}
	bodyBytes := resp.Body()
	return string(bodyBytes), nil
}
