package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var destroyEnvCmd = &cobra.Command{
	Use:   "destroy-env",
	Short: "Destroys docker-compose testing environment",
	Long:  `Runs 'docker-compose down' to stop and delete the environment containers`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("docker-compose.yml"); os.IsNotExist(err) {
			log.Fatalf("Error: docker-compose.yml file not found %v", err)
		}

		fmt.Println("Destroying the testing environment...")

		dockerCmd := exec.Command("docker-compose", "down")
		dockerCmd.Stdout = os.Stdout
		dockerCmd.Stderr = os.Stderr

		if err := dockerCmd.Run(); err != nil {
			log.Fatalf("Error executing docker-compose down: %v", err)
		}

		fmt.Println("Environment succesfully destroyed")
	},
}

func init() {
	rootCmd.AddCommand(destroyEnvCmd)

}
