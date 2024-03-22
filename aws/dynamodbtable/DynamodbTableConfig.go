// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DynamodbTableConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#name DynamodbTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// attribute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#attribute DynamodbTable#attribute}
	Attribute interface{} `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#billing_mode DynamodbTable#billing_mode}.
	BillingMode *string `field:"optional" json:"billingMode" yaml:"billingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#deletion_protection_enabled DynamodbTable#deletion_protection_enabled}.
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// global_secondary_index block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#global_secondary_index DynamodbTable#global_secondary_index}
	GlobalSecondaryIndex interface{} `field:"optional" json:"globalSecondaryIndex" yaml:"globalSecondaryIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#hash_key DynamodbTable#hash_key}.
	HashKey *string `field:"optional" json:"hashKey" yaml:"hashKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#id DynamodbTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// import_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#import_table DynamodbTable#import_table}
	ImportTable *DynamodbTableImportTable `field:"optional" json:"importTable" yaml:"importTable"`
	// local_secondary_index block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#local_secondary_index DynamodbTable#local_secondary_index}
	LocalSecondaryIndex interface{} `field:"optional" json:"localSecondaryIndex" yaml:"localSecondaryIndex"`
	// point_in_time_recovery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#point_in_time_recovery DynamodbTable#point_in_time_recovery}
	PointInTimeRecovery *DynamodbTablePointInTimeRecovery `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#range_key DynamodbTable#range_key}.
	RangeKey *string `field:"optional" json:"rangeKey" yaml:"rangeKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#read_capacity DynamodbTable#read_capacity}.
	ReadCapacity *float64 `field:"optional" json:"readCapacity" yaml:"readCapacity"`
	// replica block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#replica DynamodbTable#replica}
	Replica interface{} `field:"optional" json:"replica" yaml:"replica"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#restore_date_time DynamodbTable#restore_date_time}.
	RestoreDateTime *string `field:"optional" json:"restoreDateTime" yaml:"restoreDateTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#restore_source_name DynamodbTable#restore_source_name}.
	RestoreSourceName *string `field:"optional" json:"restoreSourceName" yaml:"restoreSourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#restore_to_latest_time DynamodbTable#restore_to_latest_time}.
	RestoreToLatestTime interface{} `field:"optional" json:"restoreToLatestTime" yaml:"restoreToLatestTime"`
	// server_side_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#server_side_encryption DynamodbTable#server_side_encryption}
	ServerSideEncryption *DynamodbTableServerSideEncryption `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#stream_enabled DynamodbTable#stream_enabled}.
	StreamEnabled interface{} `field:"optional" json:"streamEnabled" yaml:"streamEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#stream_view_type DynamodbTable#stream_view_type}.
	StreamViewType *string `field:"optional" json:"streamViewType" yaml:"streamViewType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#table_class DynamodbTable#table_class}.
	TableClass *string `field:"optional" json:"tableClass" yaml:"tableClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#tags DynamodbTable#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#tags_all DynamodbTable#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#timeouts DynamodbTable#timeouts}
	Timeouts *DynamodbTableTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#ttl DynamodbTable#ttl}
	Ttl *DynamodbTableTtl `field:"optional" json:"ttl" yaml:"ttl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/dynamodb_table#write_capacity DynamodbTable#write_capacity}.
	WriteCapacity *float64 `field:"optional" json:"writeCapacity" yaml:"writeCapacity"`
}

