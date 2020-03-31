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

	"backend/pkg"
	"backend/src/containers/service"
	"backend/src/handlers/repo"
	"backend/src/handlers/tag"
)

type HttpServer struct {
	backend  *echo.Echo
	services *service.ServiceContainer
}

func NewServer(services *service.ServiceContainer) *HttpServer {
	backend := &HttpServer{
		backend:  echo.New(),
		services: services,
	}
	backend.setDefaultConfiguration()
	backend.addMiddlewares()
	backend.addRoutesTo(services)
	return backend
}

func (s *HttpServer) GetEchoServer() *echo.Echo {
	return s.backend
}

func hi(c echo.Context) error {
	return c.String(http.StatusOK, pkg.DATETIME+" "+pkg.VERSION)
}

func (s *HttpServer) addRoutesTo(services *service.ServiceContainer) {
	e := s.backend

	e.GET("/", hi)

	api := e.Group("/api")
	api.Use(VerifyAuth(s.services))

	repo.NewRepoHandler(services).AddRoutesTo(api)
	tag.NewTagHandler(services).AddRoutesTo(api)
}

func (s *HttpServer) setDefaultConfiguration() {
	e := s.backend
	e.HideBanner = true
	e.Debug = true
}

func (s *HttpServer) addMiddlewares() {
	e := s.backend
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
	e := s.backend

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
