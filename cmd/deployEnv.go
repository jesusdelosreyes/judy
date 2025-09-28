package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var deployEnvCmd = &cobra.Command{
	Use:   "deploy-env",
	Short: "Deploys a testing environment using docker-compose",
	Long:  `Looks for a docker-compose.yml file in the current directory and executes docker-compose up -d to create and deploy the containers`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("docker-compose.yml"); os.IsNotExist(err) {
			log.Fatalf("Error: docker-compose.yml file not found %v", err)
		}

		fmt.Println("Deploying testing environment with docker-compose...")

		dockerCmd := exec.Command("docker-compose", "up", "-d")
		dockerCmd.Stdout = os.Stdout
		dockerCmd.Stderr = os.Stderr

		if err := dockerCmd.Run(); err != nil {
			log.Fatalf("Error executing docker-compose up -d: %v", err)
		}

		fmt.Println("Environment succesfully deployed!")
	},
}

func init() {
	rootCmd.AddCommand(deployEnvCmd)

}
