package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var deployEnvCmd = &cobra.Command{
	Use:   "deploy-env",
	Short: "Deploys a testing environment using docker-compose",
	Long:  `Looks for a docker-compose.yml file in the current directory and executes docker compose up -d to create and deploy the containers`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := exec.LookPath("docker"); err != nil {
			return errors.New("docker is not installed or not in the system's PATH")
		}
		if _, err := os.Stat("docker-compose.yml"); os.IsNotExist(err) {
			return fmt.Errorf("error: docker-compose.yml file not found %w", err)
		}

		fmt.Println("Deploying testing environment with docker-compose...")

		dockerCmd := exec.Command("docker", "compose", "up", "-d")
		dockerCmd.Stdout = os.Stdout
		dockerCmd.Stderr = os.Stderr

		if err := dockerCmd.Run(); err != nil {
			return fmt.Errorf("error executing docker-compose up -d: %w", err)
		}

		fmt.Println("Environment successfully deployed!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deployEnvCmd)

}
