package tweet

import (
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
)

// Create a tweet with given text
func Create(client *twitter.Client, text string) (*twitter.Tweet, *http.Response, error) {
	tweet, resp, err := client.Statuses.Update(text, nil)

	return tweet, resp, err
}
