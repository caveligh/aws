package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type myEvent struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

func handleHello(ev events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var reqBody myEvent

	resp := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                  "application/json",
			"Access-Control-Allow-Origin":   "*",
			"Access-Control-Alloe-Methodes": "GET,HEAD,OPTIONS,POST",
		},
	}

	err := json.Unmarshal([]byte(ev.Body), &reqBody)
	if err != nil {
		resp.StatusCode = http.StatusBadRequest
		return resp, err
	}

	resp.Body = fmt.Sprintf(`{"message":"Hola: %s %s"}`, reqBody.Name, reqBody.Lastname)
	resp.StatusCode = http.StatusOK

	return resp, nil
}

func main() {
	lambda.Start(handleHello)
}
