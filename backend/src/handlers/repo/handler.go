package repo

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/libs/errors"
	"backend/libs/filter"
	"backend/libs/logger"
	"backend/src/containers/service"
	"backend/src/domain"
	"backend/src/handlers"
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
	route.GET("/:id", h.ReadOne)
	route.DELETE("/:id", h.Remove)
	route.POST("/search", h.Search)
	route.GET("/sync", h.GetTotalStarredRepositories)
	route.POST("/sync", h.Sync)
	route.POST("/tags/search", h.SearchByTagsIDs)
	route.DELETE("/:id/tag/:tag_id", h.RemoveTagFromRepo)
	route.POST("/:repo_id/tag", h.AddTagToRepo)
}

func (h *RepoHandler) Create(c echo.Context) error {
	var item domain.Repo
	if err := handlers.BindAndValidate(c, &item); err != nil {
		return err
	}
	err := h.repoService.Create(item)
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusCreated, domain.Response{Data: item})
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

func (h *RepoHandler) Remove(c echo.Context) error {
	repo := domain.Repo{
		ID: c.Param("id"),
	}
	err := h.repoService.Remove(repo)
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusNoContent, "")
}

func (h *RepoHandler) GetTotalStarredRepositories(c echo.Context) error {
	total, err := h.repoService.GetTotalStarredRepositories()
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusOK, domain.Response{Data: total})
}

func (h *RepoHandler) Sync(c echo.Context) error {
	items, total, err := h.repoService.Sync()
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	logger.Info("Returning", len(items), "/", total, "repositories")
	meta := domain.Meta{
		Page:    1,
		PerPage: 100,
		Total:   int(total),
	}
	return c.JSON(http.StatusOK, domain.Response{Data: items, Meta: &meta})
}

func (h *RepoHandler) SearchByTagsIDs(c echo.Context) error {
	type TagFilter struct {
		filter.Request
		Tags []string `json:"tags"`
	}
	var req TagFilter
	if err := c.Bind(&req); err != nil {
		return err
	}
	items, total, err := h.repoService.SearchByTagsIDs(req.Tags, req.Request)
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	meta := domain.Meta{
		Page:    req.Pagination.Page,
		PerPage: req.Pagination.PerPage,
		Total:   total,
	}
	return c.JSON(http.StatusOK, domain.Response{Data: items, Meta: &meta})
}

func (h *RepoHandler) Search(c echo.Context) error {
	var req filter.Request
	if err := c.Bind(&req); err != nil {
		return err
	}
	items, total, err := h.repoService.Search(req)
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	meta := domain.Meta{
		Page:    req.Pagination.Page,
		PerPage: req.Pagination.PerPage,
		Total:   total,
	}
	return c.JSON(http.StatusOK, domain.Response{Data: items, Meta: &meta})
}

func (h *RepoHandler) RemoveTagFromRepo(c echo.Context) error {
	repo := domain.Repo{
		ID: c.Param("id"),
	}
	tag := domain.Tag{
		ID: c.Param("tag_id"),
	}
	err := h.repoService.RemoveTagFromRepo(repo, tag)
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusNoContent, "")
}

func (h *RepoHandler) AddTagToRepo(c echo.Context) error {
	repo := domain.Repo{ID: c.Param("repo_id")}
	var item domain.Tag
	if err := c.Bind(&item); err != nil {
		return err
	}
	tag, err := h.repoService.AddTagToRepo(repo, item)
	if err != nil {
		logger.Error(err)
		return c.JSON(errors.GetCodeFrom(err), domain.Response{Err: err.Error()})
	}
	return c.JSON(http.StatusOK, domain.Response{Data: tag})
}
