package models

type Config struct {
	GitHubToken    string `json:"github_token"`
	GitLabToken    string `json:"gitlab_token"`
	GitHubUsername string `json:"github_username"`
	GitLabUsername string `json:"gitlab_username"`
	Email          string `json:"email"` 
}

type UserData struct {
	Email          string   
	GitHubUsername string   
	GitHubRepos    []string
	GitLabUsername string   
	GitLabRepos    []string
}
