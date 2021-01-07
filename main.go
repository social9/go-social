package main

import (
	"log"

	"github.com/social9/go-social/twitter"
)

func main() {

	twitter, _ := twitter.NewTwitterClient(twitter.Config{
		TwitterConsumerKey:    "<Twitter Consumer Key>",
		TwitterConsumerSecret: "<Twitter consumer secret>",
		TwitterAccessToken:    "<User Access Token>",
		TwitterAccessSceret:   "<User Access Secret>",
	})

	tweet, err := twitter.PostTweet("This is the test tweet")
	if err != nil {
		log.Println(err)
	}

	log.Printf("Newly posted tweet:\n%+v", tweet)
}
