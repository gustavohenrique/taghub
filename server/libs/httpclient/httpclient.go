package httpclient

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
)

type Config struct {
	ContentType   string
	Accept        string
	UserAgent     string
	Username      string
	Password      string
	RequestedWith string
	BaseURL       string
	Token         string
	Timeout       time.Duration
}

type HTTPClient struct {
	config Config
}

func New(config Config) *HTTPClient {
	if config.ContentType == "" {
		config.ContentType = "application/json"
	}
	if config.Accept == "" {
		config.Accept = "application/json"
	}
	if config.UserAgent == "" {
		config.UserAgent = "taghub"
	}
	return &HTTPClient{config}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (h *HTTPClient) doRequest(url string, req *fasthttp.Request) ([]byte, int, error) {
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if h.config.Token != "" {
		req.Header.Set(fasthttp.HeaderAuthorization, "Bearer "+h.config.Token)
	} else if h.config.Username != "" || h.config.Password != "" {
		req.Header.Set(fasthttp.HeaderAuthorization, "Basic "+basicAuth(h.config.Username, h.config.Password))
	}

	if h.config.RequestedWith != "" {
		req.Header.Set("X-Requested-With", h.config.RequestedWith)
	}

	if h.config.BaseURL != "" {
		url = fmt.Sprintf("%s%s", h.config.BaseURL, url)
	}
	req.SetRequestURI(url)
	req.Header.SetContentType(h.config.ContentType)
	req.Header.Set(fasthttp.HeaderAccept, h.config.Accept)
	req.Header.SetUserAgent(h.config.UserAgent)
	timeout := h.config.Timeout * time.Second
	err := fasthttp.DoTimeout(req, resp, timeout)
	if err != nil {
		log.Println(err.Error())
		return []byte(""), 0, err
	}

	statusCode := resp.StatusCode()
	bodyBytes := resp.Body()
	return bodyBytes, statusCode, nil
}

func (h *HTTPClient) GET(url string) ([]byte, int, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	return h.doRequest(url, req)
}

func (h *HTTPClient) POST(url string, data []byte) ([]byte, int, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("POST")
	req.SetBody(data)
	return h.doRequest(url, req)
}

func (h *HTTPClient) PUT(url string, data []byte) ([]byte, int, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("PUT")
	req.SetBody(data)
	return h.doRequest(url, req)
}

func (h *HTTPClient) DELETE(url string) ([]byte, int, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("DELETE")
	return h.doRequest(url, req)
}

func (h *HTTPClient) PATCH(url string, data []byte) ([]byte, int, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("PATCH")
	req.SetBody(data)
	return h.doRequest(url, req)
}

func (h *HTTPClient) Failed(err error, statusCode int) bool {
	return err != nil || statusCode >= 300
}

func SaveRequest(req *http.Request, filename string) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	if filename != "" {
		return ioutil.WriteFile(filename, body, 0644)
	}
	fmt.Println(string(body))
	return nil
}
