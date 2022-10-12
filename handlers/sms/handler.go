package sms

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"sms_func/services"
)

type handler struct {
	services services.SMS
}

func (h handler) Create(c *gofr.Context) (interface{}, error) {

}

func (h handler) GetByID(c *gofr.Context) (interface{}, error) {

}

func (h handler) Udpate(c *gofr.Context) (interface{}, error) {

}

func (h handler) Delete(c *gofr.Context) (interface{}, error) {

}

// New is a factory function to inject store in handler.
// nolint:revive // handler has to be unexported
func New(service services.SMS) handler {
	return handler{services: service}
}
