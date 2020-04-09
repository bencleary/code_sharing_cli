package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/monaco-io/request"
)

type Snippet struct {
	URN         string `json:"urn"`
	Syntax      string `json:"syntax"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func main() {

	syntax := flag.String("syntax", "", "Syntax Used")
	description := flag.String("desc", "", "Description")
	content := flag.String("content", "", "Code")
	flag.Parse()

	snippet := new(Snippet)
	snippet.Syntax = *syntax
	snippet.Description = *description
	snippet.Content = *content

	snippetJSON, _ := json.Marshal(snippet)

	client := request.Client{
		URL:    "http://localhost:8000/api/snippet/",
		Method: "POST",
		Body:   snippetJSON,
	}

	resp, _ := client.Do()

	var snippetResponse Snippet
	json.Unmarshal(resp.Data, &snippetResponse)
	url := fmt.Sprintf("http://localhost:8000/api/snippet/%s/\n", string(snippetResponse.URN))
	fmt.Printf(url)

}
