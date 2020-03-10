package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func createVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Displays your dessert-cli current version number",
		Run: func(cmd *cobra.Command, args []string) {
			version()
		},
	}

	return cmd
}

func version() {
	fmt.Printf("%s\n", currentVersionCLI)
}
