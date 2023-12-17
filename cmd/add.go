/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/LakshyaNegi/todos/repo"
	"github.com/LakshyaNegi/todos/ui/add"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `This command adds a new todo item to your list of todos.`,
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if cmd.Flags().NFlag() > 0 {
				cmd.Flags().VisitAll(func(flag *pflag.Flag) {
					if !(flag.Changed) {
						return
					}

					switch flag.Name {
					case "dueIn":
						dueIn, err := cmd.Flags().GetInt("dueIn")
						if err != nil {
							log.Fatalf("failed to parse dueIn flag: %v", err)
						}

						now := time.Now()
						err = repo.GetRepo().CreateTodoWithDueDate(
							args[0], now.AddDate(0, 0, dueIn),
						)
						if err != nil {
							log.Fatalf("failed to add new todo: %v", err)
						}
					case "dueDate":
						dueDate, err := cmd.Flags().GetString("dueDate")
						if err != nil {
							log.Fatalf("failed to parse dueIn flag: %v", err)
						}

						date, err := time.Parse("2006-01-02", dueDate)
						if err != nil {
							log.Fatalf("failed to parse dueDate flag: %v", err)
						}

						err = repo.GetRepo().CreateTodoWithDueDate(args[0], date)
						if err != nil {
							log.Fatalf("failed to add new todo: %v", err)
						}
					default:
						log.Fatalf("invalid flag: %s", flag.Name)
					}
				})

				return
			}

			err := repo.GetRepo().CreateTodo(args[0])
			if err != nil {
				log.Fatalf("failed to add new todo: %v", err)
			}

			return
		}

		if _, err := tea.NewProgram(add.NewAddModel()).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().IntP("dueIn", "i", -1, "due in how many days")
	addCmd.PersistentFlags().StringP("dueDate", "d", "", "due in how many days")

	addCmd.MarkFlagsMutuallyExclusive("dueIn", "dueDate")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
