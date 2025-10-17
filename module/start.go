/*Paclage module
* @fevunge
* */
package module

import tea "github.com/charmbracelet/bubbletea"

type Start struct {
	in   int
	tabs []string
}

func InitialModel() Start {
	return Start{
		in:   0,
		tabs: []string{"Enter a Room", "Find For Room", "Generate Room", "About"},
	}
}

func (m Start) Init() tea.Cmd {
	return tea.SetWindowTitle("OK")
}

func (m Start) Update() (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Start) View() string {
	return ""
}
