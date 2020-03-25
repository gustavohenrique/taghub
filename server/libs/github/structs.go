package github

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	Viewer Viewer `json:"viewer"`
}

type Viewer struct {
	StarredRepositories StarredRepositories `json:"starredRepositories"`
}

type StarredRepositories struct {
	Edges []Edge `json:"edges"`
}

type Edge struct {
	Cursor     string
	Repository Repository `json:"repositories"`
}

type Repository struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Homepage    string `json:"homepage"`
	Cursor      string `json:"cursor"`
}
