package pkg

import (
	"errors"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Tweeter struct {
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
}

func (t *Tweeter) Tweet(tweet string) error {
	config := oauth1.NewConfig(t.ConsumerKey, t.ConsumerSecret)
	token := oauth1.NewToken(t.AccessToken, t.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	_, resp, err := client.Statuses.Update(tweet, nil)
	if err != nil {
		return err
	}

	if !is2xx(resp.StatusCode) {
		return errors.New(fmt.Sprintf("there was a problem in sending the tweet. HTTP Error Code: %v", resp.StatusCode))
	}

	return nil
}

func is2xx(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
