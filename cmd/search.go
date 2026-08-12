/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	mr "github.com/kiprex/mcmgr/internal/modrinth"
	"github.com/spf13/cobra"
	"strings"
)

func printSearchHitList(searchHitList []mr.SearchHit) {
	fmt.Println("Found " + fmt.Sprint(len(searchHitList)) + " SearchHits")
	for i, searchHit := range searchHitList {
		fmt.Println(fmt.Sprint(i+1) + ". " + searchHit.Title)
		fmt.Println(" Author: " + searchHit.Author)
		fmt.Println()
	}
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		searchArgs := strings.Join(args, " ")
		result, err := mr.Search(searchArgs)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		if len(result.Hits) == 0 {
			fmt.Println("Список пуст, показывать нечего")
			return
		}
		printSearchHitList(result.Hits)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
