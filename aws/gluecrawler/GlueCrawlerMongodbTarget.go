// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecrawler


type GlueCrawlerMongodbTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/glue_crawler#connection_name GlueCrawler#connection_name}.
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/glue_crawler#path GlueCrawler#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/glue_crawler#scan_all GlueCrawler#scan_all}.
	ScanAll interface{} `field:"optional" json:"scanAll" yaml:"scanAll"`
}

