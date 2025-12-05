// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogtransformer


type CloudwatchLogTransformerTransformerConfigSplitStringEntry struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#delimiter CloudwatchLogTransformer#delimiter}.
	Delimiter *string `field:"required" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#source CloudwatchLogTransformer#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
}

