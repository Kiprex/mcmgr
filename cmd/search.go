/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	//"os"
	"strings"
	mr "github.com/kiprex/mcmgr/internal/modrinth"
	"github.com/spf13/cobra"
)



func PrintSearchHitList(SearchHitList []mr.SearchHit) {
	fmt.Println("Found " + fmt.Sprint(len(SearchHitList)) + " SearchHits")
	for i := 0; i < len(SearchHitList); i++ {
		SearchHit := SearchHitList[i]
		fmt.Println(fmt.Sprint(i+1) + ". " + string(SearchHit.Title))
		fmt.Println(" Author: " + string(SearchHit.Author))
		fmt.Println()
	}
}

func searchMods(query string) (mr.SearchResponse, error) {
	v := url.Values{}
	v.Set("query", query)
	resp, err := http.Get("https://api.modrinth.com/v2/search?" + v.Encode())
	if err != nil {
		return mr.SearchResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return mr.SearchResponse{}, err
	}
	var result mr.SearchResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return mr.SearchResponse{}, err
	}
	return result, nil
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
		result, err := searchMods(strings.Join(args, " "))
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		if len(result.Hits) == 0 {
			fmt.Println("Список пуст, показывать нечего")
			return
		}
		PrintSearchHitList(result.Hits)
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
