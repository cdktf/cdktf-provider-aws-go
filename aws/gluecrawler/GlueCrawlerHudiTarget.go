// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecrawler


type GlueCrawlerHudiTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/glue_crawler#maximum_traversal_depth GlueCrawler#maximum_traversal_depth}.
	MaximumTraversalDepth *float64 `field:"required" json:"maximumTraversalDepth" yaml:"maximumTraversalDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/glue_crawler#paths GlueCrawler#paths}.
	Paths *[]*string `field:"required" json:"paths" yaml:"paths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/glue_crawler#connection_name GlueCrawler#connection_name}.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/glue_crawler#exclusions GlueCrawler#exclusions}.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
}

