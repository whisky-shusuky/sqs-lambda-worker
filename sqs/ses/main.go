package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

type sesMessageBody struct {
	Region    string `json:"region"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
	Sender    string `json:"sender"`
	Recipient string `json:"reciptient"`
}

func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		var res sesMessageBody
		err := json.Unmarshal([]byte(message.Body), &res)
		if err != nil {
			fmt.Println("Payload is unable to parse to json")
			return err
		}
		sendSES(&res)
		fmt.Printf("Send Following message. \n")
		fmt.Printf("Region:[%s]! \n", res.Region)
		fmt.Printf("Subject:[%s]! \n", res.Subject)
		fmt.Printf("Sender:[%s]! \n", res.Sender)
		fmt.Printf("Recipient:[%s]! \n", res.Recipient)
	}

	return nil
}

func sendSES(payload *sesMessageBody) error {
	ctx := context.Background()

	// sdk API Client 作成
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(payload.Region))
	if err != nil {
		fmt.Println("error")
		return err
	}
	client := sesv2.NewFromConfig(cfg)

	// SES API に投げ込むパラメタを作る
	input := &sesv2.SendEmailInput{
		FromEmailAddress: &payload.Sender,
		Destination: &types.Destination{
			ToAddresses: []string{payload.Recipient}, // 配列なので複数指定可能
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Data: &payload.Message,
					},
				},
				Subject: &types.Content{
					Data: &payload.Subject, // 件名
				},
			},
		},
	}

	// メール送信
	res, err := client.SendEmail(ctx, input)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(res.MessageId)
	fmt.Println("success!")
	return nil

}

func main() {
	lambda.Start(Handler)
}
