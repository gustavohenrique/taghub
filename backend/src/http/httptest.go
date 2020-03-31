package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"backend/src/containers/service"
	"backend/src/domain"
)

type HTTPServer struct {
	services *service.ServiceContainer
}

func GenerateValidTokenForTesting() string {
	return "token"
}

func With(services *service.ServiceContainer) HTTPServer {
	return HTTPServer{services}
}

func (h HTTPServer) ServeHTTP(req *http.Request) (*http.Response, domain.Response) {
	backend := NewServer(h.services)
	resp := httptest.NewRecorder()
	backend.GetEchoServer().ServeHTTP(resp, req)
	result := resp.Result()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	var response domain.Response
	json.Unmarshal(body, &response)
	return result, response
}
