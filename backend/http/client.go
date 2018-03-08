package http

import (
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {
	if response, err := http.Get(url); err != nil {
		return []byte{}, err
	} else {
		defer response.Body.Close()
		if content, err := ioutil.ReadAll(response.Body); err != nil {
			return []byte{}, err
		} else {
			return content, nil
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
