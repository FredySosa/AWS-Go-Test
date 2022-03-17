package services

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"

	"github.com/FredySosa/AWS-Go-Test/processData/internal/core/domain"
	"github.com/FredySosa/AWS-Go-Test/processData/internal/ports"
)

type SNSService struct {
	SNS   ports.SNSPort
	topic string
}

func NewSNSService(s ports.SNSPort, topic string) SNSService {
	return SNSService{
		SNS:   s,
		topic: topic,
	}
}

func (ss SNSService) PublishMessage(ctx context.Context, posts []domain.Post) error {
	for _, post := range posts {
		dataToSend := fmt.Sprintf(template, post.Text, post.URLToPost)
		_, err := ss.SNS.Publish(&sns.PublishInput{
			Message:  aws.String(dataToSend),
			TopicArn: aws.String(ss.topic),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

const template = "%s.\nPara más información da clic en el siguiente link: %s"
