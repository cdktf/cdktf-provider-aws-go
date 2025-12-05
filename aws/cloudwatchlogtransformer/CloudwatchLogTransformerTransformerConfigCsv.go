// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogtransformer


type CloudwatchLogTransformerTransformerConfigCsv struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#columns CloudwatchLogTransformer#columns}.
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#delimiter CloudwatchLogTransformer#delimiter}.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#quote_character CloudwatchLogTransformer#quote_character}.
	QuoteCharacter *string `field:"optional" json:"quoteCharacter" yaml:"quoteCharacter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#source CloudwatchLogTransformer#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

