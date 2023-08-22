/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hfs1988/jagaad_test/singleton"
	"github.com/spf13/cobra"
)

var tags []string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Command to get some data",
	Long:  `Command to get some data for each module or domain like users.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			handlers singleton.Handlers = singleton.GetHandlers()
		)

		if len(args) > 0 {
			if args[0] == "user" {
				handlers.GetUserHandler().GetUsers(tags)
			}
			return
		}
		fmt.Println("invalid command!")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringSliceVar(&tags, "tag", []string{}, "User tags")
}
