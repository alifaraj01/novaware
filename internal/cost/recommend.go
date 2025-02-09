package cost

import (
	"context"
	"sort"
)

func (a *Analyzer) generateRecommendations(ctx context.Context, data *costData) ([]Recommendation, error) {
	var recommendations []Recommendation

	// Sort services by cost
	type serviceCost struct {
		name string
		cost float64
	}
	services := make([]serviceCost, 0, len(data.byService))
	for name, cost := range data.byService {
		services = append(services, serviceCost{name, cost})
	}
	sort.Slice(services, func(i, j int) bool {
		return services[i].cost > services[j].cost
	})

	// Analyze top spending services
	for _, service := range services[:min(3, len(services))] {
		switch service.name {
		case "Amazon Elastic Compute Cloud - Compute":
			rec := analyzeEC2Costs(service.cost, data.totalCost)
			if rec != nil {
				recommendations = append(recommendations, *rec)
			}
		case "Amazon Elastic Container Service":
			rec := analyzeECSCosts(service.cost, data.totalCost)
			if rec != nil {
				recommendations = append(recommendations, *rec)
			}
		case "AmazonCloudWatch":
			rec := analyzeCloudWatchCosts(service.cost, data.totalCost)
			if rec != nil {
				recommendations = append(recommendations, *rec)
			}
		}
	}

	return recommendations, nil
}

func analyzeEC2Costs(cost, totalCost float64) *Recommendation {
	percentage := (cost / totalCost) * 100
	if percentage > 30 {
		return &Recommendation{
			Service:     "EC2",
			Action:      "Consider Reserved Instances",
			Impact:      cost * 0.4, // Potential 40% savings
			Confidence:  0.85,
			Description: "High EC2 costs detected. Consider purchasing Reserved Instances for consistent workloads to save up to 40%.",
		}
	}
	return nil
}

func analyzeECSCosts(cost, totalCost float64) *Recommendation {
	percentage := (cost / totalCost) * 100
	if percentage > 20 {
		return &Recommendation{
			Service:     "ECS",
			Action:      "Optimize Container Resources",
			Impact:      cost * 0.25, // Potential 25% savings
			Confidence:  0.75,
			Description: "High ECS costs detected. Consider optimizing container resource allocations and using Spot instances for non-critical workloads.",
		}
	}
	return nil
}

func analyzeCloudWatchCosts(cost, totalCost float64) *Recommendation {
	percentage := (cost / totalCost) * 100
	if percentage > 10 {
		return &Recommendation{
			Service:     "CloudWatch",
			Action:      "Review Log Retention",
			Impact:      cost * 0.3, // Potential 30% savings
			Confidence:  0.9,
			Description: "High CloudWatch costs detected. Consider adjusting log retention periods and metric collection frequency.",
		}
	}
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
