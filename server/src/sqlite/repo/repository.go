package repo

import (
	"context"

	"server/libs/errors"
	"server/libs/filter"
	"server/src/domain"
	"server/src/sqlite"
)

type RepoRepository struct {
	db  *sqlite.Database
	ctx context.Context
}

func NewRepoRepository(db *sqlite.Database) domain.RepoRepository {
	return RepoRepository{db, context.Background()}
}

func (repo RepoRepository) ReadOne(item domain.Repo) (domain.Repo, error) {
	var found domain.Repo
	query := "SELECT * FROM repos WHERE id = $1"
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

func (repo RepoRepository) Create(item domain.Repo) (domain.Repo, error) {
	query := `INSERT INTO repos (
        id,
        name,
        description,
        url,
        homepage,
        created_at
    )
    VALUES (?, ?, ?, ?, ?, now)
    RETURNING *`
	return repo.upsert(query, item)
}

func (repo RepoRepository) Update(item domain.Repo) (domain.Repo, error) {
	query := `UPDATE githubs
    SET name = $2,
        interval = $3,
        name = COALESCE(NULLIF($2, ''), name),
        description = COALESCE(NULLIF($2, ''), description),
        url = COALESCE(NULLIF($2, ''), url),
        homepage = COALESCE(NULLIF($2, ''), homepage),
        updated_at = now
    WHERE id = $1
    RETURNING *`
	return repo.upsert(query, item)
}

func (repo RepoRepository) upsert(query string, item domain.Repo) (domain.Repo, error) {
	var saved domain.Repo
	err := repo.db.QueryRow(query, &saved,
		item.ID,
		item.Name,
		item.Description,
		item.URL,
		item.Homepage,
	)
	if err != nil {
		code := errors.Detect(err)
		return saved, errors.New(code, err.Error())
	}
	return saved, nil
}
