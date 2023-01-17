package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type myEvent struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

func handleHello(ev myEvent) (string, error) {
	return fmt.Sprintf("Hola: %s %s", ev.Name, ev.Lastname), nil
}

func main() {
	lambda.Start(handleHello)
}
