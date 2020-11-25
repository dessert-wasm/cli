package cmd

import (
	"context"
	"fmt"
	"reflect"

	"github.com/geospace/sac"
	"github.com/machinebox/graphql"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// PublishR : response struct for mutation `mutatePublishModule`
type PublishR struct {
	CreateModule struct { // TODO: fusion module struct
		ID int `json:"id"`
	} `json:"createModule"`
}

// PublishM : query for mutation `mutatePublishModule`
var PublishM = `
mutation ($token: String!, $description: String!, $name: String!, $isCore: Boolean!, $replacements: [ModuleReplacementInput!]!) {
	createModule(token: $token, isCore: $isCore, description: $description, name: $name, replacements: $replacements) {
		id
	}
}
`

// Module : data needed to describe user's module
type Module struct {
	name        string
	description string
	is_core     bool
	replaces    []string
}

func checkPropertyString(module *Module, propName string) bool {
	if module == nil {
		return false
	}

	r := reflect.ValueOf(module)
	s := reflect.Indirect(r).FieldByName(propName)

	return s.IsValid()
}

// TODO: we may need other variables from package.json in the future
func getModuleData(sacJSON *sac.Sac) (Module, error) {
	var mod Module

	mod.name = sacJSON.GetString(JSONNameKey)
	mod.description = sacJSON.GetString(JSONDescriptionKey)
	mod.is_core = sacJSON.GetBool("dessert.is_core")
	mod.replaces = sacJSON.GetStringSlice("dessert.replaces")

	propArr := []string{JSONNameKey, JSONDescriptionKey, JSONIsCoreKey}

	for _, elem := range propArr {
		if b := checkPropertyString(&mod, elem); !b {
			return mod, fmt.Errorf("missing %s property in package.json", elem)
		}
	}

	return mod, nil
}

type Replacement struct {
	Name string `json:"name"`
}

func publish(client *graphql.Client, token string, mod Module) error {
	var respData PublishR
	req := graphql.NewRequest(PublishM)
	ctx := context.Background()

	// setting `createModule` mutation variables
	req.Var(GQLTokenKey, token)
	req.Var(GQLNameKey, mod.name)
	req.Var(GQLDescriptionKey, mod.description)
	req.Var(GQLIsCoreKey, mod.is_core)

	rpl := []Replacement{}
	// Loop
	for i := 0; i < len(mod.replaces); i++ {
		item := Replacement{Name: mod.replaces[i]}
		rpl = append(rpl, item)
	}

	//data, err := json.Marshal(test)
	// fmt.Println("Replacements", string(data), err)
	req.Var("replacements", rpl)

	// fmt.Println(req)

	if err := client.Run(ctx, req, &respData); err != nil {
		// fmt.Println(err, respData)
		return err
	}
	return nil
}

func createPublishCmd(sacYML *sac.Sac, sacJSON *sac.Sac) *cobra.Command {
	return &cobra.Command{
		Use:   "publish",
		Short: "Publishes your project to the dessert platform.",
		Long: `[needs you to be logged in]
	Publishes your project to the dessert platform.
	You will need to run "npm publish" afterwards
	`,
		PreRun: func(cmd *cobra.Command, args []string) {
			Logger.Info(logPublishingToDessert)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			Logger.Info(logPublishingSuccess)
			Logger.Warn(warnPublishNPM)
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			if err := sacYML.ReadConfig("dessert.yml"); err != nil {
				return errors.New(errDessertYMLNeeded)
			}
			if err := sacJSON.ReadConfig("package.json"); err != nil {
				return errors.New(errPackageJSONNeeded)
			}

			mod, err := getModuleData(sacJSON)
			if err != nil {
				return err
			}

			token := sacYML.GetString("token")
			if token == "" {
				return errors.New(errGettingToken)
			}

			client := initClient()

			if err := publish(client, token, mod); err != nil {
				return errors.New(errPublishingModule)
			}
			return nil
		},
	}
}
