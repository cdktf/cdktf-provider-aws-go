// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationSemanticChunkingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/bedrockagent_data_source#breakpoint_percentile_threshold BedrockagentDataSource#breakpoint_percentile_threshold}.
	BreakpointPercentileThreshold *float64 `field:"required" json:"breakpointPercentileThreshold" yaml:"breakpointPercentileThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/bedrockagent_data_source#buffer_size BedrockagentDataSource#buffer_size}.
	BufferSize *float64 `field:"required" json:"bufferSize" yaml:"bufferSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/bedrockagent_data_source#max_token BedrockagentDataSource#max_token}.
	MaxToken *float64 `field:"required" json:"maxToken" yaml:"maxToken"`
}

