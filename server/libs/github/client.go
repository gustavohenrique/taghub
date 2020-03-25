package github

import (
	"encoding/json"
	"fmt"
	"log"

	"server/libs/httpclient"
	"server/libs/stringutils"
)

type GithubClient interface {
	FetchAllStarredRepositories(all []Repository, after string) ([]Repository, error)
	GetStarredRepositories(perPage int, after string) ([]Repository, error)
}

type GithubCli struct {
	httpClient *httpclient.HTTPClient
}

func NewGithubClient(token string) GithubClient {
	config := httpclient.Config{
		Timeout: 120,
		BaseURL: "https://api.github.com/graphql",
		Token:   token,
	}
	httpClient := httpclient.New(config)
	return &GithubCli{httpClient}
}

func (s *GithubCli) FetchAllStarredRepositories(all []Repository, after string) ([]Repository, error) {
	repos, err := s.GetStarredRepositories(100, after)
	if err != nil || len(repos) == 0 {
		return all, err
	}
	all = append(all, repos...)
	last := repos[len(repos)-1]
	return s.FetchAllStarredRepositories(all, last.Cursor)
}

func (s *GithubCli) GetStarredRepositories(perPage int, after string) ([]Repository, error) {
	var repos []Repository
	params := fmt.Sprintf("first:%d", perPage)
	if after != "" {
		params = fmt.Sprintf(`%s, after:\"%s\"`, params, after)
	}
	query := stringutils.TrimSpaceNewlineInString(`{"query": "query { 
      viewer { 
        starredRepositories(%s) {
          edges {
            cursor,
            repositories:node {
              id,
              name,
              url,
              description,
              homepageUrl
            }
          },
          totalCount
        }
      }
    }"}`)
	req := fmt.Sprintf(query, params)
	body, statusCode, err := s.httpClient.POST("", []byte(req))
	if hasFail(err, statusCode) {
		log.Println("req=", req, "res=", string(body))
		return repos, fmt.Errorf("%d %s", statusCode, err)
	}
	var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return repos, fmt.Errorf("Cannot unmarshal response %s", err)
	}
	edges := res.Data.Viewer.StarredRepositories.Edges
	for _, edge := range edges {
		repo := edge.Repository
		repo.Cursor = edge.Cursor
		repos = append(repos, repo)
	}
	return repos, err
}

func hasFail(err error, statusCode int) bool {
	return err != nil || statusCode >= 300
}
