package blog

type Item struct {
	Title string `json:"-"`
	Type  string `json:"type"`
	Body  string `json:"body"`
}
