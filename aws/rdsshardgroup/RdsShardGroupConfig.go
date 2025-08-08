// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsshardgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsShardGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/rds_shard_group#db_cluster_identifier RdsShardGroup#db_cluster_identifier}.
	DbClusterIdentifier *string `field:"required" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/rds_shard_group#db_shard_group_identifier RdsShardGroup#db_shard_group_identifier}.
	DbShardGroupIdentifier *string `field:"required" json:"dbShardGroupIdentifier" yaml:"dbShardGroupIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/rds_shard_group#max_acu RdsShardGroup#max_acu}.
	MaxAcu *float64 `field:"required" json:"maxAcu" yaml:"maxAcu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/rds_shard_group#compute_redundancy RdsShardGroup#compute_redundancy}.
	ComputeRedundancy *float64 `field:"optional" json:"computeRedundancy" yaml:"computeRedundancy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/rds_shard_group#min_acu RdsShardGroup#min_acu}.
	MinAcu *float64 `field:"optional" json:"minAcu" yaml:"minAcu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/rds_shard_group#publicly_accessible RdsShardGroup#publicly_accessible}.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/rds_shard_group#region RdsShardGroup#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/rds_shard_group#tags RdsShardGroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/rds_shard_group#timeouts RdsShardGroup#timeouts}
	Timeouts *RdsShardGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

