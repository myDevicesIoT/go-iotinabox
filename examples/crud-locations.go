package main

import (
	"context"
	"log"

	"github.com/mydevicesiot/go-iotinabox/iotinabox"
	"golang.org/x/oauth2"
)

const (
	USERNAME = "YOUR_IOTINABOX_EMAIL"
	PASSWORD = "YOUR_IOTINABOX_PASSWORD"
)

func main() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_SECRET",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.iotinabox.com/auth/realms/iotinabox/protocol/openid-connect/auth",
			TokenURL: "https://accounts.iotinabox.com/auth/realms/iotinabox/protocol/openid-connect/token",
		},
	}

	token, err := conf.PasswordCredentialsToken(ctx, USERNAME, PASSWORD)

	if err != nil {
		log.Fatal("Could not get Token")
	}

	//log.Println("Token %s", token.AccessToken)
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token.AccessToken})
	tc := oauth2.NewClient(ctx, ts)

	tinaClient := iotinabox.NewClient(tc)

	// Get Location List
	locations, err := tinaClient.Locations.List(ctx)
	if err != nil {
		log.Fatalf("Error %s", err)
	}

	for _, location := range locations {
		log.Println(location)
	}

	// Get Single Location
	location, err := tinaClient.Locations.Get(ctx, 82)
	if err != nil {
		log.Fatalf("Error while getting single location %s", err)
	}

	log.Println(location)
}
