package domain

type SearchResult struct {
	Hits struct {
		Total struct {
			Value int `json:"value,omitepmty"`
		} `json:"total,omitepmty"`
		Hits []struct {
			Source Post `json:"_source,omitepmty"`
		} `json:"hits,omitepmty"`
	} `json:"hits,omitepmty"`
}

type Range struct {
	LTE   int `json:"lte,omitepmty"`
	GTE   int `json:"gte,omitepmty"`
	Boost int `json:"boost,omitepmty"`
}
