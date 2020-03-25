package repo

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"server/libs/errors"
	"server/libs/filter"
	"server/libs/logger"
	"server/src/containers/service"
	"server/src/domain"
	"server/src/handlers"
)

type RepoHandler struct {
	repoService domain.RepoService
}

func NewRepoHandler(services *service.ServiceContainer) *RepoHandler {
	return &RepoHandler{
		repoService: services.RepoService,
	}
}

func (h *RepoHandler) AddRoutesTo(api *echo.Group) {
	route := api.Group("/repo")
	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.GET("/:id", h.ReadOne)
	route.GET("/sync", h.Sync)
	route.POST("/search", h.Search)
}

func (h *RepoHandler) Create(c echo.Context) error {
	var item domain.Repo
	if err := handlers.BindAndValidate(c, &item); err != nil {
		return err
	}
	err := h.repoService.Create(item)
	if err != nil {
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusCreated, domain.Response{Data: item})
}

func (h *RepoHandler) Update(c echo.Context) error {
	var item domain.Repo
	if err := handlers.BindAndValidate(c, &item); err != nil {
		return err
	}
	item.ID = c.Param("id")
	err := h.repoService.Update(item)
	if err != nil {
		log.Println(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusOK, domain.Response{Data: item})
}

func (h *RepoHandler) ReadOne(c echo.Context) error {
	id := c.Param("id")
	item := domain.Repo{ID: id}
	found, err := h.repoService.ReadOne(item)
	if err != nil {
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusOK, domain.Response{Data: found})
}

func (h *RepoHandler) Sync(c echo.Context) error {
	repos, err := h.repoService.Sync()
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusOK, domain.Response{Data: repos})
}

func (h *RepoHandler) Search(c echo.Context) error {
	var item filter.Request
	if err := handlers.BindAndValidate(c, &item); err != nil {
		return err
	}
	items, total, err := h.repoService.Search(item)
	if err != nil {
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	meta := domain.Meta{
		Page:    item.Pagination.Page,
		PerPage: item.Pagination.PerPage,
		Total:   total,
	}
	return c.JSON(http.StatusOK, domain.Response{Data: items, Meta: &meta})
}
