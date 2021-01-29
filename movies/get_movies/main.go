package main

import (
	"awsLambdaExample/movies/repositories"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type getMoviesInterface interface {
	Handle() (string, error)
}

//Handler : adapter
type Handler func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

//Handler : Function called by start
func myHandler(uc getMoviesInterface) Handler {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

		fmt.Println("Handler", "endpoint get movies")
		res, err := uc.Handle()
		if err != nil {
			return events.APIGatewayProxyResponse{Body: "error", StatusCode: 500}, nil
		}
		return events.APIGatewayProxyResponse{Body: res, StatusCode: 200}, nil
	}
}

func main() {
	moviesRepository := repositories.NewDynamoMovie()
	usecase := NewGetMovies(moviesRepository)
	lambda.Start(myHandler(usecase))
}
