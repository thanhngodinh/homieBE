package domain

import (
	"time"

	"github.com/lib/pq"
)

type HostelFilter struct {
	Name         *string        `json:"name,omitempty"`
	Province     *string        `json:"province,omitempty"`
	District     *string        `json:"district,omitempty"`
	Ward         *string        `json:"ward,omitempty"`
	Street       *string        `json:"street,omitempty"`
	Status       *string        `json:"status,omitempty"`
	CostFrom     *int           `json:"costFrom,omitempty"`
	CostTo       *int           `json:"costTo,omitempty"`
	DepositFrom  *int           `json:"depositFrom,omitempty"`
	DepositTo    *int           `json:"depositTo,omitempty"`
	Capacity     *int           `json:"capacity,omitempty"`
	Utilities    pq.StringArray `json:"utilities,omitempty"`
	CapacityFrom *int           `json:"-"`
	CapacityTo   *int           `json:"-"`
	CreatedAt    *time.Time     `json:"createdAt,omitempty"`
	CreatedBy    *string        `json:"createdBy,omitempty"`
	PageSize     int            `json:"pageSize,omitempty"`
	PageIdx      int            `json:"pageIdx,omitempty"`
	Sort         string         `json:"sort,omitempty"`
}
