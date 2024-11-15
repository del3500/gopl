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
	Title  string `json:"title"`
	Body   string `json:"body"`
	Number int    `json:"number"`
	State  string `json:"state"`
}

type ListIssue struct {
	Title  string `json:"title"`
	NODE   string `json:"node_id"`
	Number int    `json:"number"`
}

var token = os.Getenv("GITHUB_TOKEN")

func main() {
	UpdateIssue(token, 3)
}

func UpdateIssue(token string, num int) {

	url := fmt.Sprintf("https://api.github.com/repos/del3500/go-http/issues/%d", num)
	issue := &Issue{
		Title: "sample update",
		State: "closed",
	}

	jsonData, err := json.Marshal(issue)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("error %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("error %v", err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Updated failed.\nStatus: %d\nResponse: %s", resp.StatusCode, string(body))
	}

}

func listIssues(token string) {
	var issue []ListIssue
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/repos/del3500/go-http/issues", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		json.NewDecoder(resp.Body).Decode(&issue)
		fmt.Println(issue[0].Number)
	}
}

func authenticate(token string) {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		log.Fatalf("err: %v", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var user GitHubUser
		json.NewDecoder(resp.Body).Decode(&user)
		fmt.Printf("Authentication successful. Login: %s (ID: %d)", user.Login, user.ID)
	}
}

func createIssue(token string) {
	body := &Issue{
		Title: "Sample",
		Body:  "Sample",
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
