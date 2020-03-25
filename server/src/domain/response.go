package domain

import "fmt"

type Meta struct {
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
	Total   int `json:"total,omitempty"`
}

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Err  string      `json:"error,omitempty"`
	Meta *Meta       `json:"meta,omitempty"`
}

func (r *Response) String() string {
	if r.Data != nil {
		return fmt.Sprintf("%s", r.Data)
	}
	return "Sorry, the response data is empty."
}
