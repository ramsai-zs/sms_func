package sms

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/google/uuid"
	"sms_func/models"
	"sms_func/services"
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

func (s service) GetByID(ctx *gofr.Context, id uuid.UUID) (models.SMS, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Get(ctx *gofr.Context) (models.SMS, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Delete(ctx *gofr.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
