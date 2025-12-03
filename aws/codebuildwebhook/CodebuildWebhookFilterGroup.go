// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildwebhook


type CodebuildWebhookFilterGroup struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/codebuild_webhook#filter CodebuildWebhook#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

