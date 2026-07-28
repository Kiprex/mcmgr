package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Response struct {
	Hits []Project `json:"hits"`
}

type Project struct {
	ProjectID string `json:"project_id"`
	Slug string `json:"slug"`
	Title string `json:"title"`
	Description string `json:"description"`
	Author string `json:"author"`
	Versions []string `json:"versions"`
	Categories []string `json:"categories"`
}


func printSearchResult(searchResult []Project){
	fmt.Println("Found " + fmt.Sprint(len(searchResult)) + " projects")
	for i := 0; i < len(searchResult); i++{
		project := searchResult[i]
		fmt.Println(fmt.Sprint(i + 1) + ". " + project.Title)
		fmt.Println("   Author: " + string(project.Author))
		fmt.Println()
	}
}


func search(query string) (Response, error) {
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


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: not enough args. Usage: go run . {query}")
		return
	}

	
	query := strings.Join(os.Args[1:], " ")
	
	fmt.Println(query)


	result, err := search(query)
	if err != nil{
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(len(result.Hits))
	if len(result.Hits) == 0 {
		fmt.Println("Список пуст, показывать нечего")
		return
	}
	printSearchResult(result.Hits)
}
