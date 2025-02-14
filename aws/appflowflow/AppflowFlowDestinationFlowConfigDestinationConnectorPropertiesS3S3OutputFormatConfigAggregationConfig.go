// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/appflow_flow#aggregation_type AppflowFlow#aggregation_type}.
	AggregationType *string `field:"optional" json:"aggregationType" yaml:"aggregationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/appflow_flow#target_file_size AppflowFlow#target_file_size}.
	TargetFileSize *float64 `field:"optional" json:"targetFileSize" yaml:"targetFileSize"`
}

