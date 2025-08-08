// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforce struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/appflow_flow#object AppflowFlow#object}.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/appflow_flow#data_transfer_api AppflowFlow#data_transfer_api}.
	DataTransferApi *string `field:"optional" json:"dataTransferApi" yaml:"dataTransferApi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/appflow_flow#enable_dynamic_field_update AppflowFlow#enable_dynamic_field_update}.
	EnableDynamicFieldUpdate interface{} `field:"optional" json:"enableDynamicFieldUpdate" yaml:"enableDynamicFieldUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/appflow_flow#include_deleted_records AppflowFlow#include_deleted_records}.
	IncludeDeletedRecords interface{} `field:"optional" json:"includeDeletedRecords" yaml:"includeDeletedRecords"`
}

