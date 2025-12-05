// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogtransformer


type CloudwatchLogTransformerTransformerConfigDateTimeConverter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#match_patterns CloudwatchLogTransformer#match_patterns}.
	MatchPatterns *[]*string `field:"required" json:"matchPatterns" yaml:"matchPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#source CloudwatchLogTransformer#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#target CloudwatchLogTransformer#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#locale CloudwatchLogTransformer#locale}.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#source_timezone CloudwatchLogTransformer#source_timezone}.
	SourceTimezone *string `field:"optional" json:"sourceTimezone" yaml:"sourceTimezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#target_format CloudwatchLogTransformer#target_format}.
	TargetFormat *string `field:"optional" json:"targetFormat" yaml:"targetFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#target_timezone CloudwatchLogTransformer#target_timezone}.
	TargetTimezone *string `field:"optional" json:"targetTimezone" yaml:"targetTimezone"`
}

