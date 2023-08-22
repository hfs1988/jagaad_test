/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hfs1988/jagaad_test/singleton"
	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Command to save some data",
	Long:  `Command to save some data for a module like user. Saving data to CSV file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			handlers singleton.Handlers = singleton.GetHandlers()
		)

		if len(args) > 0 {
			if args[0] == "user" {
				handlers.GetUserHandler().SaveUsers()
			}
			return
		}
		fmt.Println("invalid command!")
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
