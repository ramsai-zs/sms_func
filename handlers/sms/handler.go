package sms

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"encoding/json"
	"sms_func/models"
	"sms_func/services"
)

type handler struct {
	services services.SMS
}

// New is a factory function to inject store in handler.
// nolint:revive // handler has to be unexported
func New(service services.SMS) handler {
	return handler{services: service}
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var sms models.SMS

	err := bindRequest(ctx, &sms)
	if err != nil {
		return nil, err
	}

	// checks for mandatory fields to create a sms.
	err = createMandatoryCheck(sms)
	if err != nil {
		return nil, err
	}

	resp, err := h.services.POST(ctx, sms)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	var s models.SMS

	sms = ctx.PathParam("id")
	if user.ChannelUserProfileID == "" {
		return nil, errors.MissingParam{Param: []string{"userReferenceID"}}
	}
}

func (h handler) Udpate(c *gofr.Context) (interface{}, error) {

}

func (h handler) Delete(c *gofr.Context) (interface{}, error) {

}

// BindRequest is used to bind the request payload.
func bindRequest(ctx *gofr.Context, model interface{}) error {
	err := ctx.Bind(&model)
	if err != nil {
		e, ok := err.(*json.UnmarshalTypeError)
		if ok {
			return errors.InvalidParam{Param: []string{e.Field}}
		}

		return errors.InvalidParam{Param: []string{"body"}}
	}

	return nil
}

// createMandatoryCheck is used to check the request payload to create a SMS.
func createMandatoryCheck(sms models.SMS) error {
	params := make([]string, 0)

	if sms.Provider.Name == "" {
		params = append(params, "name")
	}

	if sms.Provider.Type == "" {
		params = append(params, "type")
	}

	if sms.Message.Message == "" {
		params = append(params, "message")
	}

	if sms.Message.Number == "" {
		params = append(params, "number")
	}

	if sms.Message.TransactionalID == "" {
		params = append(params, "transactionalID")
	}
}