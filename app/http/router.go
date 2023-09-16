package http

import (
	"github.com/JobCreator/app/http/endpoints"
	"github.com/JobCreator/app/http/middleware"
	"github.com/JobCreator/app/http/request"
	"github.com/JobCreator/app/http/response"
	"github.com/gin-gonic/gin"
	httpTransport "github.com/go-kit/kit/transport/http"
)

func InitServer() {
	r := gin.Default()

	opts := []httpTransport.ServerOption{
		httpTransport.ServerBefore(middleware.HttpToContext()),
	}

	r.POST(`/v1/passenger/:passenger_id/job`, gin.WrapH(httpTransport.NewServer(
		endpoints.CreateJob(),
		request.DecodeCreateJobRequestJson,
		response.EncodeCreateJobResponseJson,
		opts...,
	)))

	r.Run("0.0.0.0:8089")

}
