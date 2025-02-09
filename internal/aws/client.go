package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

// Client provides access to AWS services
type Client struct {
	CostExplorer *costexplorer.Client
	EC2          *ec2.Client
	EKS          *eks.Client
}

// NewClient creates a new AWS client with default configuration
func NewClient(ctx context.Context) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %w", err)
	}

	return &Client{
		CostExplorer: costexplorer.NewFromConfig(cfg),
		EC2:          ec2.NewFromConfig(cfg),
		EKS:          eks.NewFromConfig(cfg),
	}, nil
}
