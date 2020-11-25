package cmd

import (
	"context"
	"dessert-cli/prompt"

	"github.com/geospace/sac"
	"github.com/machinebox/graphql"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// LoginR : response struct for mutation `login`
type LoginR struct {
	Login struct {
		ID int `json:"id"`
	} `json:"login"`
}

// LoginM : query for mutation `login`
var LoginM = `
mutation ($email: String!, $password: String!) {
	login(email: $email, password: $password, remember: true) { id }
}
`

func login(client *graphql.Client, credentials prompt.LoginA) error {
	req := graphql.NewRequest(LoginM)
	ctx := context.Background()

	req.Var("email", credentials.Username)
	req.Var("password", credentials.Password)

	var respData LoginR
	if err := client.Run(ctx, req, &respData); err != nil {
		return err
	}
	return nil
}

func getCredentials(answers *prompt.LoginA) error {
	return prompt.LoginPrompt(answers)
}

func createLoginCmd(sacYML *sac.Sac) *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Logs you into dessert",
		PreRun: func(cmd *cobra.Command, args []string) {
			Logger.Info(logLoggingYouIn)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			Logger.Info(logInSuccess)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := sacYML.ReadConfig(dessertYML); err != nil {
				return errors.New(warnDessertConfig404)
			}
			client := initClient()
			credentials := prompt.LoginA{}

			if err := getCredentials(&credentials); err != nil {
				return err
			}

			if err := login(client, credentials); err != nil {
				return errors.New(errLogin)
			}

			// grabbing user token and saving it to config file
			token, err := getToken(client)
			if err != nil {
				return errors.New(errGetToken)
			}

			sacYML.Set("token", token)
			if err := sacYML.WriteConfig(); err != nil {
				return errors.New(errWriteToken)
			}
			Logger.Info(logTokenSaved)
			return nil
		},
	}
}
