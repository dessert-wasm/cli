package cmd

import (
	"dessert-cli/prompt"
	"errors"

	"github.com/geospace/sac"
	"github.com/spf13/cobra"
)

func createInitCmd(sacYML *sac.Sac, sacJSON *sac.Sac) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Makes your folder dessert-ready",
		PreRun: func(cmd *cobra.Command, args []string) {
			Logger.Info(logMakingFolderDReady)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			Logger.Info(logSuccessHappyDessert)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := sacJSON.ReadConfig("package.json"); err != nil {
				return errors.New(errPackageJSONNeeded)
			}
			if err := preparePJSON(sacJSON); err != nil {
				return err
			}
			if err := prepareDYML(sacYML); err != nil {
				return err
			}
			return nil
		},
	}
}

func inKeywords(word string, list []string) bool {
	for _, w := range list {
		if w == word {
			return true
		}
	}
	return false
}

func addDessertToKeywords(sacJSON *sac.Sac) error {
	keywords := sacJSON.GetStringSlice(JSONKeywordsKey)

	if inKeywords(JSONDessertKey, keywords) {
		return nil
	}

	keywords = append(keywords, JSONDessertKey)
	sacJSON.Set(JSONKeywordsKey, keywords)

	if err := sacJSON.WriteConfig(); err != nil {
		return err
	}

	return nil
}

func promptProjectType(answers *prompt.InitA) error {
	return prompt.InitPrompt(answers)
}

func setCore(sacJSON *sac.Sac) error {
	var answers prompt.InitA

	if err := promptProjectType(&answers); err != nil {
		return errors.New(errPrompt)
	}

	isCore := answers.Type == "core"
	sacJSON.Set("dessert.is_core", isCore)

	if err := sacJSON.WriteConfig(); err != nil {
		return err
	}

	if isCore {
		Logger.Info(logIsCore)
	} else {
		Logger.Info(logIsConnector)
	}
	return nil
}

func prepareDYML(sacYML *sac.Sac) error {
	sacYML.Set(YMLTokenKey, "")
	sacYML.Set(YMLVersionKey, 1)

	if err := sacYML.WriteConfig(); err != nil {
		return err
	}
	return nil
}

func preparePJSON(sacJSON *sac.Sac) error {
	if err := addDessertToKeywords(sacJSON); err != nil {
		return errors.New(errAddDessertKeyword)
	}

	if err := setCore(sacJSON); err != nil {
		return err
	}

	return nil
}
