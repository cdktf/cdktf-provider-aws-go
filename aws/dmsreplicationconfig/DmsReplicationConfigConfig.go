// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsreplicationconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsReplicationConfigConfig struct {
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
	// compute_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_replication_config#compute_config DmsReplicationConfig#compute_config}
	ComputeConfig *DmsReplicationConfigComputeConfig `field:"required" json:"computeConfig" yaml:"computeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_replication_config#replication_config_identifier DmsReplicationConfig#replication_config_identifier}.
	ReplicationConfigIdentifier *string `field:"required" json:"replicationConfigIdentifier" yaml:"replicationConfigIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_replication_config#replication_type DmsReplicationConfig#replication_type}.
	ReplicationType *string `field:"required" json:"replicationType" yaml:"replicationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_replication_config#source_endpoint_arn DmsReplicationConfig#source_endpoint_arn}.
	SourceEndpointArn *string `field:"required" json:"sourceEndpointArn" yaml:"sourceEndpointArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_replication_config#table_mappings DmsReplicationConfig#table_mappings}.
	TableMappings *string `field:"required" json:"tableMappings" yaml:"tableMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_replication_config#target_endpoint_arn DmsReplicationConfig#target_endpoint_arn}.
	TargetEndpointArn *string `field:"required" json:"targetEndpointArn" yaml:"targetEndpointArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_replication_config#id DmsReplicationConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_replication_config#replication_settings DmsReplicationConfig#replication_settings}.
	ReplicationSettings *string `field:"optional" json:"replicationSettings" yaml:"replicationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_replication_config#resource_identifier DmsReplicationConfig#resource_identifier}.
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_replication_config#start_replication DmsReplicationConfig#start_replication}.
	StartReplication interface{} `field:"optional" json:"startReplication" yaml:"startReplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_replication_config#supplemental_settings DmsReplicationConfig#supplemental_settings}.
	SupplementalSettings *string `field:"optional" json:"supplementalSettings" yaml:"supplementalSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_replication_config#tags DmsReplicationConfig#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_replication_config#tags_all DmsReplicationConfig#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/dms_replication_config#timeouts DmsReplicationConfig#timeouts}
	Timeouts *DmsReplicationConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

