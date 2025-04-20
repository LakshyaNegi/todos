package del

import (
	"fmt"
	"log"

	"github.com/LakshyaNegi/todos/internal/entity"
	"github.com/LakshyaNegi/todos/internal/repo"
	"github.com/LakshyaNegi/todos/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	selectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(ui.SelectionBgColor))
	crossStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color(ui.RedColor))
)

type model struct {
	cursor   int
	choices  []*entity.Todo
	selected map[int]*entity.Todo
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Todos")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = m.choices[m.cursor]
			}
		case "enter":
			if len(m.selected) == 0 {
				return m, tea.Quit
			}

			for _, todo := range m.selected {
				err := repo.GetRepo().DeleteByID(todo.ID)
				if err != nil {
					log.Printf("failed to delete todo: %v", err)
					return m, tea.Batch(
						tea.Quit,
					)
				}

				log.Printf("todo id %v deleted", todo.ID)
			}

			return m, tea.Batch(
				tea.Quit,
			)
		}
	}

	return m, nil
}

func (m model) View() string {
	if len(m.choices) == 0 {
		return "No todos to show!\n"
	}

	s := "\nMark todos to delete\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = crossStyle.Render("x")
		}

		if m.cursor == i {
			s += selectedItemStyle.Render(fmt.Sprintf("\n%s [", cursor))
			s += checked
			s += selectedItemStyle.Render(fmt.Sprintf("] %s", choice.Task))
		} else {
			s += fmt.Sprintf("\n%s [%s] %s", cursor, checked, choice.Task)
		}

	}

	s += "\n\nPress q to quit.\nPress space to select a todo.\nPress enter to delete selected todos.\n\n"

	return s
}
