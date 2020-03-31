package tag

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/libs/errors"
	"backend/libs/filter"
	"backend/libs/logger"
	"backend/src/containers/service"
	"backend/src/domain"
)

type TagHandler struct {
	tagService domain.TagService
}

func NewTagHandler(services *service.ServiceContainer) *TagHandler {
	return &TagHandler{
		tagService: services.TagService,
	}
}

func (h *TagHandler) AddRoutesTo(api *echo.Group) {
	route := api.Group("/tag")
	route.POST("/search", h.Search)
	route.GET("", h.ReadAll)
	route.GET("/:id/total", h.GetTotalReposByTag)
	route.DELETE("/:id", h.Remove)
	route.PUT("/:id", h.Update)
}

func (h *TagHandler) ReadAll(c echo.Context) error {
	items, err := h.tagService.ReadAll()
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusOK, domain.Response{Data: items})
}

func (h *TagHandler) GetTotalReposByTag(c echo.Context) error {
	tag := domain.Tag{ID: c.Param("id")}
	total, err := h.tagService.GetTotalReposByTag(tag)
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusOK, domain.Response{Data: total})
}

func (h *TagHandler) Update(c echo.Context) error {
	var item domain.Tag
	if err := c.Bind(&item); err != nil {
		return c.JSON(errors.Invalid, domain.Response{Err: err.Error()})
	}
	item.ID = c.Param("id")
	err := h.tagService.Update(item)
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusOK, domain.Response{Data: item})
}

func (h *TagHandler) Remove(c echo.Context) error {
	tag := domain.Tag{ID: c.Param("id")}
	err := h.tagService.Remove(tag)
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusNoContent, "")
}

func (h *TagHandler) Search(c echo.Context) error {
	var item filter.Request
	if err := c.Bind(&item); err != nil {
		return c.JSON(errors.Invalid, domain.Response{Err: err.Error()})
	}
	items, total, err := h.tagService.Search(item)
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	meta := domain.Meta{
		Page:    item.Pagination.Page,
		PerPage: item.Pagination.PerPage,
		Total:   total,
	}
	if c.QueryParam("total_repos") == "true" {
		itemsWithTotalRepos := []domain.Tag{}
		for _, tag := range items {
			total, _ := h.tagService.GetTotalReposByTag(tag)
			tag.TotalRepos = total
			itemsWithTotalRepos = append(itemsWithTotalRepos, tag)
		}
		items = itemsWithTotalRepos
	}
	return c.JSON(http.StatusOK, domain.Response{Data: items, Meta: &meta})
}
