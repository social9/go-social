package main

import (
	"fmt"

	"github.com/social9/go-social/config"
	"github.com/social9/go-social/twitter"
)

func main() {

	env := config.Env()

	twitter, _ := twitter.NewTwitterClient(twitter.TwitterConfig{
		TwitterConsumerKey:    env.TWConsumerKey,
		TwitterConsumerSecret: env.TWConsumerSecret,
		TwitterAccessToken:    "<User Access Token>",
		TwitterAccessSceret:   "<User Access Secret>",

		Verbosity: 0,
	})

	tweet, err := twitter.PostTweet("This is the test tweet")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Newly posted tweet:\n%+v\n", tweet)
}
