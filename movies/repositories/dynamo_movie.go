package repositories

import (
	"awsLambdaExample/movies/models"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
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

	table := os.Getenv("MOVIES_TABLE")

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("not", "empty", ""),
		DisableSSL:  aws.Bool(true),
		Region: aws.String(os.Getenv("REGION")),
		Endpoint:    aws.String("http://" + os.Getenv("LOCALSTACK_HOSTNAME")+":4566"),
	})

	if err != nil {
		fmt.Println("Get session failed:")
		fmt.Println((err.Error()))
		return "",err
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
		return "",err
	}
	Movies := []models.Movie{}

	for _, i := range result.Items {
		item := models.Movie{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			return "",err
		}
		Movies = append(Movies, item)
	}


	json, err := json.Marshal(Movies)
	if err != nil {
		fmt.Println("Got error marshalling:")
		fmt.Println(err.Error())
		return "",err
	}
	return string(json),nil


}
