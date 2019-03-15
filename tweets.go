package main

import (
	"fmt"
	"log"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func analysis() {

	// fmt.Printf("%#v\n", hc.Analyze("Hey you worthless scumbag"))
	// fmt.Printf("%#v\n", hc.Positivity("life is very good")) //check positive max score 6
	// fmt.Printf("%#v\n", hc.Negativity("Hey you worthless scumbag")) //check neg max score -6

}

func main() {
	//// TODO: user input to be searched
	//// TODO: sentiment the how positive/negative the tweet is

	config := oauth1.NewConfig(consumerKey, consumerSecret)                   //consumerKey, consumerSecret
	token := oauth1.NewToken(accessToken, accessSecret) //accessToken, accessSecret
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter client
	client := twitter.NewClient(httpClient) //creating a new client

	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "pi",
	})

	// twet, resp, err := client.Statuses.Update("Testing my Twitter Bot.", nil)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(resp)
	fmt.Println("\n")

	log.Println(search)	//type *twitter.Search

}
