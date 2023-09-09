// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamwritetable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TimestreamwriteTableConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/timestreamwrite_table#database_name TimestreamwriteTable#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/timestreamwrite_table#table_name TimestreamwriteTable#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/timestreamwrite_table#id TimestreamwriteTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// magnetic_store_write_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/timestreamwrite_table#magnetic_store_write_properties TimestreamwriteTable#magnetic_store_write_properties}
	MagneticStoreWriteProperties *TimestreamwriteTableMagneticStoreWriteProperties `field:"optional" json:"magneticStoreWriteProperties" yaml:"magneticStoreWriteProperties"`
	// retention_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/timestreamwrite_table#retention_properties TimestreamwriteTable#retention_properties}
	RetentionProperties *TimestreamwriteTableRetentionProperties `field:"optional" json:"retentionProperties" yaml:"retentionProperties"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/timestreamwrite_table#schema TimestreamwriteTable#schema}
	Schema *TimestreamwriteTableSchema `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/timestreamwrite_table#tags TimestreamwriteTable#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/timestreamwrite_table#tags_all TimestreamwriteTable#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

