package filter

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
)

/*
   Request example:
   {
       "condition": "($1 or $2) and $3 and $4",
       "terms": [
           {
               "id": "1",
               "field": "first_name",
               "operator": "icontains",
               "value": "gus%"
           },
           {
               "id": "2",
               "field": "last_name",
               "operator": "icontains",
               "value": "henrique"
           },
           {
               "id": "3",
               "field": "email",
               "operator": "eq",
               "value": "gustavo@gustavohenrique.net"
           },
           {
               "id": "4",
               "field": "country",
               "operator": "any",
               "value": "brasil",
           }
       ],
       "pagination": {
           "page": 1,
           "per_page": 10
       },
       "ordering": {
           "field": "name",
           "sort": "ASC"
       },
       "groupBy": "optional"
   }
*/

var operators = map[string]string{
	"eq":        "=",
	"ne":        "<>",
	"gt":        ">",
	"lt":        "<",
	"gte":       ">=",
	"lte":       "<=",
	"contains":  "LIKE",
	"icontains": "ILIKE",
	"any":       "ANY",
}

var logicalOperators = map[string]string{
	"+": "AND",
	",": "OR",
}

type Request struct {
	Condition  string     `json:"condition"`
	Pagination Pagination `json:"pagination"`
	Terms      []Term     `json:"terms"`
	Ordering   Ordering   `json:"ordering"`
	Grouping   string     `json:"grouping"`
}

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}

type Ordering struct {
	Field string `json:"field"`
	Sort  string `json:"sort"`
}

type Term struct {
	ID       string `json:"id"`
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

func Parse(raw string) (Request, error) {
	var filter Request
	err := json.Unmarshal([]byte(raw), &filter)
	return filter, err
}

func (f Request) SQL() string {
	sql := f.Where()
	sql += f.GroupBy()
	sql += f.OrderBy()
	sql += f.Limit()
	return sql
}

func (f Request) Where() string {
	sql := "WHERE 1=1 "
	if f.Condition != "" && len(f.Terms) > 0 {
		if strings.Count(f.Condition, "$") == len(f.Terms) {
			sql += "AND"
			condition := f.Condition
			for _, term := range f.Terms {
				operator := operators[strings.ToLower(term.Operator)]
				if operator == "" {
					operator = operators["eq"]
				}
				value := quoteString(removeDangerousWords(term.Value))
				field := removeDangerousWords(term.Field)
				str := field + " " + operator + " " + value

				if operator == operators["any"] {
					str = value + " ILIKE " + operator + "(" + field + ")"
				}
				key := "$" + term.ID
				condition = strings.Replace(condition, key, str, 1)
			}
			sql += " " + condition
		}
	}
	return sql
}

func (f Request) GroupBy() string {
	grouping := f.Grouping
	if grouping != "" {
		return " GROUP BY " + grouping
	}
	return ""
}

func (f Request) OrderBy() string {
	ordering := f.Ordering
	if ordering.Field != "" && ordering.Sort != "" {
		return " ORDER BY " + ordering.Field + " " + ordering.Sort
	}
	return ""
}

func (f Request) GetPerPage() int {
	if f.Pagination.PerPage == 0 {
		return 10
	}
	return f.Pagination.PerPage
}

func (f Request) Limit() string {
	pagination := f.Pagination
	perPage := f.GetPerPage()
	page := int(math.Max(1, float64(pagination.Page)))
	offset := (page - 1) * int(perPage)
	return fmt.Sprintf(" LIMIT %d OFFSET %d", perPage, offset)
}

func quoteString(str string) string {
	return "'" + strings.Replace(str, "'", "''", -1) + "'"
}

func removeDangerousWords(str string) string {
	s := strings.ReplaceAll(str, "\xbf\x27", "")
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, "\"", "")
	return s
}
