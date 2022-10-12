package sms

import (
	"developer.zopsmart.com/go/gofr/examples/using-http-service/services"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"sms_func/models"
)

type service struct {
	httpService services.HTTPService
}

// New is a factory function to inject store in handler.
// nolint:revive // handler has to be unexported
func New(httpService services.HTTPService) service {
	return service{httpService: httpService}
}

func (s service) POST(ctx *gofr.Context, sms models.SMS) (models.SMS, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) PUT(ctx *gofr.Context, sms models.SMS) error {
	//TODO implement me
	panic("implement me")
}

func (s service) GetByID(ctx *gofr.Context, id string) (models.SMS, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Get(ctx *gofr.Context) (models.SMS, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Delete(ctx *gofr.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
