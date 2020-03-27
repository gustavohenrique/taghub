package tag

import (
	"testing"

	"server/libs/testutils/assert"
	"server/src/domain"
	"server/src/sqlite"
	"server/src/sqlite/test"
)

func TestCreateTag(ts *testing.T) {
	test.Seed(ts, "Should save all data", func(t *testing.T, db *sqlite.Database) {
		item := domain.Tag{
			Name: "linux",
		}
		saved, err := NewTagRepository(db).Create(item)
		assert.Nil(t, err, "")
		assert.True(t, saved.ID != "", saved.ID)
	})

	test.Seed(ts, "Should fail when duplicate name", func(t *testing.T, db *sqlite.Database) {
		item := domain.Tag{
			Name: "golang",
		}
		_, err := NewTagRepository(db).Create(item)
		assert.NotNil(t, err, "")
	})
}
