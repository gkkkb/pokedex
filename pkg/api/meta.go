package api

// IndexMeta contains metadata of an index response
type IndexMeta struct {
	HTTPStatus int    `json:"http_status"`
	Limit      uint64 `json:"limit"`
	Offset     uint64 `json:"offset"`
	Total      int    `json:"total"`
	TotalPages uint   `json:"total_pages,omitempty"`
}
