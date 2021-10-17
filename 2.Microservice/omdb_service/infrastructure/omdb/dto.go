package omdb

type Result struct {
	Search      []MovieSummary `json:"Search"`
	TotalResult string         `json:"totalResults"`
	Response    string         `json:"Response"`
	Error       string         `json:"Error,omitempty"`
}

type MovieSummary struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type Filter struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type MovieInformation struct {
	Response string `json:"Response"`
	Error    string `json:"Error,omitempty"`
}
