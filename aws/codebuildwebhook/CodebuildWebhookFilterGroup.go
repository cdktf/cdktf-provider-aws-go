// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildwebhook


type CodebuildWebhookFilterGroup struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/codebuild_webhook#filter CodebuildWebhook#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

