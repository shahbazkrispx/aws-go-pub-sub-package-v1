package services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

func SendToSNS(topicARn string, message string, messageAttr map[string]*sns.MessageAttributeValue)  {
	session, err := BuildSession()
	if err != nil {
		fmt.Println(err.Error())
	}
	svc := sns.New(session)

	pubMessage := &sns.PublishInput{
		MessageAttributes: messageAttr,
		Message:  aws.String(message),
		TopicArn: aws.String(topicARn),
	}
	_, err = svc.Publish(pubMessage)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}