package models

import (
	"everywheretew.it/main/models/response"
	"fmt"
)

type TwitterUser struct {
	id            string
	pinnedTweetId string
	name          string
	username      string
	email         string
	pinnedTweets  []Tweet
}

func (t TwitterUser) String() string {
	return fmt.Sprintf("\n{"+
		"\n\tid: %s, "+
		"\n\tname: %s, "+
		"\n\tusername: %s"+
		"\n\tpinnedTweetId: %s"+
		"\n\temail: %s"+
		"\n\tpinnedTweets: %s"+
		"\n}", t.id, t.name, t.username, t.pinnedTweetId, t.email, t.pinnedTweets)
}

func GetTwitterUserByUsername(username string) TwitterUser {
	var twitterApiUser = response.ParseUserFromJson(username)
	var apiPinnedTweets = twitterApiUser.Includes.PinnedTweets
	var pinnedTweets []Tweet

	for _, tweet := range apiPinnedTweets {
		pinnedTweets = append(pinnedTweets, Tweet{
			id:             tweet.Id,
			authorId:       tweet.AuthorId,
			createdAt:      tweet.CreatedAt,
			text:           tweet.Text,
			conversationId: tweet.ConversationId,
			source:         tweet.Source,
		})
	}

	return TwitterUser{
		id:            twitterApiUser.Data.Id,
		pinnedTweetId: twitterApiUser.Data.PinnetTweetId,
		name:          twitterApiUser.Data.Name,
		username:      twitterApiUser.Data.Username,
		email:         "",
		pinnedTweets:  pinnedTweets,
	}
}
