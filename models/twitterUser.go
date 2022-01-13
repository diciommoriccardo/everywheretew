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
	tweets        []Tweet
}

func (t TwitterUser) String() string {
	return fmt.Sprintf("\n{"+
		"\n\tid: %s, "+
		"\n\tname: %s, "+
		"\n\tusername: %s"+
		"\n\tpinnedTweetId: %s"+
		"\n\temail: %s"+
		"\n\ttweets: %s"+
		"\n}", t.id, t.name, t.username, t.pinnedTweetId, t.email, t.tweets)
}

func GetTwitterUserByUsername(username string) TwitterUser {
	var twitterApiUser = response.ParseResponseFromJson(username)
	var tweets = GetTweetsByUsername(username)

	return TwitterUser{
		id:            twitterApiUser.Data.Id,
		pinnedTweetId: twitterApiUser.Data.PinnetTweetId,
		name:          twitterApiUser.Data.Name,
		username:      twitterApiUser.Data.Username,
		email:         "",
		tweets:        tweets,
	}
}
