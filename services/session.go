package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/shahbazkrispx/aws-go-pub-sub-package-v1/config"
	"github.com/shahbazkrispx/aws-go-pub-sub-package-v1/models"
	"log"
)

func BuildSession() (*session.Session, error) {
	creds := getCredentials()

	sessionConfig := &aws.Config{
		Region: aws.String(creds.Region),
		Credentials: credentials.NewStaticCredentials(creds.AccessKey, creds.SecretKey, ""),
	}

	sess, err := session.NewSession(sessionConfig)
	if err != nil {
		log.Println("Error Establishing session")
		return nil, err
	}
	return sess, nil
}


func getCredentials() models.Credentials {
	return models.Credentials{
		AccessKey: config.Env("AWS_ACCESS_KEY"),
		SecretKey: config.Env("AWS_SECRET"),
		Region: config.Env("AWS_REGION"),
	}
}