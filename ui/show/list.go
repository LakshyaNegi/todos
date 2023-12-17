package show

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	defaultWidth = 20
	listHeight   = 8
)

var (
	titleStyle = lipgloss.NewStyle().MarginLeft(2)
	itemStyle  = lipgloss.NewStyle().PaddingLeft(4)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	// if index == m.Index() {
	// 	fn = func(s ...string) string {
	// 		return selectedItemStyle.Render("> " + strings.Join(s, " "))
	// 	}
	// }

	fmt.Fprint(w, fn(str))
}
