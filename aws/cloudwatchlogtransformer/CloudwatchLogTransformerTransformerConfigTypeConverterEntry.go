// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogtransformer


type CloudwatchLogTransformerTransformerConfigTypeConverterEntry struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#key CloudwatchLogTransformer#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#type CloudwatchLogTransformer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

