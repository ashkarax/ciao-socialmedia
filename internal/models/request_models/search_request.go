package requestmodels

type SearchRequest struct {
	SearchText string `json:"search_text"`
	MyId       string
}
