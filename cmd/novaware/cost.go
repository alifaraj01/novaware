package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/your-username/novaware/internal/aws"
	"github.com/your-username/novaware/internal/cost"
)

func newCostCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cost",
		Short: "Cost optimization and analysis",
		Long:  `Analyze AWS costs and get AI-powered optimization recommendations.`,
	}

	cmd.AddCommand(newCostAnalyzeCmd())
	return cmd
}

func newCostAnalyzeCmd() *cobra.Command {
	var days int

	cmd := &cobra.Command{
		Use:   "analyze",
		Short: "Analyze costs and get recommendations",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			// Create AWS client
			client, err := aws.NewClient(ctx)
			if err != nil {
				return fmt.Errorf("failed to create AWS client: %w", err)
			}

			// Create analyzer
			analyzer := cost.NewAnalyzer(client)

			// Run analysis
			analysis, err := analyzer.Analyze(ctx, days)
			if err != nil {
				return fmt.Errorf("analysis failed: %w", err)
			}

			// Print results
			printAnalysis(analysis)
			return nil
		},
	}

	cmd.Flags().IntVarP(&days, "days", "d", 30, "Number of days to analyze")
	return cmd
}

func printAnalysis(a *cost.Analysis) {
	fmt.Printf("\n=== Cost Analysis Report ===\n\n")
	fmt.Printf("Total Cost: $%.2f\n\n", a.TotalCost)

	fmt.Println("Top Services:")
	for service, cost := range a.ServiceCosts {
		fmt.Printf("  • %s: $%.2f\n", service, cost)
	}

	fmt.Println("\nRecommendations:")
	for _, rec := range a.Recommendations {
		fmt.Printf("\n  • %s\n", rec.Service)
		fmt.Printf("    Action: %s\n", rec.Action)
		fmt.Printf("    Impact: $%.2f\n", rec.Impact)
		fmt.Printf("    Confidence: %.0f%%\n", rec.Confidence*100)
		fmt.Printf("    %s\n", rec.Description)
	}
}
