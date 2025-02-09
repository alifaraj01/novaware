package cost

import (
	"context"
	"testing"

	"github.com/alifaraj01/novaware/internal/aws"
	"github.com/alifaraj01/novaware/internal/mock"
)

func TestAnalyzeEC2Costs(t *testing.T) {
	tests := []struct {
		name      string
		cost      float64
		totalCost float64
		wantRec   bool
	}{
		{
			name:      "high ec2 costs",
			cost:      1000,
			totalCost: 2000,
			wantRec:   true,
		},
		{
			name:      "low ec2 costs",
			cost:      100,
			totalCost: 2000,
			wantRec:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := analyzeEC2Costs(tt.cost, tt.totalCost)
			if (rec != nil) != tt.wantRec {
				t.Errorf("analyzeEC2Costs() got recommendation = %v, want %v", rec != nil, tt.wantRec)
			}
		})
	}
}

func TestAnalyzer_Analyze(t *testing.T) {
	// Create mock client
	mockClient := &aws.Client{
		CostExplorer: &mock.CostExplorerClient{},
	}

	analyzer := &Analyzer{
		client: mockClient,
	}

	ctx := context.Background()
	analysis, err := analyzer.Analyze(ctx, 30)
	if err != nil {
		t.Fatalf("Analyze() error = %v", err)
	}

	if analysis == nil {
		t.Fatal("Analyze() returned nil analysis")
	}

	// Verify the analysis results
	if analysis.TotalCost != 1500.00 { // 1000 + 500 from mock data
		t.Errorf("Expected total cost of 1500.00, got %.2f", analysis.TotalCost)
	}

	// Check EC2 costs
	if ec2Cost := analysis.ServiceCosts["Amazon Elastic Compute Cloud - Compute"]; ec2Cost != 1000.00 {
		t.Errorf("Expected EC2 cost of 1000.00, got %.2f", ec2Cost)
	}

	// Check recommendations
	if len(analysis.Recommendations) == 0 {
		t.Error("Expected at least one recommendation")
	}
}

func TestAnalyzeECSCosts(t *testing.T) {
	tests := []struct {
		name      string
		cost      float64
		totalCost float64
		wantRec   bool
	}{
		{
			name:      "high ecs costs",
			cost:      500,
			totalCost: 1000,
			wantRec:   true,
		},
		{
			name:      "low ecs costs",
			cost:      50,
			totalCost: 1000,
			wantRec:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := analyzeECSCosts(tt.cost, tt.totalCost)
			if (rec != nil) != tt.wantRec {
				t.Errorf("analyzeECSCosts() got recommendation = %v, want %v", rec != nil, tt.wantRec)
			}
		})
	}
}
