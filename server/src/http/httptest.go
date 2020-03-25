package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"server/src/containers/service"
	"server/src/domain"
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
	server := NewServer(h.services)
	resp := httptest.NewRecorder()
	server.GetEchoServer().ServeHTTP(resp, req)
	result := resp.Result()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	var response domain.Response
	json.Unmarshal(body, &response)
	return result, response
}
