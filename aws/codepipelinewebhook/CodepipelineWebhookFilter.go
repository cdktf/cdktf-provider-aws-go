// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipelinewebhook


type CodepipelineWebhookFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.31.0/docs/resources/codepipeline_webhook#json_path CodepipelineWebhook#json_path}.
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.31.0/docs/resources/codepipeline_webhook#match_equals CodepipelineWebhook#match_equals}.
	MatchEquals *string `field:"required" json:"matchEquals" yaml:"matchEquals"`
}

