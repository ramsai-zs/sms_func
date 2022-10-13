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
	_, err := s.create(ctx, sms)
	if err != nil {
		return models.SMS{}, err
	}

	_, err = s.createMessage(ctx, sms)
	if err != nil {
		return models.SMS{}, err
	}

	return models.SMS{}, nil
}

func (s service) PUT(ctx *gofr.Context, sms models.SMS) error {
	err := s.update(ctx, sms)
	if err != nil {
		return err
	}

	err = s.updateMessage(ctx, sms)
	if err != nil {
		return err
	}

	return nil
}

func (s service) GetByID(ctx *gofr.Context, id uuid.UUID) (models.SMS, error) {
	_, err := s.getByID(ctx, id)
	if err != nil {
		return models.SMS{}, err
	}

	_, err = s.getByProviderID(ctx, id)
	if err != nil {
		return models.SMS{}, err
	}

	return models.SMS{}, err
}

func (s service) Get(ctx *gofr.Context) ([]models.SMS, error) {
	_, err := s.get(ctx)
	if err != nil {
		return nil, err
	}

	_, err = s.getAll(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s service) Delete(ctx *gofr.Context, id uuid.UUID) error {
	err := s.delete(ctx, id)
	if err != nil {
		return err
	}

	err = s.deleteMessage(ctx, id)
	if err != nil {
		return err
	}

	return err
}
