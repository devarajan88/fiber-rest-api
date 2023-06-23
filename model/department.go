package model

import "time"

type Department struct {
	Id               uint      `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	HeadOfDepartment string    `json:"headOfDepartment"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
