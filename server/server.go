package server

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/suga0828/go-twitter-bot/client"
	"github.com/suga0828/go-twitter-bot/dtos"
	"github.com/suga0828/go-twitter-bot/scheduler"
	"github.com/suga0828/go-twitter-bot/tweet"
)

// Create a new echo server
func Create() {
	client, err := client.Get(&dtos.CredentialsDTO{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	})
	if err != nil {
		log.Printf("error getting Twitter Client: %v", err)
	}

	scheduler.Run("1h", func() {
		text := "Tweet created at: " + time.Now().Format(time.Kitchen)
		tweet, _, err := tweet.Create(client, text)
		if err != nil {
			fmt.Printf("There was an error creating tweet: %v", err)
		}

		fmt.Println(tweet.Text)
	})

	e := echo.New()

	e.Logger.Fatal(e.Start(":1323"))
}
