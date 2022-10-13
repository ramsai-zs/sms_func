package models

type SMS struct {
	Provider provider `json:"provider"`
	Message  message  `json:"message"`
}
