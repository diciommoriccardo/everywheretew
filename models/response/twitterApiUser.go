package response

import (
	"encoding/json"
	"log"

	"everywheretew.it/main/common"
)

type TwitterApiUser struct {
	Data struct {
		Username      string `json:"username"`
		PinnetTweetId string `json:"pinned_tweet_id"`
		Name          string `json:"name"`
		Id            string `json:"id"`
	} `json:"data"`
	Includes struct {
		Tweets []struct {
			AuthorId       string `json:"author_id"`
			Id             string `json:"id"`
			CreatedAt      string `json:"created_at"`
			Text           string `json:"text"`
			ConversationId string `json:"conversation_id"`
			Source         string `json:"source"`
		} `json:"tweets"`
	} `json:"includes"`
}

func ParseResponseFromJson(username string) TwitterApiUser {
	url := "https://api.twitter.com/2/users/by/username/" + username + "?expansions=pinned_tweet_id&tweet.fields=author_id,conversation_id,created_at,entities,geo,id,in_reply_to_user_id,lang,possibly_sensitive,referenced_tweets,source,text,withheld"
	var responseData []byte = common.GetRequest(url)

	var responseObject TwitterApiUser

	err := json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatal(err)
		return TwitterApiUser{}
	}

	return responseObject
}
