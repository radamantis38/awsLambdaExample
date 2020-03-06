package repositories

import (
	"awsLambdaExample/movies/models"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
)

type DynamoMovie struct {
}

func NewDynamoMovie() *DynamoMovie {
	return &DynamoMovie{}
}

func (rep *DynamoMovie) GetAll() (string, error) {
	//Inicia envio de mensaje a la cola
	table := os.Getenv("MOVIES_TABLE")

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION"))},
	)

	if err != nil {
		fmt.Println("Get session failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}
	// Create a SQS service client.
	svc := dynamodb.New(sess)

	// snippet-start:[dynamodb.go.scan_items.call]
	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName:                 aws.String(table),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}
	Movies := []models.Movie{}

	for _, i := range result.Items {
		item := models.Movie{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		Movies = append(Movies, item)
	}


	json, err := json.Marshal(Movies)
	if err != nil {
		fmt.Println("Got error marshalling:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(json),nil
}
