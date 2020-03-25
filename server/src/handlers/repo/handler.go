package repo

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"server/libs/errors"
	"server/libs/filter"
	"server/src/containers/service"
	"server/src/domain"
	"server/src/handlers"
)

type RepoHandler struct {
	service domain.RepoService
}

func NewRepoHandler(services *service.ServiceContainer) *RepoHandler {
	return &RepoHandler{services.RepoService}
}

func (h *RepoHandler) AddRoutesTo(api *echo.Group) {
	route := api.Group("/repo")
	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.GET("/:id", h.ReadOne)
	route.POST("/search", h.Search)
}

func (h *RepoHandler) Create(c echo.Context) error {
	var item domain.Repo
	if err := handlers.BindAndValidate(c, &item); err != nil {
		return err
	}
	saved, err := h.service.Create(item)
	if err != nil {
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusCreated, domain.Response{Data: saved})
}

func (h *RepoHandler) Update(c echo.Context) error {
	var item domain.Repo
	if err := handlers.BindAndValidate(c, &item); err != nil {
		return err
	}
	item.ID = c.Param("id")
	saved, err := h.service.Update(item)
	if err != nil {
		log.Println(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusOK, domain.Response{Data: saved})
}

func (h *RepoHandler) ReadOne(c echo.Context) error {
	id := c.Param("id")
	item := domain.Repo{ID: id}
	found, err := h.service.ReadOne(item)
	if err != nil {
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusOK, domain.Response{Data: found})
}

func (h *RepoHandler) Search(c echo.Context) error {
	var item filter.Request
	if err := handlers.BindAndValidate(c, &item); err != nil {
		return err
	}
	items, total, err := h.service.Search(item)
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
