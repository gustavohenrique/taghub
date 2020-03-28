package tag

import (
	"server/libs/errors"
	"server/libs/filter"
	"server/src/domain"
	"server/src/sqlite"
)

type TagRepository struct {
	db *sqlite.Database
}

func NewTagRepository(db *sqlite.Database) domain.TagRepository {
	return TagRepository{db}
}

func (r TagRepository) Create(item domain.Tag) (domain.Tag, error) {
	query := "INSERT INTO tags (name) VALUES (?)"
	id, err := r.db.ExecAndGetLastID(query,
		item.Name,
	)
	if err != nil {
		code := errors.Detect(err)
		return item, errors.New(code, err.Error())
	}
	item.ID = id
	return item, nil
}

func (r TagRepository) Update(item domain.Tag) error {
	query := "UPDATE tags SET name = ? WHERE id = ?"
	err := r.db.Exec(query,
		item.Name,
		item.ID,
	)
	if err != nil {
		code := errors.Detect(err)
		return errors.New(code, err.Error())
	}
	return nil
}

func (r TagRepository) ReadOne(item domain.Tag) (domain.Tag, error) {
	query := "SELECT id, name FROM tags WHERE id = ? OR name = ? LIMIT 1"
	var found domain.Tag
	err := r.db.QueryRow(query, &found, item.ID, item.Name)
	if err != nil {
		code := errors.Detect(err)
		return found, errors.New(code, err.Error())
	}
	return found, nil
}

func (r TagRepository) ReadAll() ([]domain.Tag, error) {
	query := "SELECT id, name FROM tags ORDER BY name"
	var items []domain.Tag
	err := r.db.QueryAll(query, &items)
	if err != nil {
		code := errors.Detect(err)
		return items, errors.New(code, err.Error())
	}
	return items, nil
}

func (r TagRepository) Remove(item domain.Tag) error {
	query := "DELETE FROM tags WHERE id = ?"
	err := r.db.Exec(query, item.ID)
	if err != nil {
		code := errors.Detect(err)
		return errors.New(code, err.Error())
	}
	return nil
}

func (r TagRepository) GetTotalReposByTag(item domain.Tag) (int, error) {
	query := "SELECT COUNT(*) FROM mapping WHERE tag_id = ?"
	var total int
	err := r.db.Get(query, &total, item.ID)
	if err != nil {
		code := errors.Detect(err)
		return total, errors.New(code, err.Error())
	}
	return total, nil
}

func (r TagRepository) Search(filters filter.Request) ([]domain.Tag, int, error) {
	var items []domain.Tag
	var total int
	err := r.db.Select().From("tags").Applying(filters).WithTotal(&total).ForEach(&domain.Tag{}, func(row interface{}) {
		item := row.(*domain.Tag)
		items = append(items, *item)
	})
	if err != nil {
		code := errors.Detect(err)
		return items, total, errors.New(code, err.Error())
	}
	return items, total, nil
}
