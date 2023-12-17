/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/LakshyaNegi/todos/repo"
	"github.com/LakshyaNegi/todos/ui/del"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal("please provide a valid integer todo id")
			}

			err = repo.GetRepo().DeleteByID(id)
			if err != nil {
				log.Printf("failed to delete todo: %v", err)
			}

			return
		}

		todos, err := repo.GetRepo().GetTodos()
		if err != nil {
			log.Fatalf("failed to get todos: %v", err)
		}

		if _, err := tea.NewProgram(del.NewDeleteModelFromTodos(todos)).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
