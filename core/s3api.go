package core

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func AwsSession() *session.Session {

	os.Setenv("AWS_PROFILE", "test-account")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewSharedCredentials("", "test-account"),
	})
	if err != nil {
		log.Fatal("error %v", err)
	}
	return sess
}
