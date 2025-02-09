package cost

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/your-username/novaware/internal/aws"
)

type costData struct {
	totalCost float64
	byService map[string]float64
	byRegion  map[string]float64
}

func (a *Analyzer) fetchCostData(ctx context.Context, start, end time.Time) (*costData, error) {
	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &types.DateInterval{
			Start: aws.String(start.Format("2006-01-02")),
			End:   aws.String(end.Format("2006-01-02")),
		},
		Granularity: types.GranularityMonthly,
		GroupBy: []types.GroupDefinition{
			{
				Type: types.GroupDefinitionTypeDimension,
				Key:  aws.String("SERVICE"),
			},
			{
				Type: types.GroupDefinitionTypeDimension,
				Key:  aws.String("REGION"),
			},
		},
		Metrics: []string{"UnblendedCost"},
	}

	result, err := a.client.CostExplorer.GetCostAndUsage(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get cost data: %w", err)
	}

	data := &costData{
		byService: make(map[string]float64),
		byRegion:  make(map[string]float64),
	}

	// Parse results
	for _, resultByTime := range result.ResultsByTime {
		for _, group := range resultByTime.Groups {
			cost, _ := group.Metrics["UnblendedCost"].Amount.Float64()
			data.totalCost += cost

			// Parse service and region from group keys
			if len(group.Keys) >= 2 {
				service := group.Keys[0]
				region := group.Keys[1]
				data.byService[service] += cost
				data.byRegion[region] += cost
			}
		}
	}

	return data, nil
}
