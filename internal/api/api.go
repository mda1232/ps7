package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"resume-website/internal/models"
)

func GetGitHubRepos(cfg *models.Config) ([]string, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", cfg.GitHubUsername)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if cfg.GitHubToken != "" {
		req.Header.Set("Authorization", "Bearer "+cfg.GitHubToken)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API error: %s", body)
	}

	var repos []struct {
		Name string `json:"name"`
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &repos); err != nil {
		return nil, err
	}

	var repoNames []string
	for _, r := range repos {
		repoNames = append(repoNames, r.Name)
	}
	return repoNames, nil
}

func GetGitLabRepos(cfg *models.Config) ([]string, error) {
	url := fmt.Sprintf("https://gitlab.com/api/v4/users/%s/projects", cfg.GitLabUsername)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if cfg.GitLabToken != "" {
		req.Header.Set("PRIVATE-TOKEN", cfg.GitLabToken)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitLab API error: %s", body)
	}

	var repos []struct {
		Name string `json:"name"`
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &repos); err != nil {
		return nil, err
	}

	var repoNames []string
	for _, r := range repos {
		repoNames = append(repoNames, r.Name)
	}
	return repoNames, nil
}

func GetAllUserData(cfg *models.Config) (models.UserData, error) {
	githubRepos, err := GetGitHubRepos(cfg)
	if err != nil {
		return models.UserData{}, err
	}

	gitlabRepos, err := GetGitLabRepos(cfg)
	if err != nil {
		return models.UserData{}, err
	}

	return models.UserData{
		Email:          cfg.Email,
		GitHubUsername: cfg.GitHubUsername,
		GitHubRepos:    githubRepos,
		GitLabUsername: cfg.GitLabUsername,
		GitLabRepos:    gitlabRepos,
	}, nil
}
