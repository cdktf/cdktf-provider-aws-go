// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingProvisionedPollerConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/lambda_event_source_mapping#maximum_pollers LambdaEventSourceMapping#maximum_pollers}.
	MaximumPollers *float64 `field:"optional" json:"maximumPollers" yaml:"maximumPollers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/lambda_event_source_mapping#minimum_pollers LambdaEventSourceMapping#minimum_pollers}.
	MinimumPollers *float64 `field:"optional" json:"minimumPollers" yaml:"minimumPollers"`
}

