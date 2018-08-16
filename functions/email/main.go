package main

import (
	"encoding/json"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serinth/serverless-emailer/api"
	"github.com/serinth/serverless-emailer/util"
	"github.com/serinth/serverless-emailer/validators"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var cfg *util.Config

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var req api.SendEmailRequest

	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return api.NewResponseBuilder().Status(http.StatusBadRequest).Build(), nil
	}

	requestErrors := validators.GetSendEmailRequestErrors(req)
	if requestErrors != nil {
		invalidRequestResponse, _ := json.Marshal(api.InvalidEmailRequest(requestErrors))
		return api.NewResponseBuilder().
			Body(string(invalidRequestResponse)).
			Status(http.StatusBadRequest).
			Build(), nil
	}

	if err := sendEmail(&req, cfg); err != nil {
		internalErrorResponse, _ := json.Marshal(api.InternalError(err.Error()))
		return api.NewResponseBuilder().
			Body(string(internalErrorResponse)).
			Status(http.StatusInternalServerError).
			Build(), nil
	}

	return api.NewResponseBuilder().
		Status(http.StatusOK).
		Build(), nil
}

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	cfg = util.LoadConfig()

	hystrixConfig := hystrix.CommandConfig{
		Timeout:               cfg.HystrixTimeout,
		MaxConcurrentRequests: cfg.HystrixMaxConcurrentRequests,
		ErrorPercentThreshold: cfg.HystrixErrorPercentThreshold,
	}

	hystrix.ConfigureCommand(cfg.MetricsCommandName, hystrixConfig)

	if cfg.Stage == "local" {
		local()
	} else {
		lambda.Start(Handler)
	}
}
