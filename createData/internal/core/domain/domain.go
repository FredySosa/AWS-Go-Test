package domain

import "encoding/json"

type (
	CreationRequest struct {
		URLToImage string `json:"urlToImage"`
		Text       string `json:"text"`
	}
	Response struct {
		ID string `json:"id"`
	}
	PostToSave struct {
		ID         string `json:"id" dynamodbav:"id"`
		URLToImage string `json:"urlToImage" dynamodbav:"url_to_image"`
		Text       string `json:"text" dynamodbav:"text"`
	}
)

func (r Response) String() string {
	data, err := json.Marshal(r)
	if err != nil {
		return ""
	}

	return string(data)
}
