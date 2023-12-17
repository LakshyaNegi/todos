package show

import (
	"strconv"

	"github.com/LakshyaNegi/todos/entity"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func NewModelFromTodos(todos []*entity.Todo) model {
	t := table.New(
		table.WithFocused(true),
		table.WithHeight(7),
	)

	t.SetColumns([]table.Column{
		{Title: "ID", Width: 3},
		{Title: "Task", Width: 25},
		{Title: "Status", Width: 8},
		{Title: "Due Date", Width: 10},
		{Title: "Completed At", Width: 20},
	})

	rows := []table.Row{}
	for _, todo := range todos {
		dd := ""
		if todo.DueDate.Valid {
			dd = todo.DueDate.Time.Format("2006-01-02")
		}

		ca := ""
		if todo.CompletedAt.Valid {
			ca = todo.CompletedAt.Time.String()
		}

		rows = append(rows, table.Row{
			strconv.Itoa(todo.ID),
			todo.Task,
			todo.Status.String(),
			dd,
			ca,
		})
	}

	t.SetRows(rows)

	// Table styling
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	// menu
	items := []list.Item{
		item("Press 'c' to complete a todo"),
		item("Press 'i' to mark todo pending"),
		item("Press 'x' to delete a todo"),
	}

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Options"
	l.Styles.Title = titleStyle

	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	return model{table: t, menu: l}
}
