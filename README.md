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
- Your _Pull Request_ branch must be rebased with the `dev` branch i.e. have a linear history
- One or more maintainers will review your PR once associated to an issue.

> Do append the issue ID in the pull request title e.g. **Implemented a functionality closes #20** where **20** is the issue number

## License

MIT
