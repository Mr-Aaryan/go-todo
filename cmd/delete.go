/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"todo/database"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type item struct {
    ID         string
    IsSelected bool
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		todos := database.ReadDB()
		prompt := promptui.Select{
			Label: "🗑️  Select a task to delete",
			Items: todos,
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}",
				Active:   "👉 \033[36;4m{{ .Title }}\033[0m",
				Inactive: "   {{ .Title }}",
				Selected: "✅ Successfully deleted: \033[31m{{ .Title }}\033[0m",
			},
			Size: 8,
		}

		i, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		result := todos[i]
		database.DeleteFromDB(result.Id)
		
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}