package infrastructure

import (
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"

	"github.com/kaliazur/demo-clean-architecture/v2/pkg/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const __TABLE_USER__ = "test_user"

type UserRepo struct {
	sess         *session.Session
	svc_DynamoDB *dynamodb.DynamoDB
}

var usersMap map[string]*domain.User = map[string]*domain.User{}

func (u *UserRepo) Init(config map[string]interface{}) (err error) {
	os.Setenv("AWS_ACCESS_KEY_ID", config["AWS_ACCESS_KEY_ID"].(string))
	os.Setenv("AWS_SECRET_ACCESS_KEY", config["AWS_SECRET_ACCESS_KEY"].(string))
	// os.Setenv("AWS_SESSION_TOKEN", config["AWS_SESSION_TOKEN"].(string))

	creds := credentials.NewEnvCredentials()
	creds.Get()

	u.sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
		// Endpoint:    aws.String("http://localhost:8000"),
		Credentials: creds,
	}))

	u.svc_DynamoDB = dynamodb.New(u.sess)

	return nil
}

func (u *UserRepo) CreateUser(user domain.User) (createdUser *domain.User, err error) {
	user.Id = uuid.NewString()

	// Create item in table
	tableName := __TABLE_USER__
	fmt.Printf("Inserting event in %+v\n", tableName)
	_, err = u.svc_DynamoDB.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(user.Id),
			},
			"firstname": {
				S: aws.String(user.Firstname),
			},
			"username": {
				S: aws.String(user.Username),
			},
			// "age": {
			// 	N: aws.String(user.Age),
			// },
			"is_admin": {
				BOOL: aws.Bool(user.IsAdmin),
			},
		},
	})
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetUserById(id string) (user *domain.User, err error) {

	tableName := __TABLE_USER__
	fmt.Printf("Using table: %+v\n", tableName)

	result, err := u.svc_DynamoDB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if result.Item == nil {
		fmt.Printf("Could not find customer id: %+v\n", id)
		return nil, errors.New(fmt.Sprintf("User Not Found: %s", id))
	}
	fmt.Printf("result.Item: %+v\n", result.Item)

	user = &domain.User{}
	err = dynamodbattribute.UnmarshalMap(result.Item, user)
	if err != nil {
		fmt.Printf("Error dynamodbattribute.UnmarshalMap(): %+v\n", result.Item)
		return nil, err
	}
	fmt.Printf("user: %+v\n", user)

	return
}
