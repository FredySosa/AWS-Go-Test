package domain

type Post struct {
	ID        string `json:"id" dynamodbav:"id"`
	URLToPost string `json:"urlToPost" dynamodbav:"url_to_post"`
	Text      string `json:"text" dynamodbav:"text"`
}
