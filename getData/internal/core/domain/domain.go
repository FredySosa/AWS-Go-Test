package domain

import "encoding/json"

type (
	Post struct {
		ID        string `json:"id" dynamodbav:"id"`
		URLToPost string `json:"urlToPost" dynamodbav:"url_to_post"`
		Text      string `json:"text" dynamodbav:"text"`
	}

	Response struct {
		Posts []Post `json:"posts"`
	}
)

func (r Response) String() string {
	data, err := json.Marshal(r)
	if err != nil {
		return ""
	}

	return string(data)
}
