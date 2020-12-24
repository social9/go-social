package twitter

import (
	"errors"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"

	libLogger "github.com/social9/go-queues/lib/logger"
)

type TwitterConfig struct {
	TwitterConsumerKey    string
	TwitterConsumerSecret string
	TwitterAccessToken    string
	TwitterAccessSceret   string

	// 0-5 : debug, info, warn, error, fatal
	Verbosity int

	client *twitter.Client
	logger libLogger.Logger
}

type TwitterClient interface {
	GetTweets(count int) ([]twitter.Tweet, error)
	PostTweet(status string) (*twitter.Tweet, error)
	GetUserInfo() (*twitter.User, error)
}

func NewTwitterClient(opts TwitterConfig) (TwitterClient, error) {

	logger := libLogger.NewLogger(libLogger.Config{Level: opts.Verbosity})
	// Validate parameters
	validateErr := validateOpts(opts)
	if validateErr != nil {
		logger.Debug(validateErr)
		return nil, validateErr
	}

	config := oauth1.NewConfig(opts.TwitterConsumerKey, opts.TwitterConsumerSecret)
	if config == nil {
		logger.Debug("Invalid Twitter Configuration")
		return nil, errors.New("Something is wrong with your twitter app")
	}
	token := oauth1.NewToken(opts.TwitterAccessToken, opts.TwitterAccessSceret)
	if token == nil {
		logger.Debug("Invalid Twitter Access Token")
		return nil, errors.New("Something is wrong with your User Access Token")
	}

	// httpClient will automatically authorize http.Request's
	httpClient := config.Client(oauth1.NoContext, token)
	if httpClient == nil {
		logger.Debug("Twitter Authorization Failed")
		return nil, errors.New("Twitter Authorization Failed")
	}

	// Nicer: Pass OAuth1 client to go-twitter API
	client := twitter.NewClient(httpClient)
	if client == nil {
		logger.Debug("Failed to create go-twitter API Client")
		return nil, errors.New("Failed to create go-twitter API Client")
	}

	logger.Info("go-twitter API client created successfully")

	opts.client = client
	opts.logger = logger

	return &opts, nil
}

func (s *TwitterConfig) GetTweets(count int) ([]twitter.Tweet, error) {

	tweets, _, err := s.client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: count,
	})
	return tweets, err
}

func (s *TwitterConfig) PostTweet(status string) (*twitter.Tweet, error) {
	tweet, _, err := s.client.Statuses.Update(status, nil)
	return tweet, err
}

func (s *TwitterConfig) GetUserInfo() (*twitter.User, error) {
	user, _, err := s.client.Accounts.VerifyCredentials(nil)
	return user, err
}

func validateOpts(opts TwitterConfig) error {
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
