/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/IgnacioFan/todo-cli/model"
	"github.com/spf13/cobra"
)

var limit int

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todo items",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := model.ReadAll(dataFile)
		if err != nil {
			log.Printf("%v\n", err)
		}
		if len(todos) == 0 {
			log.Printf("Todo list is empty.\n")
		} else {
			log.Printf("Todo list:\n")
			if limit > len(todos) {
				limit = len(todos)
			}
			for i := 0; i < limit; i++ {
				todo := todos[i]
				log.Printf("%d. %s\n", i+1, todo.Title)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().IntVarP(&limit, "limit", "l", 3, "Display limit(default 3)")
}
