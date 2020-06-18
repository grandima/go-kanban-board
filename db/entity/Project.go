package entity

import "database/sql"

type Project struct {
	Id	int `json:"id"`
	Name string `json:"name"`
	Description sql.NullString `json:"description"`
}