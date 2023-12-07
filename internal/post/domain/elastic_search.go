package domain

type SearchResult struct {
	Hits struct {
		Total struct {
			Relation string `json:"relation,omitepmty"`
			Value    int64  `json:"value,omitepmty"`
		} `json:"total,omitepmty"`
		Hits []struct {
			Source Post    `json:"_source,omitepmty"`
			Score  float64 `json:"_score,omitepmty"`
			Id     string  `json:"_id,omitepmty"`
			Index  string  `json:"_index,omitepmty"`
		} `json:"hits,omitepmty"`
	} `json:"hits,omitepmty"`
}

type Range struct {
	LTE   int `json:"lte,omitepmty"`
	GTE   int `json:"gte,omitepmty"`
	Boost int `json:"boost,omitepmty"`
}
