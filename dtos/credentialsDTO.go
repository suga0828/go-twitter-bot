package dtos

// CredentialsDTO stores all of our access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.
type CredentialsDTO struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}
