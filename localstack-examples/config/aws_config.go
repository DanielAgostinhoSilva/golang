package config

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type AWSClient struct {
	session *session.Session
}

func NewAWSClient() *AWSClient {
	cfg := aws.Config{
		Credentials: credentials.NewStaticCredentials("access_key_id", "secret_access_key", "session_token"),
		Endpoint:    aws.String("http://localhost:4566"), // URL do LocalStack SNS
		Region:      aws.String("us-east-1"),             // Região (não é usada com LocalStack)
	}
	sess := session.Must(session.NewSession(&cfg))
	return &AWSClient{
		session: sess,
	}
}

func (c *AWSClient) ListSQSQueues() error {
	svc := sqs.New(c.session)
	result, err := svc.ListQueues(nil) // for all queues
	if err != nil {
		return err
	}

	for _, url := range result.QueueUrls {
		fmt.Println(*url)
	}

	return nil
}

func (c *AWSClient) ListSNSTopics() error {
	svc := sns.New(c.session)
	result, err := svc.ListTopics(nil) // for all topics
	if err != nil {
		return err
	}

	for _, topic := range result.Topics {
		fmt.Println(*topic.TopicArn)
	}

	return nil
}
