package cmd

import (
	"fmt"

	"github.com/dimeko/assets-proxy/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run:   version,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func version(command *cobra.Command, args []string) {
	fmt.Println("Version: " + config.Version)
	fmt.Println("Build Date: " + config.BuildDate)
}
