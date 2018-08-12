package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serinth/serverless-emailer/api"
	"github.com/serinth/serverless-emailer/validators"
	"os"

	"net/http"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var req api.SendEmailRequest

	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return api.NewResponseBuilder().Status(http.StatusBadRequest).Build(), nil
	}

	requestErrors := validators.GetSendEmailRequestErrors(req)
	if len(requestErrors) > 0 {
		invalidRequestResponse, _ := json.Marshal(api.InvalidEmailRequest(requestErrors))
		return api.NewResponseBuilder().
			Body(string(invalidRequestResponse)).
			Status(http.StatusBadRequest).
			Build(), nil
	}

	//TODO actually send the email here with fallback and circuitbreaker

	return api.NewResponseBuilder().
		Body("OK").
		Status(http.StatusOK).
		Build(), nil
}

func main() {
	stage := os.Getenv("STAGE")

	if len(stage) == 0 {
		panic("Error! Mandatory STAGE environment variable not set!")
	}

	if stage == "local" {
		local()
	} else {
		lambda.Start(Handler)
	}
}
