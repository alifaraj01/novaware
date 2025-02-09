package cost

import (
	"context"
	"fmt"
	"time"

	"github.com/your-username/novaware/internal/aws"
)

// Analyzer handles cost analysis and optimization recommendations
type Analyzer struct {
	client *aws.Client
}

// NewAnalyzer creates a new cost analyzer
func NewAnalyzer(client *aws.Client) *Analyzer {
	return &Analyzer{client: client}
}

// Analysis represents a cost analysis result
type Analysis struct {
	TotalCost       float64
	ServiceCosts    map[string]float64
	RegionCosts     map[string]float64
	Recommendations []Recommendation
}

// Recommendation represents a cost optimization suggestion
type Recommendation struct {
	Service     string
	Action      string
	Impact      float64
	Confidence  float64
	Description string
}

// Analyze performs cost analysis and generates recommendations
func (a *Analyzer) Analyze(ctx context.Context, days int) (*Analysis, error) {
	end := time.Now()
	start := end.AddDate(0, 0, -days)

	// Get cost and usage data
	costs, err := a.fetchCostData(ctx, start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cost data: %w", err)
	}

	// Generate recommendations
	recommendations, err := a.generateRecommendations(ctx, costs)
	if err != nil {
		return nil, fmt.Errorf("failed to generate recommendations: %w", err)
	}

	return &Analysis{
		TotalCost:       costs.totalCost,
		ServiceCosts:    costs.byService,
		RegionCosts:     costs.byRegion,
		Recommendations: recommendations,
	}, nil
}
