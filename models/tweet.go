package models

import (
	"everywheretew.it/main/models/response"
	"fmt"
	"strings"
)

type Tweet struct {
	id             string
	authorId       string
	createdAt      string
	text           string
	conversationId string
	source         string
}

func (t Tweet) String() string {
	return fmt.Sprintf("\n{"+
		"\n\tid: %s,"+
		"\n\tauthorId: %s, "+
		"\n\tcreatedAt: %s,"+
		"\n\ttext: %s,"+
		"\n\tconversationId: %s,"+
		"\n\tsource: %s"+
		"\n}", t.id, t.authorId, t.createdAt, strings.Replace(t.text, "\n", " ", -1), t.conversationId, t.source)
}

// GetTweetsByUsername return a list of tweets (tweet.go) by Twitter username
func GetTweetsByUsername(username string) []Tweet {
	var twitterApiUser = response.ParseResponseFromJson(username)
	var tweets []Tweet

	apiTweets := twitterApiUser.Includes.Tweets

	for _, tweet := range apiTweets {
		tweets = append(tweets, Tweet{
			id:             tweet.Id,
			authorId:       tweet.AuthorId,
			createdAt:      tweet.CreatedAt,
			text:           tweet.Text,
			conversationId: tweet.ConversationId,
			source:         tweet.Source,
		})
	}

	return tweets
}
