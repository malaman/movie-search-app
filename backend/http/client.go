package http

import (
	"io/ioutil"
	"net/http"

	"github.com/valyala/fasthttp"
)

type HTTPClient interface {
	Get(url string) ([]byte, error)
	Post(url string, body []byte) (string, error)
}

type Client struct{}

func (client *Client) Get(url string) ([]byte, error) {
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

func (client *Client) Post(url string, body []byte) (string, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.SetBody(body)

	resp := fasthttp.AcquireResponse()
	fastHTTPClient := &fasthttp.Client{}
	if err := fastHTTPClient.Do(req, resp); err != nil {
		return "", err
	}
	bodyBytes := resp.Body()
	return string(bodyBytes), nil
}
