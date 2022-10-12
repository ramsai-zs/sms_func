package services

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/service"
	"encoding/json"
)

func GetError(res *service.Response) error {
	resp := struct {
		Errors []errors.Response `json:"errors"`
	}{}

	err := json.Unmarshal(res.Body, &resp)
	if err != nil {
		return err
	}

	if len(resp.Errors) > 0 {
		err := &resp.Errors[0]
		err.StatusCode = res.StatusCode

		return err
	}

	return nil
}
