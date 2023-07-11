package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ultra-supara/glink/utils"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "a cli tool which generates cat",
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintTxtFromTxt("cat.txt")
	},
}

func init() {
    rootCmd.AddCommand(catCmd)
}
