/*
Copyright © 2025 Jesus De Los Reyes Larraga jesusdelosreyes@outlook.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const AppVersion = "v0.1.0-alpha"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the current version of Judy",
	Long:  `Judy is intended to be a toolkit for Quality Assurance Engineers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Judy %s\n", AppVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
