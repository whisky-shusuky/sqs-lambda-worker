package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type sesMessageBody struct {
	Region    string `json:"region"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
	Sender    string `json:"sender"`
	Recipient string `json:"reciptient"`
}

type SQSSendMessageAPI interface {
	GetQueueUrl(ctx context.Context,
		params *sqs.GetQueueUrlInput,
		optFns ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error)

	SendMessage(ctx context.Context,
		params *sqs.SendMessageInput,
		optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}

func GetQueueURL(c context.Context, api SQSSendMessageAPI, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	return api.GetQueueUrl(c, input)
}

func SendMsg(c context.Context, api SQSSendMessageAPI, input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	return api.SendMessage(c, input)
}

func main() {
	ctx := context.Background()
	region := "ap-northeast-1"
	// SQS QueueのSender アドレスを入力
	queueURL := ""
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		fmt.Println("error")
		return
	}
	client := sqs.NewFromConfig(cfg)

	// Recipient に送信先のメールアドレスを指定
	// Sender に送信元のメールアドレスを指定する(SESであらかじめドメイン認証しておくことが必要)
	data := sesMessageBody{Region: "ap-northeast-1", Subject: "件名", Sender: "", Recipient: "", Message: "test send"}
	dataJson, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error")
		return
	}

	sMInput := &sqs.SendMessageInput{
		DelaySeconds: 10,
		MessageBody:  aws.String(string(dataJson)),
		QueueUrl:     &queueURL,
	}

	resp, err := SendMsg(context.TODO(), client, sMInput)
	if err != nil {
		fmt.Println("Got an error sending the message:")
		fmt.Println(err)
		return
	}
	fmt.Println("Sent message with ID: " + *resp.MessageId)
}
