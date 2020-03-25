package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"server/libs/httpvalidator"
	"server/pkg"
	"server/src/containers/service"
	"server/src/handlers/repo"
)

type HttpServer struct {
	server   *echo.Echo
	services *service.ServiceContainer
}

func NewServer(services *service.ServiceContainer) *HttpServer {
	server := &HttpServer{
		server:   echo.New(),
		services: services,
	}
	server.setDefaultConfiguration()
	server.addMiddlewares()
	server.addRoutesTo(services)
	return server
}

func (s *HttpServer) GetEchoServer() *echo.Echo {
	return s.server
}

func hi(c echo.Context) error {
	return c.String(http.StatusOK, pkg.DATETIME+" "+pkg.VERSION)
}

func (s *HttpServer) addRoutesTo(services *service.ServiceContainer) {
	e := s.server

	e.GET("/", hi)

	api := e.Group("/api")
	api.Use(VerifyAuth(s.services))

	repo.NewRepoHandler(services).AddRoutesTo(api)
}

func (s *HttpServer) setDefaultConfiguration() {
	e := s.server
	e.Validator = httpvalidator.New()
	e.HideBanner = true
	e.Debug = true
}

func (s *HttpServer) addMiddlewares() {
	e := s.server
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodOptions, http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{
			"Authorization",
			"X-Requested-With",
			"X-Request-ID",
			"Content-Type",
			"Accept",
			"User-Agent",
			"X-Amzn-Trace-Id",
			"X-Forwarded-For",
			"X-Forwarded-Port",
			"X-Real-Ip",
		},
	}))
	e.Use(middleware.BodyLimit("10M"))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
}

func (s *HttpServer) Start(port string) {
	e := s.server

	go func() {
		log.Fatal(e.Start(port))
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGQUIT)
	<-quit
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
