// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfig struct {
	// aggregation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/appflow_flow#aggregation_config AppflowFlow#aggregation_config}
	AggregationConfig *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfig `field:"optional" json:"aggregationConfig" yaml:"aggregationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/appflow_flow#file_type AppflowFlow#file_type}.
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
	// prefix_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/appflow_flow#prefix_config AppflowFlow#prefix_config}
	PrefixConfig *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfig `field:"optional" json:"prefixConfig" yaml:"prefixConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/appflow_flow#preserve_source_data_typing AppflowFlow#preserve_source_data_typing}.
	PreserveSourceDataTyping interface{} `field:"optional" json:"preserveSourceDataTyping" yaml:"preserveSourceDataTyping"`
}

