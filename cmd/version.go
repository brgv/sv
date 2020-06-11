package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "My short description",
	Long:  "My long description",
	Run:   versionCmdFunc,
}

var versionCmdFunc = func(cmd *cobra.Command, args []string) {
	//
}
