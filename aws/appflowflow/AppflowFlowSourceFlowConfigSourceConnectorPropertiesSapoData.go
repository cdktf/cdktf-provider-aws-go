// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/appflow_flow#object_path AppflowFlow#object_path}.
	ObjectPath *string `field:"required" json:"objectPath" yaml:"objectPath"`
	// pagination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/appflow_flow#pagination_config AppflowFlow#pagination_config}
	PaginationConfig *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataPaginationConfig `field:"optional" json:"paginationConfig" yaml:"paginationConfig"`
	// parallelism_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/appflow_flow#parallelism_config AppflowFlow#parallelism_config}
	ParallelismConfig *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataParallelismConfig `field:"optional" json:"parallelismConfig" yaml:"parallelismConfig"`
}

