package repo

import (
	"testing"

	"server/libs/testutils/assert"
	"server/src/domain"
	"server/src/sqlite"
	"server/src/sqlite/test"
)

func TestCreateRepo(ts *testing.T) {
	test.Seed(ts, "Should save all data", func(t *testing.T, db *sqlite.Database) {
		item := domain.Repo{
			ID:          "MDEwOlJlcG9zaXRvcnk5ODQ2NDQ=",
			Name:        "narrative",
			Description: "A framework for building behaviour-driven tests in fluent Java.",
			URL:         "https://github.com/tim-group/narrative",
			Homepage:    "http://youdevise.github.com/narrative",
		}
		err := NewRepoRepository(db).Create(item)
		assert.Nil(t, err, "")
		var found domain.Repo
		test.NewSQLite(db).QueryRow("repos", &found, item.ID)
		assert.Equal(t, found.ID, item.ID, "")
		assert.Equal(t, found.Name, item.Name, "")
		assert.Equal(t, found.Description, item.Description, "")
		assert.Equal(t, found.URL, item.URL, "")
		assert.Equal(t, found.Homepage, item.Homepage, "")
	})
}
