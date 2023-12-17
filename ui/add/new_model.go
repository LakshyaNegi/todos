package add

import "github.com/charmbracelet/bubbles/textinput"

func NewAddModel() model {
	m := model{
		inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Task"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Due Date (YYYY-MM-DD)"
			t.CharLimit = 64
		}

		m.inputs[i] = t
	}

	return m
}
