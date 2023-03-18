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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add todo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := model.ReadAll(dataFile)
		if err != nil {
			log.Printf("%v\n", err)
		}
		title := strings.Join(args, " ")
		todos = append(todos, &model.Todo{
			Title: title,
		})
		if err = model.SaveAll(dataFile, todos); err != nil {
			log.Printf("%v\n", err)
		}
		log.Printf("Todo item added.\n")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
