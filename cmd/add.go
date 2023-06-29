package cmd

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var validate = func(input string) error {
	if govalidator.IsDNSName(input) {
		return nil
	}

	return errors.New("Invalid domain name")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Prompt{
			Label: "Host",
			Validate: validate,
		}
		prompt.Run()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
