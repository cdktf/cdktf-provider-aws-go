// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakeformationpermissions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LakeformationPermissionsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/lakeformation_permissions#permissions LakeformationPermissions#permissions}.
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/lakeformation_permissions#principal LakeformationPermissions#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/lakeformation_permissions#catalog_id LakeformationPermissions#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/lakeformation_permissions#catalog_resource LakeformationPermissions#catalog_resource}.
	CatalogResource interface{} `field:"optional" json:"catalogResource" yaml:"catalogResource"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/lakeformation_permissions#database LakeformationPermissions#database}
	Database *LakeformationPermissionsDatabase `field:"optional" json:"database" yaml:"database"`
	// data_cells_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/lakeformation_permissions#data_cells_filter LakeformationPermissions#data_cells_filter}
	DataCellsFilter *LakeformationPermissionsDataCellsFilter `field:"optional" json:"dataCellsFilter" yaml:"dataCellsFilter"`
	// data_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/lakeformation_permissions#data_location LakeformationPermissions#data_location}
	DataLocation *LakeformationPermissionsDataLocation `field:"optional" json:"dataLocation" yaml:"dataLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/lakeformation_permissions#id LakeformationPermissions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lf_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/lakeformation_permissions#lf_tag LakeformationPermissions#lf_tag}
	LfTag *LakeformationPermissionsLfTag `field:"optional" json:"lfTag" yaml:"lfTag"`
	// lf_tag_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/lakeformation_permissions#lf_tag_policy LakeformationPermissions#lf_tag_policy}
	LfTagPolicy *LakeformationPermissionsLfTagPolicy `field:"optional" json:"lfTagPolicy" yaml:"lfTagPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/lakeformation_permissions#permissions_with_grant_option LakeformationPermissions#permissions_with_grant_option}.
	PermissionsWithGrantOption *[]*string `field:"optional" json:"permissionsWithGrantOption" yaml:"permissionsWithGrantOption"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/lakeformation_permissions#table LakeformationPermissions#table}
	Table *LakeformationPermissionsTable `field:"optional" json:"table" yaml:"table"`
	// table_with_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/lakeformation_permissions#table_with_columns LakeformationPermissions#table_with_columns}
	TableWithColumns *LakeformationPermissionsTableWithColumns `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
}

