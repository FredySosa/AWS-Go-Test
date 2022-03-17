package ports

import "github.com/aws/aws-sdk-go/service/sns"

type SNSPort interface {
	Publish(input *sns.PublishInput) (*sns.PublishOutput, error)
}
