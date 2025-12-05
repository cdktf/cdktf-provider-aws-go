// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogtransformer


type CloudwatchLogTransformerTransformerConfigParseToOcsf struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#event_source CloudwatchLogTransformer#event_source}.
	EventSource *string `field:"required" json:"eventSource" yaml:"eventSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#ocsf_version CloudwatchLogTransformer#ocsf_version}.
	OcsfVersion *string `field:"required" json:"ocsfVersion" yaml:"ocsfVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#source CloudwatchLogTransformer#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

