package cmd

import (
	"fmt"
	"github.com/brgv/sv/internal/version"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   version.Executable + "ctl",
	Short: "My short description",
	Long:  "My long description",
	Run:   rootCmdFunc,
}

var rootCmdFunc = func(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
