// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfig struct {
	// prefix_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/appflow_flow#prefix_config AppflowFlow#prefix_config}
	PrefixConfig *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfig `field:"required" json:"prefixConfig" yaml:"prefixConfig"`
	// aggregation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/appflow_flow#aggregation_config AppflowFlow#aggregation_config}
	AggregationConfig *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfig `field:"optional" json:"aggregationConfig" yaml:"aggregationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/appflow_flow#file_type AppflowFlow#file_type}.
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
}

