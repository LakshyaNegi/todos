/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/LakshyaNegi/todos/internal/entity"
	"github.com/LakshyaNegi/todos/internal/repo"
	"github.com/LakshyaNegi/todos/internal/ui/show"
	"github.com/LakshyaNegi/todos/pkg/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all todos",
	Long:  `Show all todos. If no flags are provided, all incomplete todos will be shown. If a todo id is provided, only that todo will be shown.`,
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() > 0 && len(args) > 0 {
			log.Fatal("please provide either flags or arguments, not both")

		}

		todos := []*entity.Todo{}
		var err error

		if cmd.Flags().NFlag() > 0 {
			cmd.Flags().VisitAll(func(flag *pflag.Flag) {
				if !(flag.Changed) {
					return
				}

				switch flag.Name {
				case "all":
					todos, err = repo.GetRepo().GetTodos()
					if err != nil {
						log.Fatalf("failed to get all todos: %v", err)
					}
				case "complete":
					days, err := cmd.Flags().GetInt("complete")
					if err != nil {
						log.Fatalf("failed to parse complete flag: %v", err)
					}

					if days < 0 {
						todos, err = repo.GetRepo().GetCompletedTodos()
						if err != nil {
							log.Fatalf("failed to get completed todos: %v", err)
						}
					} else {
						now := utils.Bod(time.Now()).AddDate(0, 0, -1*days)

						todos, err = repo.GetRepo().GetCompletedTodosAfter(now)
						if err != nil {
							log.Fatalf("failed to get completed todos: %v", err)
						}
					}
				case "pending":
					todos, err = repo.GetRepo().GetPendingTodos()
					if err != nil {
						log.Fatalf("failed to get pending todos: %v", err)
					}
				case "due":
					days, err := cmd.Flags().GetInt("due")
					if err != nil {
						log.Fatalf("failed to parse due flag: %v", err)
					}

					if days < 0 {
						todos, err = repo.GetRepo().GetDueTodos()
						if err != nil {
							log.Fatalf("failed to get due todos: %v", err)
						}
					} else {
						now := utils.Bod(time.Now()).AddDate(0, 0, days)

						todos, err = repo.GetRepo().GetDueTodosBefore(now)
						if err != nil {
							log.Fatalf("failed to get due todos: %v", err)
						}
					}
				case "overdue":
					todos, err = repo.GetRepo().GetOverdueTodos()
					if err != nil {
						log.Fatalf("failed to get overdue todos: %v", err)
					}
				case "today":
					now := utils.Bod(time.Now()).AddDate(0, 0, 1)

					todos, err = repo.GetRepo().GetDueTodosBefore(now)
					if err != nil {
						log.Fatalf("failed to get completed todos: %v", err)
					}
				default:
					log.Fatalf("invalid flag: %s", flag.Name)
				}
			})
		} else if len(args) == 1 {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal("please provide a valid integer todo id")
			}

			todo, err := repo.GetRepo().GetTodoById(id)
			if err != nil {
				log.Printf("failed to get todos: %v", err)
			}

			todos = append(todos, todo)
		} else {
			todos, err = repo.GetRepo().GetIncompleteTodosOrderedByDueDateAsc()
			if err != nil {
				log.Printf("failed to get todos: %v", err)
			}
		}

		num := len(todos)
		if cmd.Flag("num").Changed {
			num, err = cmd.Flags().GetInt("num")
			if err != nil {
				log.Fatalf("failed to parse num flag: %v", err)
			}
		}

		m := show.NewModelFromTodos(todos[:num])

		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.PersistentFlags().IntP("num", "n", -1, "number of todos to show")

	showCmd.PersistentFlags().BoolP("all", "a", false, "show all todos")
	showCmd.PersistentFlags().BoolP("overdue", "o", false, "show all overdue todos")
	showCmd.PersistentFlags().BoolP("today", "t", false, "show all todos due today")

	showCmd.PersistentFlags().IntP("complete", "c", -1, "show todos completed in last n days")
	showCmd.Flag("complete").NoOptDefVal = "-1"

	showCmd.PersistentFlags().BoolP("pending", "p", false, "show todos pending in next n days")
	showCmd.PersistentFlags().IntP("due", "d", -1, "show all todos with due dates")
	showCmd.Flag("due").NoOptDefVal = "-1"

	showCmd.MarkFlagsMutuallyExclusive(
		"all",
		"complete",
		"pending",
		"due",
		"overdue",
		"today",
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
