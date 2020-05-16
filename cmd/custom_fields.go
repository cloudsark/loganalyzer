package cmd

import (
	"github.com/cloudsark/loganalyzer/logalizer"
	"github.com/spf13/cobra"
)

var customFields = &cobra.Command{
	Use:   "custom",
	Short: "Search for a custom field",
	Long:  "Search for a custom field. Provide the regex for the field and loganalyzer will look for it",
	Run: func(cmd *cobra.Command, args []string) {
		fieldRegex, _ := cmd.Flags().GetString("field-regex")
		logalizer.ExecuteCustomRegex(&fieldRegex, &filename)
	},
}

func init() {
	customFields.Flags().String("field-regex", "", "Custom field search")
	customFields.MarkFlagRequired("field-regex")
	rootCmd.AddCommand(customFields)
}
