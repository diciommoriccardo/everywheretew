package response

import (
	"encoding/json"
	"everywheretew.it/main/common"
	"log"
	url3 "net/url"
)

type TwitterApiTweet struct {
	Data []struct {
		AuthorId       string `json:"author_id"`
		Id             string `json:"id"`
		CreatedAt      string `json:"created_at"`
		Text           string `json:"text"`
		ConversationId string `json:"conversation_id"`
		Source         string `json:"source"`
	} `json:"data"`
}

func ParseTweetFromJson(userId string) TwitterApiTweet {
	url := "https://api.twitter.com/2/users/" + userId + "/tweets?tweet.fields=author_id,conversation_id,created_at,source"

	var responseData []byte = common.GetRequest(url)

	var responseObject TwitterApiTweet

	err := json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatal(err)
		return TwitterApiTweet{}
	}

	return responseObject
}

func ParseQuerySearchFromJson(queryParameter string) TwitterApiTweet {
	url, err := url3.Parse("https://api.twitter.com/2/tweets/search/recent")

	params := url3.Values{}
	params.Add("query", queryParameter)
	url.RawQuery = params.Encode()
	var responseData []byte = common.GetRequest(url.String())

	var responseObject TwitterApiTweet

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatal(err)
		return TwitterApiTweet{}
	}

	return responseObject
}
