/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strings"

	"github.com/IgnacioFan/todo-cli/model"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for todo items with specific strings",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := model.ReadAll(dataFile)
		if err != nil {
			log.Printf("%v\n", err)
		}
		searchString := strings.Join(args, " ")
		match := false
		log.Printf("Todo items containing '%s':\n", searchString)
		for i, todo := range todos {
			if strings.Contains(todo.Title, searchString) {
				match = true
				log.Printf("%d. %s \n", i+1, todo.Title)
			}
		}
		if !match {
			log.Printf("No todo item found!\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
