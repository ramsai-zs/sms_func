package services

import (
	"context"
	"github.com/google/uuid"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	svc "developer.zopsmart.com/go/gofr/pkg/service"
	"sms_func/models"
)

type HTTPService interface {
	Post(ctx context.Context, api string, params map[string]interface{}, body []byte) (*svc.Response, error)
	Patch(ctx context.Context, api string, params map[string]interface{}, body []byte) (*svc.Response, error)
	Get(ctx context.Context, api string, params map[string]interface{}) (*svc.Response, error)
	Delete(ctx context.Context, api string, body []byte) (*svc.Response, error)
	Bind(resp []byte, i interface{}) error
}

type SMS interface {
	POST(ctx *gofr.Context, sms models.SMS) (models.SMS, error)
	PUT(ctx *gofr.Context, sms models.SMS) error
	GetByID(ctx *gofr.Context, id uuid.UUID) (models.SMS, error)
	Get(ctx *gofr.Context) ([]models.SMS, error)
	Delete(ctx *gofr.Context, id uuid.UUID) error
}
