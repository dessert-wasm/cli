package cmd

import (
	"errors"

	"github.com/geospace/sac"
	"github.com/spf13/cobra"
)

func createReplacesCmd(sacJSON *sac.Sac) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "replaces",
		Short: "Indicate which module you want to replace",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("replaces requires at least one module to replace")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := sacJSON.ReadConfig("package.json"); err != nil {
				return errors.New(errPackageJSONNeeded)
			}
			sacJSON.Set("dessert.replaces", args)
			Logger.Infof("replacing %v", args)
			return sacJSON.WriteConfig()
		},
	}

	return cmd
}
