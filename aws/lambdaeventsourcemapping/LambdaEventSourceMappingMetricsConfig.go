// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingMetricsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/lambda_event_source_mapping#metrics LambdaEventSourceMapping#metrics}.
	Metrics *[]*string `field:"required" json:"metrics" yaml:"metrics"`
}

