package cmd

import (
	"context"

	"github.com/geospace/sac"

	"github.com/machinebox/graphql"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// DeleteTokenM : query for mutation `deleteToken`
var DeleteTokenM = `
mutation($token: String) {
	deleteToken(token: $token)
}
`

// DeleteTokenR : response struct for mutation `deleteToken`
type DeleteTokenR struct {
	DeleteToken bool `json:"deleteToken"`
}

func createLogoutCmd(sacYML *sac.Sac) *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Logs you out from dessert.",
		PreRun: func(cmd *cobra.Command, args []string) {
			Logger.Info(logLoggingYouOut)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			Logger.Info(logOutSuccess)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := sacYML.ReadConfig(dessertYML); err != nil {
				return err
			}
			if err := logout(sacYML); err != nil {
				return errors.Wrap(err, errLoggingOut)
			}
			return nil
		},
	}
}

func logout(sacYML *sac.Sac) error {
	token := sacYML.GetString("token")

	if token == "" {
		Logger.Info("already logged out")
		return nil
	}

	client := initClient()

	var respData DeleteTokenR
	req := graphql.NewRequest(DeleteTokenM)
	ctx := context.Background()

	req.Var("token", token)
	if err := client.Run(ctx, req, &respData); err != nil {
		return err
	}

	if !respData.DeleteToken {
		return errors.New(errBadServerResp)
	}

	// removing token from dessert_config.yml
	sacYML.Set("token", "")
	return sacYML.WriteConfig()
}
