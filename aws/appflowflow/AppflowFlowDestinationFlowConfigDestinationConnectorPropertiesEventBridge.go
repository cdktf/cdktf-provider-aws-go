// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridge struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/appflow_flow#object AppflowFlow#object}.
	Object *string `field:"required" json:"object" yaml:"object"`
	// error_handling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/appflow_flow#error_handling_config AppflowFlow#error_handling_config}
	ErrorHandlingConfig *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridgeErrorHandlingConfig `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
}

