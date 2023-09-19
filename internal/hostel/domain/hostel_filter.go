package domain

import (
	"time"
)

type HostelFilter struct {
	Name      *string    `json:"name,omitempty"`
	Province  *string    `json:"province,omitempty"`
	District  *string    `json:"district,omitempty"`
	Ward      *string    `json:"ward,omitempty"`
	Street    *string    `json:"street,omitempty"`
	PostType  *string    `json:"type,omitempty"`
	Status    *string    `json:"status,omitempty"`
	CostFrom  *int       `json:"costFrom,omitempty"`
	CostTo    *int       `json:"costTo,omitempty"`
	Capacity  *int       `json:"capacity,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	CreatedBy *string    `json:"createdBy,omitempty"`
	PageSize  int
	PageIdx   int
	Sort      string
}
