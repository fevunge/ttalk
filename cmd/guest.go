/*
Package cmd
Copyright © 2025 Fernando Vunge <fevunge.outlook.com>
*/
package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// guestCmd represents the guest command
type model struct {
	cursor   int
	choices  []string
	selected map[int]struct{}
}

var textStyle = gloss.NewStyle().Bold(true).Foreground(gloss.Color("5")).PaddingLeft(2)

func initialModel() model {
	return model{
		2,
		[]string{"Execute git add ", "Commit all changes", "Push it to remore repo"},
		make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Exemple")
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
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "What should we buy at the market?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		checked := " "

		if m.cursor == i {
			cursor = ">"
		}
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		choice = textStyle.Render(choice)
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress q to quit.\n"
	return s
}

var guestCmd = &cobra.Command{
	Use:   "guest",
	Short: "Enter on server",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := tea.NewProgram(initialModel())
		app.Run()
		name := viper.GetString("name")
		fmt.Println(name)
	},
}

func init() {
	rootCmd.AddCommand(guestCmd)
	guestCmd.Flags().String("name", "fevunge", "Name to showed in network!")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// guestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// guestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
