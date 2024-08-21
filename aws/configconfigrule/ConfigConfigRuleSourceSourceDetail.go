// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configconfigrule


type ConfigConfigRuleSourceSourceDetail struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/config_config_rule#event_source ConfigConfigRule#event_source}.
	EventSource *string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/config_config_rule#maximum_execution_frequency ConfigConfigRule#maximum_execution_frequency}.
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/config_config_rule#message_type ConfigConfigRule#message_type}.
	MessageType *string `field:"optional" json:"messageType" yaml:"messageType"`
}

