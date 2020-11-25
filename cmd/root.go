package cmd

import (
	"github.com/geospace/sac"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// Fs enables us to create tmp files for our tests and manage io
var Fs afero.Fs

// Logger for dessert
var Logger = logrus.New()

// CreateRoot creates cobra's root command. It contains all the other dessert-cli commands.
func CreateRoot() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "dessert-cli",
		Short: "Dessert command line tool, written in go using cobra and viper.",
		Long: `Dessert command line tool, written in go using cobra and viper.
	* Link to the project : https://dessert.dev/
	* Check out the latest news about Dessert ! https://dessert.dev/blog
	`,
		SilenceUsage:  true, // See : https://github.com/spf13/cobra/issues/340
		SilenceErrors: true,
	}

	Logger.SetFormatter(new(myFormatter))

	// Sacs
	sacYML := sac.New()
	sacYML.ChangeFS(Fs)
	sacYML.Path = "dessert.yml"
	sacYML.ConfigType = sac.YAML

	sacJSON := sac.New()
	sacJSON.ChangeFS(Fs)
	sacJSON.Path = "package.json"
	sacJSON.ConfigType = sac.JSON

	rootCmd.AddCommand(createVersionCmd())
	rootCmd.AddCommand(createLogoutCmd(sacYML))
	rootCmd.AddCommand(createLoginCmd(sacYML))
	rootCmd.AddCommand(createInitCmd(sacYML, sacJSON))
	rootCmd.AddCommand(createPublishCmd(sacYML, sacJSON))
	rootCmd.AddCommand(createReplacesCmd(sacJSON))
	rootCmd.AddCommand(createRecommendCmd())
	return rootCmd
}
