package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GitHubUser struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
}

var token = os.Getenv("GITHUB_TOKEN")

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		fmt.Errorf("err: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("err %v:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var user GitHubUser
		json.NewDecoder(resp.Body).Decode(&user)
		fmt.Printf("Authorization Successful. Login: %s (ID: %d)", user.Login, user.ID)
	}
}
