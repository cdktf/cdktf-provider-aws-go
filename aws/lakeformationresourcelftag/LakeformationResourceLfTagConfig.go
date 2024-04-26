// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakeformationresourcelftag

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LakeformationResourceLfTagConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/lakeformation_resource_lf_tag#catalog_id LakeformationResourceLfTag#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/lakeformation_resource_lf_tag#database LakeformationResourceLfTag#database}
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// lf_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/lakeformation_resource_lf_tag#lf_tag LakeformationResourceLfTag#lf_tag}
	LfTag interface{} `field:"optional" json:"lfTag" yaml:"lfTag"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/lakeformation_resource_lf_tag#table LakeformationResourceLfTag#table}
	Table interface{} `field:"optional" json:"table" yaml:"table"`
	// table_with_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/lakeformation_resource_lf_tag#table_with_columns LakeformationResourceLfTag#table_with_columns}
	TableWithColumns interface{} `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/lakeformation_resource_lf_tag#timeouts LakeformationResourceLfTag#timeouts}
	Timeouts *LakeformationResourceLfTagTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

