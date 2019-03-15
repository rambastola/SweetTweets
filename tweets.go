package main

// OAuth1
import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {

	config := oauth1.NewConfig(consumerKey, consumerSecret) //consumerKey, consumerSecret
	token := oauth1.NewToken(accessToken, accessSecret)     //accessToken, accessSecret
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter client
	client := twitter.NewClient(httpClient) //creating a new client
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "trump",
	})

	// twet, resp, err := client.Statuses.Update("Testing my Twitter Bot.", nil)
	if err != nil {
		fmt.Println(err)
	}

	// defer resp.Body.Close()

	fmt.Println(resp)
	fmt.Println("\n")
	fmt.Println(search)
}
