// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package service2

import (
	"tg-example/contracts"

	"github.com/gofiber/fiber/v2"
)

type httpService2 struct {
	errorHandler     ErrorHandler
	maxBatchSize     int
	maxParallelBatch int
	svc              *serverService2
	base             contracts.Service2
}

func NewService2(svcService2 contracts.Service2) (srv *httpService2) {

	srv = &httpService2{
		base: svcService2,
		svc:  newServerService2(svcService2),
	}
	return
}

func (http *httpService2) Service() *serverService2 {
	return http.svc
}

func (http *httpService2) WithLog() *httpService2 {
	http.svc.WithLog()
	return http
}

func (http *httpService2) WithTrace() *httpService2 {
	http.svc.WithTrace()
	return http
}

func (http *httpService2) WithMetrics() *httpService2 {
	http.svc.WithMetrics()
	return http
}

func (http *httpService2) WithErrorHandler(handler ErrorHandler) *httpService2 {
	http.errorHandler = handler
	return http
}

func (http *httpService2) SetRoutes(route *fiber.App) {
	route.Post("/v1/service2", http.serveBatch)
	route.Post("/v1/service2/set", http.serveSet)
}
