// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecrawler


type GlueCrawlerDeltaTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/glue_crawler#delta_tables GlueCrawler#delta_tables}.
	DeltaTables *[]*string `field:"required" json:"deltaTables" yaml:"deltaTables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/glue_crawler#write_manifest GlueCrawler#write_manifest}.
	WriteManifest interface{} `field:"required" json:"writeManifest" yaml:"writeManifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/glue_crawler#connection_name GlueCrawler#connection_name}.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/glue_crawler#create_native_delta_table GlueCrawler#create_native_delta_table}.
	CreateNativeDeltaTable interface{} `field:"optional" json:"createNativeDeltaTable" yaml:"createNativeDeltaTable"`
}

