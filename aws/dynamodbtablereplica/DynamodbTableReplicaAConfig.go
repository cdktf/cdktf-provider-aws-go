// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbtablereplica

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DynamodbTableReplicaAConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/dynamodb_table_replica#global_table_arn DynamodbTableReplicaA#global_table_arn}.
	GlobalTableArn *string `field:"required" json:"globalTableArn" yaml:"globalTableArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/dynamodb_table_replica#id DynamodbTableReplicaA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/dynamodb_table_replica#kms_key_arn DynamodbTableReplicaA#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/dynamodb_table_replica#point_in_time_recovery DynamodbTableReplicaA#point_in_time_recovery}.
	PointInTimeRecovery interface{} `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/dynamodb_table_replica#table_class_override DynamodbTableReplicaA#table_class_override}.
	TableClassOverride *string `field:"optional" json:"tableClassOverride" yaml:"tableClassOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/dynamodb_table_replica#tags DynamodbTableReplicaA#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/dynamodb_table_replica#tags_all DynamodbTableReplicaA#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/dynamodb_table_replica#timeouts DynamodbTableReplicaA#timeouts}
	Timeouts *DynamodbTableReplicaTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

