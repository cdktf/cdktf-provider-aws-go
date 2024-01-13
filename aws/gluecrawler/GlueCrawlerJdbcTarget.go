// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecrawler


type GlueCrawlerJdbcTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/glue_crawler#connection_name GlueCrawler#connection_name}.
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/glue_crawler#path GlueCrawler#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/glue_crawler#enable_additional_metadata GlueCrawler#enable_additional_metadata}.
	EnableAdditionalMetadata *[]*string `field:"optional" json:"enableAdditionalMetadata" yaml:"enableAdditionalMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/glue_crawler#exclusions GlueCrawler#exclusions}.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
}

