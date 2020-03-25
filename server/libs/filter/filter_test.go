package filter_test

import (
	"testing"

	"server/libs/filter"
)

var req filter.Request

func beforeEach() {
	req = filter.Request{}
	req.Condition = "$1 or $2"
	term1 := filter.Term{
		ID:       "1",
		Field:    "name",
		Operator: "eq",
		Value:    "gustavo",
	}
	term2 := filter.Term{
		ID:       "2",
		Field:    "age",
		Operator: "gte",
		Value:    "20",
	}
	req.Terms = []filter.Term{term1, term2}
	req.Pagination = filter.Pagination{
		PerPage: 10,
		Page:    1,
	}
	req.Ordering = filter.Ordering{
		Field: "id",
		Sort:  "DESC",
	}
}

func TestContainsShouldUseUnaccent(t *testing.T) {
	req = filter.Request{}
	req.Terms = []filter.Term{
		filter.Term{
			ID:       "1",
			Value:    "%questões%",
			Field:    "name",
			Operator: "contains",
		},
		filter.Term{
			ID:       "2",
			Value:    "36",
			Field:    "age",
			Operator: "eq",
		},
	}
	req.Condition = "$1 AND $2"
	sql := req.SQL()
	expected := "WHERE 1=1 AND (name LIKE '%questões%' OR unaccent(name) LIKE '%questões%') AND age = '36' LIMIT 10 OFFSET 0"
	if expected != sql {
		t.Errorf("EXPECTED: %s | GOT: %s", expected, sql)
	}
}

func TestReturnParametersAccordingOfRequest(t *testing.T) {
	beforeEach()
	sql := req.SQL()
	expected := "WHERE 1=1 AND name = 'gustavo' or age >= '20' ORDER BY id DESC LIMIT 10 OFFSET 0"
	if expected != sql {
		t.Errorf("EXPECTED: %s | GOT: %s", expected, sql)
	}
}

func TestDefaultOrdering(t *testing.T) {
	beforeEach()
	req.Ordering = filter.Ordering{}
	sql := req.SQL()
	if "WHERE 1=1 AND name = 'gustavo' or age >= '20' LIMIT 10 OFFSET 0" != sql {
		t.Errorf("Failed")
	}
}

func TestDefaultPagination(t *testing.T) {
	beforeEach()
	req.Pagination = filter.Pagination{}
	sql := req.SQL()
	if "WHERE 1=1 AND name = 'gustavo' or age >= '20' ORDER BY id DESC LIMIT 10 OFFSET 0" != sql {
		t.Errorf("Failed")
	}
}

func TestUsingAnyInArrays(t *testing.T) {
	req = filter.Request{}
	req.Condition = "$1 and $2"
	term1 := filter.Term{
		ID:       "1",
		Field:    "name",
		Operator: "eq",
		Value:    "gustavo",
	}
	term2 := filter.Term{
		ID:       "2",
		Field:    "country",
		Operator: "any",
		Value:    "brasil",
	}
	req.Terms = []filter.Term{term1, term2}
	sql := req.SQL()
	expected := "WHERE 1=1 AND name = 'gustavo' and 'brasil' ILIKE ANY(country) LIMIT 10 OFFSET 0"
	if expected != sql {
		t.Errorf("Failed. Expected %s but got %s", expected, sql)
	}
}

func TestIgnoreTermWhenItIsEmpty(t *testing.T) {
	beforeEach()
	req.Terms = []filter.Term{}
	sql := req.SQL()
	if "WHERE 1=1  ORDER BY id DESC LIMIT 10 OFFSET 0" != sql {
		t.Errorf("Failed")
	}
}

func TestIgnoreRequestWhenTotalTermsMismatchTotalParams(t *testing.T) {
	beforeEach()
	req.Condition = "$1 or $2 and $3"
	sql := req.SQL()
	if "WHERE 1=1  ORDER BY id DESC LIMIT 10 OFFSET 0" != sql {
		t.Errorf("Failed")
	}
}

func TestSQLInjection(t *testing.T) {
	beforeEach()
	req.Terms[0] = filter.Term{
		ID:       "1",
		Field:    "name",
		Operator: "eq",
		Value:    "\xbf\x27 or (select amount from orders limit 1) as created_at or name='",
	}
	sql := req.SQL()
	expected := "WHERE 1=1 AND name = ' or (select amount from orders limit 1) as created_at or name=' or age >= '20' ORDER BY id DESC LIMIT 10 OFFSET 0"
	if expected != sql {
		t.Errorf("EXPECTED: %s | GOT: %s", expected, sql)
	}
}
