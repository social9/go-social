<p align="center">
  <a href="https://pkg.go.dev/github.com/social9/go-social/?tab=doc">
    <img src="https://img.shields.io/badge/%F0%9F%93%9A%20godoc-pkg-00ACD7.svg?color=00ACD7&style=flat">
  </a>
  <a href="https://goreportcard.com/report/github.com/social9/go-social">
    <img src="https://img.shields.io/badge/%F0%9F%93%9D%20goreport-A%2B-75C46B">
  </a>
  <a href="https://gocover.io/github.com/social9/go-social">
    <img src="https://img.shields.io/badge/coverage-0%25-orange">
  </a>
</p>

# go-social (in Development)

This is a generic project for integrating all the social profiles.

## Quick Start

```
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

	tweets, err := twitter.PostTweet("This is the test tweet")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Newly posted tweet:\n%+v\n", tweets)
}
```

## Setup

- Clone the repo `git clone git@github.com:social9/go-social`
- Congigurable enviroments variables
| Parameter              | Description  | Default | Allowed |
  |------------------------|--------------|---------|---------|
  |`TWConsumerKey`     |Twitter Consumer Key|`""`|`string`|
  |`TWConsumerSecret` |Twitter Consumer Secret|`""`|`string`|


## Contribution Guidelines

- Fork this repo to your GitHub account
- You can either create an issue or pick from the existing and seek maintainers' attention before developement
- Your _Pull Request_ branch must be rebased with the `master` branch i.e. have a linear history
- One or more maintainers will review your PR once associated to an issue.

> Do append the issue ID in the pull request title e.g. **Implemented a functionality closes #20** where **20** is the issue number

## License

MIT
