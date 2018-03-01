package http

import (
	"github.com/valyala/fasthttp"
)

func GET(url string) (string, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		return "", err
	}
	bodyBytes := resp.Body()
	return string(bodyBytes), nil
}

func POST(url string, body []byte) (string, error) {
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
