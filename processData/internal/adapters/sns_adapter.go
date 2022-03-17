package adapters

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/FredySosa/AWS-Go-Test/processData/internal/core/domain"
	"github.com/FredySosa/AWS-Go-Test/processData/internal/ports"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const insertType = "INSERT"

type Handler struct {
	SNSService ports.SNSServicePort
}

func NewSNSHandler(ss ports.SNSServicePort) Handler {
	return Handler{
		SNSService: ss,
	}
}

func (h Handler) ProcessRequest(ctx context.Context, e events.DynamoDBEvent) error {
	posts := make([]domain.Post, 0)
	for _, record := range e.Records {
		if !strings.EqualFold(record.EventName, insertType) {
			continue
		}

		post, err := parsePost(record.Change.NewImage)
		if err != nil {
			log.Println(err)
			return nil
		}
		posts = append(posts, post)
	}

	if err := h.SNSService.PublishMessage(ctx, posts); err != nil {
		return err
	}

	return nil
}

func parsePost(newImage map[string]events.DynamoDBAttributeValue) (domain.Post, error) {
	dbAttrMap := make(map[string]*dynamodb.AttributeValue)

	for k, v := range newImage {

		var dbAttr dynamodb.AttributeValue

		bytes, err := v.MarshalJSON()
		if err != nil {
			return domain.Post{}, err
		}

		if err = json.Unmarshal(bytes, &dbAttr); err != nil {
			return domain.Post{}, err
		}
		dbAttrMap[k] = &dbAttr
	}

	var post domain.Post
	if err := dynamodbattribute.UnmarshalMap(dbAttrMap, &post); err != nil {
		return domain.Post{}, err
	}
	return post, nil
}
