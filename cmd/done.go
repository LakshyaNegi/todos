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
	"github.com/LakshyaNegi/todos/ui/done"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a todo as done",
	Long:  `Mark a todo as done by providing the todo id. If no id is provided, a list of todos will be shown to choose from.`,
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal("please provide a valid integer todo id")
			}

			err = repo.GetRepo().UpdateTodoCompletedByID(id)
			if err != nil {
				log.Printf("failed to update todo: %v", err)
			}

			return
		}

		todos, err := repo.GetRepo().GetPendingTodos()
		if err != nil {
			log.Fatalf("failed to get pending todos: %v", err)
		}

		if _, err := tea.NewProgram(done.NewDoneModelFromTodos(todos)).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
