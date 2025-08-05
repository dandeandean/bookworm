package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dandeandean/bookworm/internal"
)

type model struct {
	choices  []internal.BookMark // items on the to-do list
	cursor   int                 // which to-do list item our cursor is pointing at
	selected map[int]struct{}    // which to-do items are selected
}

func TeaModel() model {
	choices := make([]internal.BookMark, 0)
	for _, bm := range Bw.Cfg.BookMarks {
		choices = append(choices, *bm)
	}
	return model{
		choices:  choices,
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	s := "Bookmarks\n\n"

	for i, choice := range m.choices {

		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice.Name)
	}

	s += "\nPress q to quit.\n"

	return s
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
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
				internal.OpenURL(m.choices[m.cursor].Link)
				return m, tea.Quit
			}
		}
	}
	return m, nil
}
