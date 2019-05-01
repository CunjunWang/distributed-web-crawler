package model

type SearchResult struct {
	Hits  int
	Start int
	Query string
	Items []interface{}
}
