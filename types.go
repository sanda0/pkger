package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type config struct {
	PkgRoot   string `json:"package_root"`
	Prefix    string `json:"prefix"`
	Namespace string `json:"namespace"`
}

type NewPkgFlags struct {
	AuthorName  string
	AuthorEmail string
	PkgName     string
}

type folderSelectForm struct {
	Choices     []string
	Cursor      int
	Selected    map[int]struct{}
	Config      *config
	PkgName     string
	Quit        bool
	AuthorName  string
	AuthorEmail string
}

// Note: folderSelectForm{} is tea.Model thats why i use "m" in these methods
func (m folderSelectForm) Init() tea.Cmd {
	return nil
}

func (m folderSelectForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.Quit = true
			return m, tea.Quit
		case "up", "l":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "k":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case " ":
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = struct{}{}
			}
		case "enter":
			m.Quit = false
			return m, tea.Quit
		}

	}
	return m, nil
}

func (m folderSelectForm) View() string {
	s := fmt.Sprintf(`
New Package : %s/%s%s

Select folders you want in package
(use space key to select)


`, m.Config.PkgRoot, m.Config.Prefix, m.PkgName)

	for i, choice := range m.Choices {
		cursor := "  "
		if m.Cursor == i {
			cursor = "👉️"
		}
		checked := " "
		if _, ok := m.Selected[i]; ok {
			checked = "✅"
		}
		s += fmt.Sprintf("%s %s %s\n", cursor, checked, choice)
	}

	s += "\n\npress q to quit or enter to continue"
	return s
}
