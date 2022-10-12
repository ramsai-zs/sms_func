package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/service"
	smshandler "sms_func/handlers/sms"
	"sms_func/services/sms"
)

func main() {
	app := gofr.New()

	smsDataService := service.NewHTTPServiceWithOptions(app.Config.Get(""), app.Logger, nil)
	smsService := sms.New(smsDataService)
	handler := smshandler.New(smsService)

	app.POST("/sms", handler.Create)
	app.GET("/sms/{id}", handler.GetByID)
	app.PUT("/sms/{id}", handler.Update)
	app.DELETE("/sms", handler.Delete)

	app.Start()
}
