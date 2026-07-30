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

	"github.com/spf13/cobra"
)

type Response struct {
	Hits []Project `json:"hits"`
}
type Project struct {
	ProjectID   string   `json:"project_id"`
	Slug        string   `json:"slug"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	Versions    []string `json:"versions"`
	Categories  []string `json:"categories"`
}

func PrintProjectList(projectList []Project) {
	fmt.Println("Found " + fmt.Sprint(len(projectList)) + " projects")
	for i := 0; i < len(projectList); i++ {
		project := projectList[i]
		fmt.Println(fmt.Sprint(i+1) + ". " + string(project.Title))
		fmt.Println(" Author: " + string(project.Author))
		fmt.Println()
	}
}

func searchMods(query string) (Response, error) {
	v := url.Values{}
	v.Set("query", query)
	resp, err := http.Get("https://api.modrinth.com/v2/search?" + v.Encode())
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}
	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Response{}, err
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
		PrintProjectList(result.Hits)
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
