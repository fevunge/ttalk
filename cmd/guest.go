/*
Package cmd
Copyright © 2025 Fernando Vunge <fevunge.outlook.com>
*/
package cmd

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// guestCmd represents the guest command
var guestCmd = &cobra.Command{
	Use:   "guest",
	Short: "Enter on server",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
		name := viper.GetString("name")
		box := tview.NewBox().SetBorder(true).SetTitle(name)
		nameTextArea := tview.NewTextArea().SetPlaceholder("murfer drone")
		box.SetBorderColor(tcell.Color112)
		if err := app.SetRoot(box, false).Run(); err != nil {
			panic(err)
		}
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
