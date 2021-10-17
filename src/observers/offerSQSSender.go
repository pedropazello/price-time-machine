package observers

import (
	"encoding/json"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joho/godotenv"
	"github.com/pedropazello/price-time-machine/infra"
	"github.com/pedropazello/price-time-machine/src/models"
)

type OfferSQSSender struct {
}

func (o *OfferSQSSender) Execute(offer models.Offer) error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	sqsClient := infra.NewSQSClient()

	queueUrl := os.Getenv("SQS_OFFER_PROCESSED_URL")

	offerJson, err := json.Marshal(offer)

	if err != nil {
		return err
	}

	_, err = sqsClient.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(string(offerJson)),
		QueueUrl:     &queueUrl,
	})

	return err
}
