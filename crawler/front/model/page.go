package page

type SearchResult struct {
	Q     string
	Hits  int64
	Start int
	Items []interface{}
}
