package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type GitHubUser struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
}

type Issue struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Labels []string `json:"labels"`
}

var token = os.Getenv("GITHUB_TOKEN")

func main() {
	authenticate(token)
	createIssue(token)
}

func authenticate(token string) {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		fmt.Errorf("err: %v", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("err: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var user GitHubUser
		json.NewDecoder(resp.Body).Decode(&user)
		fmt.Printf("Authentication successful. Login: %s (ID: %d)", user.Login, user.ID)
	}
}

/*
func createIssue(token string) {
	body := &Issue{
		Title:  "Sample",
		Body:   "Sample",
		Labels: "Sample",
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, "/repos/del3500/go-http/issues", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
	defer resp.Body.Close()

	resbdy, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	if resp.StatusCode == 201 {
		fmt.Println("Post request successful: Response: ", string(resbdy))
	} else {
		fmt.Errorf("Received a non-200 response status: %s. Response: %s", resp.Status, string(resbdy))
	}

}*/

func createIssue(token string) {
	body := &Issue{
		Title:  "Sample",
		Body:   "Sample",
		Labels: []string{"Sample"},
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Ensure you have the full URL
	url := "https://api.github.com/repos/del3500/go-http/issues" // Use the correct full URL
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Check if response is nil
	if resp == nil {
		log.Fatalf("Received nil response")
	}

	resbdy, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Post request successful: Response: ", string(resbdy))
	} else {
		log.Printf("Received a non-200 response status: %s. Response: %s", resp.Status, string(resbdy))
	}
}
