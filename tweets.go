package main

import (
	"bufio"
	"fmt"
	"os"

	hc "cirello.io/HumorChecker"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func analysis(tweets []twitter.Tweet) map[string]twitter.Tweet {
	/*
		Sentiment the tweets
	*/
	m := make(map[string]twitter.Tweet) //map to store the tweet its analysis

	for i := 0; i < len(tweets); i++ {
		analyzed := fmt.Sprintf("%#v\n", hc.Positivity(tweets[i].Text))
		m[analyzed] = tweets[i]
	}
	return m
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to be searched on twitter: ")
	input, _ := reader.ReadString('\n') //holds the string of user input

	config := oauth1.NewConfig("consumerKey", "consumerSecret") //consumerKey, consumerSecret
	token := oauth1.NewToken("accessToken", "accessSecret")     //accessToken, accessSecret
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient) //creating a new client

	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{ //get the tweets
		Query: input, Count: 20,
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(analysis(search.Statuses))

}
