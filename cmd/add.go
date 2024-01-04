/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Olowodev/tri/todo"
	"github.com/spf13/cobra"
)

var priority int

func addRun(cmd *cobra.Command, args []string) {
	// fmt.Println("add called")
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		fmt.Printf("%v", err)
	}
	for _, x := range args {
		// fmt.Println(x)
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}
	err = todo.SaveItems(dataFile, items)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	// fmt.Printf("%#v\n", items)
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("add called")
	// },
	Run: addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
