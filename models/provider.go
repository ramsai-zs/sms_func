package models

import (
	"time"
)

type provider struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	CreationTime time.Time `json:"creationTime"`
}
