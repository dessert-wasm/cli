package cmd

import (
	"context"

	"github.com/machinebox/graphql"
)

/* Utility functions to get token for auth */

// GetTokenR : response struct for mutation `mutateTokenCreate`
type GetTokenR struct {
	Token string `json:"createToken"`
}

// GetTokenM : query for mutation `mutateTokenCreate`
const GetTokenM = `
mutation {
	createToken(description: "CLI_login")
}
`

func getToken(client *graphql.Client) (token string, err error) {
	var respData GetTokenR
	req := graphql.NewRequest(GetTokenM)

	ctx := context.Background()
	if err := client.Run(ctx, req, &respData); err != nil {
		return "", err
	}

	return respData.Token, nil
}
