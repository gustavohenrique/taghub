package repo

import (
	"fmt"
	"strings"

	"server/libs/errors"
	"server/libs/filter"
	"server/libs/logger"
	"server/src/domain"
	"server/src/sqlite"
)

type RepoRepository struct {
	db *sqlite.Database
}

func NewRepoRepository(db *sqlite.Database) domain.RepoRepository {
	return RepoRepository{db}
}

func (r RepoRepository) ReadOne(item domain.Repo) (domain.Repo, error) {
	var found domain.Repo
	query := "SELECT * FROM repos WHERE id = ?"
	err := r.db.QueryRow(query, &found, item.ID)
	if err != nil {
		code := errors.Detect(err)
		return found, errors.New(code, err.Error())
	}
	return found, nil
}

func (r RepoRepository) SearchByTagsIDs(tags []string, filters filter.Request) ([]domain.Repo, int, error) {
	var repos []domain.Repo
	var total int

	query := `SELECT COUNT(DISTINCT id) FROM repos JOIN mapping WHERE mapping.repo_id = repos.id AND mapping.tag_id IN ('%s')`
	query = fmt.Sprintf(query, strings.Join(tags, "', '"))
	err := r.db.Get(query, &total)
	if err != nil {
		code := errors.Detect(err)
		return repos, total, errors.New(code, err.Error())
	}

	query = `SELECT repos.* FROM repos JOIN mapping WHERE mapping.repo_id = repos.id AND mapping.tag_id IN ('%s') GROUP BY repos.name %s %s`
	query = fmt.Sprintf(query, strings.Join(tags, "', '"), filters.OrderBy(), filters.Limit())
	err = r.db.QueryAll(query, &repos)
	if err != nil {
		code := errors.Detect(err)
		return repos, total, errors.New(code, err.Error())
	}

	var items []domain.Repo
	for _, item := range repos {
		var tags []domain.Tag
		r.db.QueryAll("SELECT id, name FROM tags JOIN mapping ON mapping.tag_id = tags.id WHERE mapping.repo_id = ?", &tags, item.ID)
		item.Tags = tags
		items = append(items, item)
	}
	return items, total, nil
}

func (r RepoRepository) Search(filters filter.Request) ([]domain.Repo, int, error) {
	var items []domain.Repo
	var total int
	err := r.db.
		Select().
		From("repos").
		Applying(filters).
		WithTotal(&total).
		ForEach(&domain.Repo{}, func(row interface{}) {
			item := row.(*domain.Repo)
			var tags []domain.Tag
			err := r.db.QueryAll("SELECT id, name FROM tags JOIN mapping ON mapping.tag_id = tags.id WHERE mapping.repo_id = ?", &tags, item.ID)
			if err != nil {
				logger.Error(err)
			}
			item.Tags = tags
			items = append(items, *item)
		})
	if err != nil {
		code := errors.Detect(err)
		return items, total, errors.New(code, err.Error())
	}
	return items, total, nil
}

func (r RepoRepository) Create(item domain.Repo) error {
	query := `INSERT INTO repos (
        id,
        name,
        description,
        url,
        homepage
    )
    VALUES (?, ?, ?, ?, ?)`
	err := r.db.Exec(query,
		item.ID,
		item.Name,
		item.Description,
		item.URL,
		item.Homepage,
	)
	if err != nil {
		code := errors.Detect(err)
		return errors.New(code, err.Error())
	}
	return nil
}

func (r RepoRepository) GetTagByName(name string) (domain.Tag, error) {
	var found domain.Tag
	query := "SELECT id, name FROM tags WHERE name = ?"
	err := r.db.QueryRow(query, &found, name)
	if err != nil {
		code := errors.Detect(err)
		return found, errors.New(code, err.Error())
	}
	return found, nil
}

func (r RepoRepository) AddTagToRepo(item domain.Repo, tag domain.Tag) error {
	query := "INSERT INTO mapping (repo_id, tag_id) VALUES (?, ?)"
	err := r.db.Exec(query,
		item.ID,
		tag.ID,
	)
	if err != nil {
		code := errors.Detect(err)
		return errors.New(code, err.Error())
	}
	return nil
}

func (r RepoRepository) RemoveTagFromRepo(item domain.Repo, tag domain.Tag) error {
	query := "DELETE FROM mapping WHERE repo_id = (?) AND tag_id = (?)"
	err := r.db.Exec(query,
		item.ID,
		tag.ID,
	)
	if err != nil {
		code := errors.Detect(err)
		return errors.New(code, err.Error())
	}
	return nil
}
