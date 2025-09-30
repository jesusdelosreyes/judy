package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var destroyEnvCmd = &cobra.Command{
	Use:   "destroy-env",
	Short: "Destroys docker-compose testing environment",
	Long:  `Runs 'docker-compose down' to stop and delete the environment containers`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := exec.LookPath("docker"); err != nil {
			return errors.New("docker is not installed or not in the system's PATH")
		}
		if _, err := os.Stat("docker-compose.yml"); os.IsNotExist(err) {
			return fmt.Errorf("error: docker-compose.yml file not found %w", err)
		}

		fmt.Println("Destroying the testing environment...")

		dockerCmd := exec.Command("docker-compose", "down")
		dockerCmd.Stdout = os.Stdout
		dockerCmd.Stderr = os.Stderr

		if err := dockerCmd.Run(); err != nil {
			return fmt.Errorf("error executing docker-compose down: %w", err)
		}

		fmt.Println("Environment succesfully destroyed")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(destroyEnvCmd)

}
