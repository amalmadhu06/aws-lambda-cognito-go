package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func main() {
	lambda.Start(handler)
}

type Person struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

type ResponseBody struct {
	Message *string `json:"message"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// declare a variable for storing user details that comes in request
	var person Person

	// bind json request to `person` variable
	if err := json.Unmarshal([]byte(request.Body), &person); err != nil {
		return events.APIGatewayProxyResponse{}, nil
	}

	// create a message string for sending back response
	msg := fmt.Sprintf("Hello %v %v", *person.FirstName, *person.LastName)

	// place the message string inside the response body struct
	responseBody := ResponseBody{
		Message: &msg,
	}

	// convert response body to a json
	jBytes, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, nil
	}

	// send back response
	response := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jBytes), //use convert json to string using type casting and send as body
	}

	return response, nil
}
