package sms

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"sms_func/models"
	"sms_func/services"
)

func (s service) create(ctx *gofr.Context, sms models.SMS) (models.SMS, error) {
	// Marshal converts Go objects into JSON Format.
	body, err := json.Marshal(sms)
	if err != nil {
		return models.SMS{}, errors.Error("Marshal Error")
	}

	resp, err := s.httpService.Post(ctx, "", nil, body)
	if err != nil {
		return models.SMS{}, err
	}

	// checking status code for the response.
	if resp.StatusCode != http.StatusCreated {
		return models.SMS{}, services.GetError(resp)
	}

	var msg models.SMS

	// getResponse binds the data from data service response  to SMS.
	err = s.getResponse(resp.Body, &msg)
	if err != nil {
		return models.SMS{}, errors.EntityNotFound{}
	}

	return msg, nil
}

func (s service) getByID(ctx *gofr.Context, id uuid.UUID) (models.SMS, error) {
	resp, err := s.httpService.Get(ctx, "", nil)
	if err != nil {
		return models.SMS{}, err
	}

	// checking status code for the response.
	if resp.StatusCode != http.StatusOK {
		return models.SMS{}, services.GetError(resp)
	}

	var msg models.SMS

	// getResponse binds the data from data service response  to SMS.
	err = s.getResponse(resp.Body, &msg)
	if err != nil {
		return models.SMS{}, errors.EntityNotFound{}
	}

	return msg, nil
}

func (s service) get(ctx *gofr.Context) (models.SMS, error) {
	resp, err := s.httpService.Get(ctx, "", nil)
	if err != nil {
		return models.SMS{}, err
	}

	// checking status code for the response.
	if resp.StatusCode != http.StatusOK {
		return models.SMS{}, services.GetError(resp)
	}

	var msg models.SMS

	// getResponse binds the data from data service response  to SMS.
	err = s.getResponse(resp.Body, &msg)
	if err != nil {
		return models.SMS{}, errors.EntityNotFound{}
	}

	return msg, nil
}

func (s service) update(ctx *gofr.Context, sms models.SMS) error {
	// Marshal converts Go objects into JSON Format.
	body, err := json.Marshal(sms)
	if err != nil {
		return errors.Error("Marshal Error")
	}

	resp, err := s.httpService.Patch(ctx, "", nil, body)
	if err != nil {
		return err
	}

	// checking status code for the response.
	if resp.StatusCode != http.StatusOK {
		return services.GetError(resp)
	}

	var msg models.SMS

	// getResponse binds the data from data service response  to SMS.
	err = s.getResponse(resp.Body, &msg)
	if err != nil {
		return errors.EntityNotFound{}
	}

	return nil
}

func (s service) delete(ctx *gofr.Context, id uuid.UUID) error {
	resp, err := s.httpService.Delete(ctx, "", nil)
	if err != nil {
		return err
	}

	// checking status code for the response.
	if resp.StatusCode != http.StatusOK {
		return services.GetError(resp)
	}

	var msg models.SMS

	// getResponse binds the data from data service response  to SMS.
	err = s.getResponse(resp.Body, &msg)
	if err != nil {
		return errors.EntityNotFound{}
	}

	return nil
}

// getResponse unmarshalls the data response.
// If error occurs while unmarshalling, it returns Bind Error.
func (s service) getResponse(body []byte, data interface{}) error {
	resp := struct {
		Data interface{} `json:"data"`
	}{Data: data}

	return s.bind(body, &resp)
}

// bind unmarshalls response body to data and returns Bind Error if an error occurs.
func (s service) bind(body []byte, data interface{}) error {
	err := s.httpService.Bind(body, data)
	if err != nil {
		return errors.InvalidParam{Param: []string{fmt.Sprintf("unable to bind : %v", err)}}
	}

	return nil
}
