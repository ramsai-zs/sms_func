package models

import (
	"github.com/google/uuid"
	"time"
)

type message struct {
	MessageRefID    uuid.UUID `json:"messageRefID"`
	Message         string    `json:"message"`
	Number          string    `json:"number"`
	TransactionalID string    `json:"transactionalID"`
	Status          bool      `json:"status"`
	CreationTime    time.Time `json:"creationTime"`
}
