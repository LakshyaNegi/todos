/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"log"
	"os"

	"github.com/LakshyaNegi/todos/internal/repo"
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the database",
	Long:  `Reset the database. This will delete all the todos and create a new database file.`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Remove("data/todos.db")

		file, err := os.Create("data/todos.db")
		if err != nil {
			log.Fatal("failed to create database file", err)
		}

		file.Close()

		repo.InitRepo()
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
