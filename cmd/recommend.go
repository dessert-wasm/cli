package cmd

import (
	"context"
	"errors"

	"github.com/machinebox/graphql"
	"github.com/spf13/cobra"
)

var recommendQ = `
query($dependencies: [JSDependencyInput!]!) {
  recommend(dependencies: $dependencies) {name}
}
`

type module struct {
	Name string `json:"name"`
}

type recommendR struct {
	Recommend []module `json:"recommend"`
}

func createRecommendCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recommend",
		Short: "Lists dependencies you can replace with Dessert",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("recommend requires at least one module to give you recommendations")
			}
			return recommend(args)
		},
	}

	return cmd
}

func displayRecommendations(modules []module) {
	if len(modules) == 0 {
		Logger.Info("No recommendations found")
		return
	}
	Logger.Infof("Found %d module(s):", len(modules))

	for i := 0; i < len(modules); i++ {
		Logger.Info(modules[i].Name)
	}
}

func fillDeps(args []string) []module {
	var dependencies []module

	for i := 0; i < len(args); i++ {
		m := module{Name: args[i]}
		dependencies = append(dependencies, m)
	}

	return dependencies
}

func recommend(args []string) error {
	client := initClient()

	var respData recommendR

	req := graphql.NewRequest(recommendQ)
	ctx := context.Background()

	Logger.Infof("Finding replacements for %v...", args)

	dependencies := fillDeps(args)
	req.Var(GQLDependencies, dependencies)

	if err := client.Run(ctx, req, &respData); err != nil {
		// fmt.Println(err)
		return err
	}

	displayRecommendations(respData.Recommend)

	return nil
}
