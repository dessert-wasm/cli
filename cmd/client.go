package cmd

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"

	"github.com/machinebox/graphql"
)

// GraphQLURI has a default value, might be overriden by the environment
var GraphQLURI = "https://prod.dessert.vodka/"

// CookieJar can't fail, but you never know !
func initHTTPClient() *http.Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(fatalCookieJar)
	}
	return &http.Client{Jar: jar, Timeout: time.Second * 10}
}

func initClient() *graphql.Client {
	// Try to retrieve the GraphQL endpoint from the environment
	// if it is provided
	remote := os.Getenv("DESSERT_GRAPHQL_URI")
	if len(remote) > 0 {
		GraphQLURI = remote
	}

	httpClient := initHTTPClient()
	client := graphql.NewClient(GraphQLURI, graphql.WithHTTPClient(httpClient))
	return client
}
