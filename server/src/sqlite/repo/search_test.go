package repo

import (
	"fmt"
	"testing"

	"server/libs/filter"
	"server/libs/testutils/assert"
	"server/src/sqlite"
	"server/src/sqlite/test"
)

func TestSearchRepoReturnAllItems(ts *testing.T) {
	test.Seed(ts, "Should return all items ordering by name", func(t *testing.T, db *sqlite.Database) {
		item := filter.Request{
			Ordering: filter.Ordering{
				Field: "created_at",
				Sort:  "asc",
			},
		}
		items, total, err := NewRepoRepository(db).Search(item)
		assert.Nil(t, err, "")
		assert.Equal(t, len(items), 3, fmt.Sprintf("Got %d", len(items)))
		assert.Equal(t, total, 3, fmt.Sprintf("Got total %d", total))

		first := items[0]
		assert.Equal(t, first.ID, "MDEwOlJlcG9zaXRvcnkzMDY1NDU0", "")
		assert.Equal(t, len(first.Tags), 2, "")
	})

	test.Seed(ts, "Should return item found by name", func(t *testing.T, db *sqlite.Database) {
		item := filter.Request{
			Condition: "$1",
			Terms: []filter.Term{
				filter.Term{
					ID:       "1",
					Field:    "name",
					Operator: "contains",
					Value:    "%impre%",
				},
			},
		}
		items, total, err := NewRepoRepository(db).Search(item)
		assert.Nil(t, err, "")
		assert.Equal(t, len(items), 1, fmt.Sprintf("Got %d", len(items)))
		assert.Equal(t, total, 1, fmt.Sprintf("Got total %d", total))

		first := items[0]
		assert.Equal(t, first.ID, "MDEwOlJlcG9zaXRvcnkzMDY1NDU0", "")
		assert.Equal(t, len(first.Tags), 2, "")
	})
}
