package cms

type YTUser struct {
	Page  YTPageInfo  `json:"pageInfo"`
	Items []YTChannel `json:"items"`
}

type YTChannel struct {
	Id      string       `json:"id"`
	Page    YTPageInfo   `json:"pageInfo"`
	Snippet YTSnippet    `json:"snippet"`
	Items   []YTPlaylist `json:"items"`
}

type YTPlaylist struct {
	Id      string     `json:"id"`
	Page    YTPageInfo `json:"pageInfo"`
	Snippet YTSnippet  `json:"snippet"`
	Items   []YTVideo  `json:"items"`
}

type YTVideo struct {
	Id      string     `json:"id"`
	Page    YTPageInfo `json:"pageInfo"`
	Snippet YTSnippet  `json:"snippet"`
}

type YTSnippet struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type YTPageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}
