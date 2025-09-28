package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var projectType string

var setupTestsCmd = &cobra.Command{
	Use:   "setup-tests [projectName]",
	Short: "Creates folder structure for a new testing project",
	Long:  `Creates a new folder structure with directories and standard files to start a new automation test project as fast as possible`,

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		if projectType != "api" {
			log.Fatalf("Error: project type '%s' is invalid. There's only api support for now", projectType)
		}

		fmt.Printf("Creating API Testing project '%s' ...\n", projectName)

		if err := os.Mkdir(projectName, 0755); err != nil {
			log.Fatalf("Error when creating project directory: %v", err)
		}

		dirs := []string{"cmd", "configs", "tests"}
		for _, dir := range dirs {
			dirPath := filepath.Join(projectName, dir)
			if err := os.Mkdir(dirPath, 0755); err != nil {
				log.Fatalf("Error when creating subdirectory '%s': %v", dir, err)
			}
		}

		files := []string{
			"cmd/main_test.go",
			"configs/settings.json",
			"tests/api_test.go",
			"go.mod",
		}
		for _, file := range files {
			filePath := filepath.Join(projectName, file)
			if _, err := os.Create(filePath); err != nil {
				log.Fatalf("Error creating file '%s': %v", file, err)
			}
		}

		fmt.Println("Project succesfully created")
	},
}

func init() {
	rootCmd.AddCommand(setupTestsCmd)

	setupTestsCmd.Flags().StringVarP(&projectType, "type", "t", "", "Project type")

	if err := setupTestsCmd.MarkFlagRequired("type"); err != nil {
		log.Fatalf("Error setting Flag 'type' as required: %v", err)
	}
}
