package cmd

import "github.com/spf13/cobra"

var migrateSince uint
var migrateMessage string

func init() {
	rootCmd.AddCommand(migrateCmd)

	migrateCmd.Flags().UintVarP(
		&migrateSince,
		"since",
		"s", 0,
		"since for migration",
	)

	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)

	migrateCmd.AddCommand(migrateVersionCmd)

	migrateCmd.AddCommand(migrateNewCmd)

	migrateNewCmd.Flags().StringVarP(
		&migrateMessage,
		"message",
		"m",
		"noname",
		"message for migration files",
	)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate to version",
	Long:  "Migrate to version",
	Run:   migrateCmdFunc,
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Upgrade to a later version",
	Long:  "Upgrade to a later version",
	Run:   migrateUpCmdFunc,
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Revert to a previous version",
	Long:  "Revert to a previous version",
	Run:   migrateDownCmdFunc,
}

var migrateNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new revision files",
	Long:  "Create a new revision files",
	Run:   migrateNewCmdFunc,
}

var migrateVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the current revision for a database",
	Long:  "Display the current revision for a database",
	Run:   migrateVersionCmdFunc,
}

func migrateCmdFunc(cmd *cobra.Command, args []string) {}

func migrateUpCmdFunc(cmd *cobra.Command, args []string) {}

func migrateDownCmdFunc(cmd *cobra.Command, args []string) {}

func migrateNewCmdFunc(cmd *cobra.Command, args []string) {}

func migrateVersionCmdFunc(cmd *cobra.Command, args []string) {}
