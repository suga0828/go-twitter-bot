package server

import (
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/suga0828/go-twitter-bot/client"
	"github.com/suga0828/go-twitter-bot/dtos"
)

// Create a new echo server
func Create() {
	_, err := client.Get(&dtos.CredentialsDTO{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	})
	if err != nil {
		log.Printf("error getting Twitter Client: %v", err)
	}

	e := echo.New()

	e.Logger.Fatal(e.Start(":1323"))
}
