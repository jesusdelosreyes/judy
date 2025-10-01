package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

const goModTemplate = `module {{.ProjectName}}

go 1.25.1

require (
		github.com/spf13/viper v1.17.0
		github.com/stretchr/testify v1.8.4
)
`

const makefileTemplate = `.PHONY: test test-report tidy

test:
		go test -v ./tests/...

test-report:
		go test -v -json ./tests/... > test-results.json

tidy:
		go mod tidy
		go fmt ./...
`

const configYmlTemplate = `environments:
	local:
		base_url: "https://reqres.in/api"
	staging:
		base_url: "https://staging.myapi.com/v1"

default_environment: "local"
`

const configGoTemplate = `package config
import "github.com/spf13/viper"

type Config struct {
		BaseURL string ` + "`mapstructure:\"base_url\"`" + `
}

func LoadConfig() (*Config, error) {
		viper.SetConfigName("config")
		viper.AddConfigPath("./config")
		viper.SetConfigType("yml")

		if err := viper.ReadInConfig(); err != nil {
				return nil, err
		}

		env := viper.GetString("default_environment")

		var config Config
		if err := viper.UnmarshalKey("environments."+env, &config); err != nil {
				return nil, err
		}
		
		return &config, nil
}
`

const httpClientTemplate = `package http_client

import (
		"{{.ProjectName}}/config"
		"net/http"
		"time"
}

type Client struct {
		BaseURL		string
		HttpClient	*http.Client
}

func NewClient(cfg *config.Config) *Client {
		return &Client{
				BaseURL: cfg.BaseURL,
				HttpClient: &http.Client{
						Timeout: 10 * time.Second,
				},
		}
}

func (c *Client) Get(path string) (*http.Response, error) {
		fullURL := c.BaseURL + path
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
				return nil, err
		}
		// TODO: place for common headers
		// req.Header.Add("Authorization", "Bearer token"
		return c.HttpClient.Do(req)
}
`

const apiTestTemplate = `package tests

import (
		"{{.ProjectName}}/config"
		"{{.ProjectName}}/pkg/http_client"
		"io"
		"net/http"
		"testing"

		"github.com/stretchr/testify/assert"
		"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
		cfg, err := config.LoadConfig()
		require.NoError(t, err, "Failed to load config")

		client := http_client.NewClient(cfg)

		resp, err := client.Get("/users?page=2")
		require.NoError(t, err, "HTTP request failed")
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.NotEmpty(t, body, "Response body should not be empty")

		t.Logf("Successfully received user data from %s", cfg.BaseURL)
}
`

var projectType string

var setupTestsCmd = &cobra.Command{
	Use:   "setup-tests [projectName]",
	Short: "Creates folder structure for a new Go API test project",
	Long:  `Creates a complete, ready-to-run project structure for Go API testing, including configuration management, an HTTP client, and an example test.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		fmt.Printf("Creating a new test project: %s\n", projectName)

		dirs := []string{
			projectName,
			filepath.Join(projectName, "config"),
			filepath.Join(projectName, "pkg", "http_client"),
			filepath.Join(projectName, "tests"),
		}

		for _, dir := range dirs {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", dir, err)
			}
			fmt.Printf("   ✓ Created directory: %s\n", dir)
		}

		templateData := struct {
			ProjectName string
		}{
			ProjectName: projectName,
		}

		filesToCreate := map[string]struct {
			templateContent string
			data            interface{}
		}{
			"go.mod":                        {goModTemplate, templateData},
			"Makefile":                      {makefileTemplate, nil},
			"config/config.yml":             {configYmlTemplate, nil},
			"config/config.go":              {configGoTemplate, templateData},
			"pkg/http_client/client.go":     {httpClientTemplate, templateData},
			"tests/api_get_example_test.go": {apiTestTemplate, templateData},
		}

		for path, fileInfo := range filesToCreate {
			fullPath := filepath.Join(projectName, path)
			if err := createFileFromTemplate(fullPath, fileInfo.templateContent, fileInfo.data); err != nil {
				return err
			}
			fmt.Printf("   ✓ Created file:       %s\n", fullPath)
		}

		fmt.Println("\n Project scaffolding complete!")
		fmt.Printf("To get started, run:\n\n cd %s\n make tidy\n make test\n\n", projectName)

		return nil
	},
}

func createFileFromTemplate(path string, tmplContent string, data interface{}) error {
	tmpl, err := template.New(path).Parse(tmplContent)
	if err != nil {
		return fmt.Errorf("failed to parse template for %s: %w", path, err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", path, err)
	}
	defer file.Close()
	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template for %s: %w", path, err)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(setupTestsCmd)

	setupTestsCmd.Flags().StringVarP(&projectType, "type", "t", "", "Project type")

	if err := setupTestsCmd.MarkFlagRequired("type"); err != nil {
		panic(err)
	}
}
