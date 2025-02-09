package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"
	commit  = "development"
)

func main() {
	rootCmd := &cobra.Command{
		Use:          "novaware",
		Short:        "NOVAWARE - AI-Driven Cloud Automation Platform",
		SilenceUsage: true,
	}

	// Add commands
	rootCmd.AddCommand(
		newVersionCmd(),
		newCostCmd(),     // Cost optimization commands
		newIncidentCmd(), // Incident management
		newObserveCmd(),  // Observability
		newDeployCmd(),   // Deployment management
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("novaware %s (commit: %s)\n", version, commit)
		},
	}
}
