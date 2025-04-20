/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the application",
	Long:  `Login to the application using your credentials. If you don't have an account, you can create one using the 'register' command.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
