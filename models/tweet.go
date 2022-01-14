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

// GetTweetsByUsername returns a list of tweets (tweet.go) by Twitter username
func GetTweetsByUsername(username string) []Tweet {
	var userId = GetTwitterUserByUsername(username).id
	var twitterApiTweet = response.ParseTweetFromJson(userId)
	var tweets []Tweet

	apiTweets := twitterApiTweet.Data

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

// GetTweetsByQuerySearch returns a list of tweets (tweet.go) by Twitter query parameter (in order to generate product score with results count)
func GetTweetsByQuerySearch(searchParameter string) []Tweet {
	var twitterApiTweet = response.ParseQuerySearchFromJson(searchParameter)
	var tweets []Tweet

	apiTweets := twitterApiTweet.Data

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
