// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslakeformationpermissions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsLakeformationPermissionsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/data-sources/lakeformation_permissions#principal DataAwsLakeformationPermissions#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/data-sources/lakeformation_permissions#catalog_id DataAwsLakeformationPermissions#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/data-sources/lakeformation_permissions#catalog_resource DataAwsLakeformationPermissions#catalog_resource}.
	CatalogResource interface{} `field:"optional" json:"catalogResource" yaml:"catalogResource"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/data-sources/lakeformation_permissions#database DataAwsLakeformationPermissions#database}
	Database *DataAwsLakeformationPermissionsDatabase `field:"optional" json:"database" yaml:"database"`
	// data_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/data-sources/lakeformation_permissions#data_location DataAwsLakeformationPermissions#data_location}
	DataLocation *DataAwsLakeformationPermissionsDataLocation `field:"optional" json:"dataLocation" yaml:"dataLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/data-sources/lakeformation_permissions#id DataAwsLakeformationPermissions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lf_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/data-sources/lakeformation_permissions#lf_tag DataAwsLakeformationPermissions#lf_tag}
	LfTag *DataAwsLakeformationPermissionsLfTag `field:"optional" json:"lfTag" yaml:"lfTag"`
	// lf_tag_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/data-sources/lakeformation_permissions#lf_tag_policy DataAwsLakeformationPermissions#lf_tag_policy}
	LfTagPolicy *DataAwsLakeformationPermissionsLfTagPolicy `field:"optional" json:"lfTagPolicy" yaml:"lfTagPolicy"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/data-sources/lakeformation_permissions#table DataAwsLakeformationPermissions#table}
	Table *DataAwsLakeformationPermissionsTable `field:"optional" json:"table" yaml:"table"`
	// table_with_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/data-sources/lakeformation_permissions#table_with_columns DataAwsLakeformationPermissions#table_with_columns}
	TableWithColumns *DataAwsLakeformationPermissionsTableWithColumns `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
}

