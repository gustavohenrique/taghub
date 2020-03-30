package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"server/libs/configuration"
	"server/libs/filter"
	"server/libs/logger"
	"server/src/containers/repository"
	"server/src/domain"
	"server/src/sqlite"
)

var config *configuration.AppConfig

func init() {
	config = configuration.Load()
}

func main() {
	db := sqlite.New(sqlite.Config{
		URL: config.DatabaseURL,
	})
	db.Connect()
	fmt.Println(config.DatabaseURL)

	repositories := repository.NewRepositoryContainer(db)
	tagRepository := repositories.TagRepository
	items, err := tagRepository.ReadAll()
	if err != nil {
		logger.Error(err)
		return
	}
	var tags []domain.Tag
	for _, tag := range items {
		total, _ := tagRepository.GetTotalReposByTag(tag)
		tag.TotalRepos = total
		tags = append(tags, tag)
	}
	filesave("tags.json", tags)

	repoRepository := repositories.RepoRepository
	totalRepos := repoRepository.GetTotal()
	filters := filter.Request{
		Pagination: filter.Pagination{
			Page:    1,
			PerPage: totalRepos,
		},
		Ordering: filter.Ordering{
			Field: "created_at",
			Sort:  "DESC",
		},
	}
	for _, tag := range tags {
		repos, _, err := repoRepository.SearchByTagsIDs([]string{tag.ID}, filters)
		if err != nil {
			logger.Error(err)
			return
		}
		resp := domain.Response{
			Data: repos,
			Meta: &domain.Meta{
				Page:    1,
				PerPage: totalRepos,
				Total:   len(repos),
			},
		}
		filesave("tags/"+tag.ID+".json", resp)
	}

	perPage := 10
	filters.Pagination.PerPage = perPage
	maxPage := totalRepos / perPage
	fmt.Println(perPage, "of", maxPage, "/", totalRepos)
	for page := 1; page < maxPage; page++ {
		filters.Pagination.Page = page
		repos, _, err := repoRepository.Search(filters)
		if err != nil {
			logger.Error(err)
			return
		}
		filename := fmt.Sprintf("repos/%d.json", page)
		resp := domain.Response{
			Data: repos,
			Meta: &domain.Meta{
				Page:    page,
				PerPage: perPage,
				Total:   totalRepos,
			},
		}
		filesave(filename, resp)
	}

	fmt.Println("Finished. Enjoy it!")
}

func filesave(filename string, items interface{}) {
	b, err := json.Marshal(items)
	if err != nil {
		logger.Error(err)
		return
	}
	destDir := config.ExportDir
	out := fmt.Sprintf("%s/%s", destDir, filename)
	err = ioutil.WriteFile(out, b, os.ModePerm)
	if err != nil {
		logger.Error(err)
	}
}
