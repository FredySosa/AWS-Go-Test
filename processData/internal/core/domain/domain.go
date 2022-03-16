package domain

type Post struct {
	ID         string `json:"id" dynamodbav:"id"`
	URLToImage string `json:"urlToImage" dynamodbav:"url_to_image"`
	Text       string `json:"text" dynamodbav:"text"`
}
