// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecrawler


type GlueCrawlerRecrawlPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/glue_crawler#recrawl_behavior GlueCrawler#recrawl_behavior}.
	RecrawlBehavior *string `field:"optional" json:"recrawlBehavior" yaml:"recrawlBehavior"`
}

