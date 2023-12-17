package show

import (
	"log"
	"strconv"

	"github.com/LakshyaNegi/todos/entity"
	"github.com/LakshyaNegi/todos/repo"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
	menu  list.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "c":
			id, err := strconv.Atoi(m.table.SelectedRow()[0])
			if err != nil {
				log.Printf("failed to update todo: %v", err)
				return m, tea.Batch(
					tea.Quit,
				)
			}

			todo, err := repo.GetRepo().GetTodoById(id)
			if err != nil {
				log.Printf("failed to get todo by id %v: %v", id, err)
				return m, tea.Batch(
					tea.Quit,
				)
			}

			if todo.Status == entity.TodoStatusDone {
				return m, tea.Batch(
					tea.Printf("Todo id %v is already completed", id),
				)
			}

			err = repo.GetRepo().UpdateTodoCompletedByID(id)
			if err != nil {
				log.Printf("failed to update todo: %v", err)
				return m, tea.Batch(
					tea.Quit,
				)
			}

			rows := m.table.Rows()
			cursor := m.table.Cursor()

			if cursor == 0 {
				rows = rows[1:]
			} else if cursor == len(rows)-1 {
				rows = rows[:cursor]
			} else {
				rows = append(rows[:cursor], rows[cursor+1:]...)
			}

			m.table.SetRows(rows)
			cursor = 0

			return m, tea.Batch(
				tea.Printf("Completed todo id %v!", id),
			)
		case "i":
			id, err := strconv.Atoi(m.table.SelectedRow()[0])
			if err != nil {
				log.Printf("failed to update todo: %v", err)
				return m, tea.Batch(
					tea.Quit,
				)
			}

			todo, err := repo.GetRepo().GetTodoById(id)
			if err != nil {
				log.Printf("failed to get todo by id %v: %v", id, err)
				return m, tea.Batch(
					tea.Quit,
				)
			}

			if todo.Status == entity.TodoStatusPending {
				return m, tea.Batch(
					tea.Printf("Todo id %v is already pending", id),
				)
			}

			err = repo.GetRepo().UpdateTodoIncompleteByID(id)
			if err != nil {
				log.Printf("failed to update todo: %v", err)
				return m, tea.Batch(
					tea.Quit,
				)
			}

			rows := m.table.Rows()
			cursor := m.table.Cursor()

			if cursor == 0 {
				rows = rows[1:]
			} else if cursor == len(rows)-1 {
				rows = rows[:cursor]
			} else {
				rows = append(rows[:cursor], rows[cursor+1:]...)
			}

			m.table.SetRows(rows)
			cursor = 0

			return m, tea.Batch(
				tea.Printf("Todo id %v was updated to pending", id),
			)
		case "x":
			id, err := strconv.Atoi(m.table.SelectedRow()[0])
			if err != nil {
				log.Printf("failed to update todo: %v", err)
				return m, tea.Batch(
					tea.Quit,
				)
			}

			err = repo.GetRepo().DeleteByID(id)
			if err != nil {
				log.Printf("failed to delete todo: %v", err)
				return m, tea.Batch(
					tea.Quit,
				)
			}

			rows := m.table.Rows()
			cursor := m.table.Cursor()

			if cursor == 0 {
				rows = rows[1:]
			} else if cursor == len(rows)-1 {
				rows = rows[:cursor]
			} else {
				rows = append(rows[:cursor], rows[cursor+1:]...)
			}

			m.table.SetRows(rows)
			cursor = 0

			return m, tea.Batch(
				tea.Printf("Todo id %v was deleted", id),
			)
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if len(m.table.Rows()) == 0 {
		return "No todos to show!\n"
	}

	return baseStyle.Render(m.table.View()) + "\n" + m.menu.View() + "\n"
}
