package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/shahbazkrispx/aws-go-pub-sub-package-v1/services"
)

func main() {
	fmt.Println("PUB/SUB")
}

func SendToSNS(topicARn string, message string, messageAttr map[string]*sns.MessageAttributeValue) {

	session, err := services.BuildSession()
	if err != nil {
		fmt.Println(err.Error())
	}
	svc := sns.New(session)

	pubMessage := &sns.PublishInput{
		MessageAttributes: messageAttr,
		Message:           aws.String(message),
		TopicArn:          aws.String(topicARn),
	}
	_, err = svc.Publish(pubMessage)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func SendToSQS(queueUrl string, message string, messageAttr map[string]*sqs.MessageAttributeValue) {
	session, err := services.BuildSession()
	if err != nil {
		fmt.Println(err.Error())
	}

	svc := sqs.New(session, nil)

	sendInput := &sqs.SendMessageInput{
		MessageAttributes: messageAttr,
		MessageBody:       aws.String(message),
		QueueUrl:          aws.String(queueUrl),
	}

	_, er := svc.SendMessage(sendInput)
	if er != nil {
		fmt.Println(er.Error())
		return
	}
}

func SubscribeSQS(queueUrl string, cancel <-chan os.Signal) ([]*sqs.Message, error) {
	awsSession, err := services.BuildSession()
	if err != nil {
		fmt.Println(err.Error())
	}
	svc := sqs.New(awsSession, nil)

	messages, err := receiveMessages(svc, queueUrl)
	if messages == nil && len(messages) == 0 {
		return nil, err
	}

	//select {
	//case <-cancel:
	//	return nil, errors.New("")
	//case <-time.After(100 * time.Millisecond):
	//}
	return messages, nil
}

func receiveMessages(svc *sqs.SQS, queueUrl string) ([]*sqs.Message, error) {

	receiveMessagesInput := &sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            aws.String(queueUrl),
		MaxNumberOfMessages: aws.Int64(10), // max 10
		WaitTimeSeconds:     aws.Int64(3),  // max 20
		VisibilityTimeout:   aws.Int64(20), // max 20
	}

	receiveMessageOutput, err :=
		svc.ReceiveMessage(receiveMessagesInput)

	if err != nil {
		return nil, err
	}

	if receiveMessageOutput == nil || len(receiveMessageOutput.Messages) == 0 {
		return nil, errors.New("Message not found.")
	}

	return receiveMessageOutput.Messages, nil
}

func DeleteMessage(svc *sqs.SQS, queueUrl string, handle *string) {
	delInput := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueUrl),
		ReceiptHandle: handle,
	}
	_, err := svc.DeleteMessage(delInput)

	if err != nil {
		fmt.Println("Delete Error", err)
		return
	}
}
