// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elb


type ElbHealthCheck struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/elb#healthy_threshold Elb#healthy_threshold}.
	HealthyThreshold *float64 `field:"required" json:"healthyThreshold" yaml:"healthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/elb#interval Elb#interval}.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/elb#target Elb#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/elb#timeout Elb#timeout}.
	Timeout *float64 `field:"required" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/elb#unhealthy_threshold Elb#unhealthy_threshold}.
	UnhealthyThreshold *float64 `field:"required" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

