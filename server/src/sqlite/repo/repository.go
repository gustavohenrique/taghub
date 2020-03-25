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
	err := repo.db.Select().From("repos").Applying(filters).WithTotal(&total).ForEach(func(row interface{}) {
		item := row.(domain.Repo)
		items = append(items, item)
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
	return repo.upsert(query, item)
}

func (repo RepoRepository) Update(item domain.Repo) error {
	query := `UPDATE githubs
    SET name = ?,
        interval = ?,
        name = COALESCE(NULLIF(?, ''), name),
        description = COALESCE(NULLIF(?, ''), description),
        url = COALESCE(NULLIF(?, ''), url),
        homepage = COALESCE(NULLIF(?, ''), homepage),
        updated_at = now
    WHERE id = ?`
	return repo.upsert(query, item)
}

func (repo RepoRepository) upsert(query string, item domain.Repo) error {
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
