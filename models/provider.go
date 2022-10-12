package models

import "time"

type provider struct {
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	CreationTime time.Time `json:"creationTime"`
}
