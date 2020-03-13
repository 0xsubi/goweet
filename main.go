package main

import (
	"flag"
	"fmt"
	"github.com/sudhabindu1/goweet/pkg"
	"os"
	"strings"
)

func main()  {
	var handle, tweet, consumerKey, consumerSecret, accessToken, accessTokenSecret string
	flag.StringVar(&handle,"handle", os.Getenv("GOWEET_HANDLE"), "your twitter handle")
	flag.StringVar(&tweet,"tweet", "", "The tweet to be sent")
	flag.StringVar(&consumerKey, "consumer_key", os.Getenv("GOWEET_CONSUMER_KEY"), "consumer key for your twitter app")
	flag.StringVar(&consumerSecret, "consumer_secret", os.Getenv("GOWEET_CONSUMER_SECRET"), "consumer secret for your twitter app")
	flag.StringVar(&accessToken, "access_token", os.Getenv("GOWEET_ACCESS_TOKEN"), "consumer key for your twitter app")
	flag.StringVar(&accessTokenSecret, "access_token_secret", os.Getenv("GOWEET_ACCESS_TOKEN_SECRET"), "consumer key for your twitter app")

	flag.Parse()

	if tweet == "" {
		fmt.Println("tweet can't be empty")
		os.Exit(1)
	}

	if !strings.HasPrefix(handle, "@") {
		handle = "@" + handle
	}

	t := pkg.Tweeter{
		ConsumerKey:consumerKey,
		ConsumerSecret:consumerSecret,
		AccessToken:accessToken,
		AccessTokenSecret:accessTokenSecret,
	}

	err := t.Tweet(tweet)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Your Tweet was posted successfully")

}
