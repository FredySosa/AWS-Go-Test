package domain

import "encoding/json"

type (
	CreationRequest struct {
		URLToPost string `json:"urlToPost"`
		Text      string `json:"text"`
	}
	Response struct {
		ID string `json:"id"`
	}
	PostToSave struct {
		ID        string `json:"id" dynamodbav:"id"`
		URLToPost string `json:"urlToPost" dynamodbav:"url_to_post"`
		Text      string `json:"text" dynamodbav:"text"`
	}
)

func (r Response) String() string {
	data, err := json.Marshal(r)
	if err != nil {
		return ""
	}

	return string(data)
}
