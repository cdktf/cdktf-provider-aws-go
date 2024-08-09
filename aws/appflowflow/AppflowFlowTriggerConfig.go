// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowTriggerConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/appflow_flow#trigger_type AppflowFlow#trigger_type}.
	TriggerType *string `field:"required" json:"triggerType" yaml:"triggerType"`
	// trigger_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/appflow_flow#trigger_properties AppflowFlow#trigger_properties}
	TriggerProperties *AppflowFlowTriggerConfigTriggerProperties `field:"optional" json:"triggerProperties" yaml:"triggerProperties"`
}

