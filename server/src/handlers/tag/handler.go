package tag

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"server/libs/errors"
	"server/libs/filter"
	"server/libs/logger"
	"server/src/containers/service"
	"server/src/domain"
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
}

func (h *TagHandler) ReadAll(c echo.Context) error {
	items, err := h.tagService.ReadAll()
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusOK, domain.Response{Data: items})
}

func (h *TagHandler) Search(c echo.Context) error {
	var item filter.Request
	if err := c.Bind(&item); err != nil {
		return err
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
	return c.JSON(http.StatusOK, domain.Response{Data: items, Meta: &meta})
}
