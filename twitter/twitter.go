package twitter

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Config twitter instance configuration
type Config struct {
	TwitterConsumerKey    string
	TwitterConsumerSecret string
	TwitterAccessToken    string
	TwitterAccessSceret   string

	client *twitter.Client
}

// Client twitter client methods
type Client interface {
	GetTweets(count int) ([]twitter.Tweet, error)
	PostTweet(status string) (*twitter.Tweet, error)
	DeleteTweet(id string) (*twitter.Tweet, error)
	GetUserInfo() (*twitter.User, error)
}

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "(go-social)", log.Lshortfile)
}

// NewTwitterClient create a twitter client instance
func NewTwitterClient(opts Config) (Client, error) {
	// Validate parameters
	validateErr := validateOpts(opts)
	if validateErr != nil {
		logger.Println(validateErr)
		return nil, validateErr
	}

	config := oauth1.NewConfig(opts.TwitterConsumerKey, opts.TwitterConsumerSecret)
	if config == nil {
		logger.Println("Invalid Twitter Configuration")
		return nil, errors.New("Something is wrong with your twitter app")
	}
	token := oauth1.NewToken(opts.TwitterAccessToken, opts.TwitterAccessSceret)
	if token == nil {
		logger.Println("Invalid Twitter Access Token")
		return nil, errors.New("Something is wrong with your User Access Token")
	}

	// httpClient will automatically authorize http.Request's
	httpClient := config.Client(oauth1.NoContext, token)
	if httpClient == nil {
		logger.Println("Twitter Authorization Failed")
		return nil, errors.New("Twitter Authorization Failed")
	}

	// Nicer: Pass OAuth1 client to go-twitter API
	client := twitter.NewClient(httpClient)
	if client == nil {
		logger.Println("Failed to create go-twitter API Client")
		return nil, errors.New("Failed to create go-twitter API Client")
	}

	logger.Println("go-twitter API client created successfully")

	opts.client = client
	return &opts, nil
}

// GetTweets retrieve user tweets
func (s *Config) GetTweets(count int) ([]twitter.Tweet, error) {

	tweets, _, err := s.client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: count,
	})
	return tweets, err
}

// PostTweet create a tweet
func (s *Config) PostTweet(status string) (*twitter.Tweet, error) {
	tweet, _, err := s.client.Statuses.Update(status, nil)
	return tweet, err
}

// GetUserInfo retrieve user info
func (s *Config) GetUserInfo() (*twitter.User, error) {
	user, _, err := s.client.Accounts.VerifyCredentials(nil)
	return user, err
}

// DeleteTweet based on the id
func (s *Config) DeleteTweet(id string) (*twitter.Tweet, error) {
	delID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	tweet, _, err := s.client.Statuses.Destroy(delID, nil)
	return tweet, err
}

func validateOpts(opts Config) error {
	if opts.TwitterAccessSceret == "" {
		return errors.New("TwitterAccessSceret is required")
	}
	if opts.TwitterAccessToken == "" {
		return errors.New("TwitterAccessToken is required")
	}
	if opts.TwitterConsumerKey == "" {
		return errors.New("TwitterConsumerKey is required")
	}
	if opts.TwitterConsumerSecret == "" {
		return errors.New("TwitterConsumerSecret is required")
	}

	return nil
}
