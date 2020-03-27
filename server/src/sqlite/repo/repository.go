package repo

import (
	"server/libs/errors"
	"server/libs/filter"
	"server/src/domain"
	"server/src/sqlite"
)

type RepoRepository struct {
	db *sqlite.Database
}

func NewRepoRepository(db *sqlite.Database) domain.RepoRepository {
	return RepoRepository{db}
}

func (repo RepoRepository) ReadOne(item domain.Repo) (domain.Repo, error) {
	var found domain.Repo
	query := "SELECT * FROM repos WHERE id = ?"
	err := repo.db.QueryRow(query, &found, item.ID)
	if err != nil {
		code := errors.Detect(err)
		return found, errors.New(code, err.Error())
	}
	return found, nil
}

func (repo RepoRepository) Search(filters filter.Request) ([]domain.Repo, int, error) {
	var items []domain.Repo
	var total int
	err := repo.db.
		Select().
		From("repos").
		Applying(filters).
		WithTotal(&total).
		ForEach(&domain.Repo{}, func(row interface{}) {
			item := row.(*domain.Repo)
			var tags []domain.Tag
			repo.db.QueryAll("SELECT id, name FROM tags JOIN mapping ON mapping.tag_id = tags.id WHERE mapping.repo_id = ?", &tags, item.ID)
			item.Tags = tags
			items = append(items, *item)
		})
	if err != nil {
		code := errors.Detect(err)
		return items, total, errors.New(code, err.Error())
	}
	return items, total, nil
}

func (repo RepoRepository) SearchTag(filters filter.Request) ([]domain.Tag, int, error) {
	var items []domain.Tag
	var total int
	err := repo.db.Select().From("tags").Applying(filters).WithTotal(&total).ForEach(&domain.Tag{}, func(row interface{}) {
		item := row.(*domain.Tag)
		items = append(items, *item)
	})
	if err != nil {
		code := errors.Detect(err)
		return items, total, errors.New(code, err.Error())
	}
	return items, total, nil
}

func (repo RepoRepository) Create(item domain.Repo) error {
	query := `INSERT INTO repos (
        id,
        name,
        description,
        url,
        homepage
    )
    VALUES (?, ?, ?, ?, ?)`
	err := repo.db.Exec(query,
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

func (repo RepoRepository) GetTagByName(name string) (domain.Tag, error) {
	var found domain.Tag
	query := "SELECT id, name FROM tags WHERE name = ?"
	err := repo.db.QueryRow(query, &found, name)
	if err != nil {
		code := errors.Detect(err)
		return found, errors.New(code, err.Error())
	}
	return found, nil
}

func (repo RepoRepository) AddTagToRepo(item domain.Repo, tag domain.Tag) error {
	query := "INSERT INTO mapping (repo_id, tag_id) VALUES (?, ?)"
	err := repo.db.Exec(query,
		item.ID,
		tag.ID,
	)
	if err != nil {
		code := errors.Detect(err)
		return errors.New(code, err.Error())
	}
	return nil
}

func (repo RepoRepository) RemoveTagFromRepo(item domain.Repo, tag domain.Tag) error {
	query := "DELETE FROM mapping WHERE repo_id = (?) AND tag_id = (?)"
	err := repo.db.Exec(query,
		item.ID,
		tag.ID,
	)
	if err != nil {
		code := errors.Detect(err)
		return errors.New(code, err.Error())
	}
	return nil
}
