package common

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

var C = cron.New()

func init() {
	//C.AddFunc("@every 10s", tweetJob)
	//C.Start()
}

func tweetJob() {
	fmt.Println("test")
	//TweetAnalysis()
}
