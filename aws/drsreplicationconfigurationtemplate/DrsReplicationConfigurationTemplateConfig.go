// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package drsreplicationconfigurationtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DrsReplicationConfigurationTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#associate_default_security_group DrsReplicationConfigurationTemplate#associate_default_security_group}.
	AssociateDefaultSecurityGroup interface{} `field:"required" json:"associateDefaultSecurityGroup" yaml:"associateDefaultSecurityGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#bandwidth_throttling DrsReplicationConfigurationTemplate#bandwidth_throttling}.
	BandwidthThrottling *float64 `field:"required" json:"bandwidthThrottling" yaml:"bandwidthThrottling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#create_public_ip DrsReplicationConfigurationTemplate#create_public_ip}.
	CreatePublicIp interface{} `field:"required" json:"createPublicIp" yaml:"createPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#data_plane_routing DrsReplicationConfigurationTemplate#data_plane_routing}.
	DataPlaneRouting *string `field:"required" json:"dataPlaneRouting" yaml:"dataPlaneRouting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#default_large_staging_disk_type DrsReplicationConfigurationTemplate#default_large_staging_disk_type}.
	DefaultLargeStagingDiskType *string `field:"required" json:"defaultLargeStagingDiskType" yaml:"defaultLargeStagingDiskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#ebs_encryption DrsReplicationConfigurationTemplate#ebs_encryption}.
	EbsEncryption *string `field:"required" json:"ebsEncryption" yaml:"ebsEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#replication_server_instance_type DrsReplicationConfigurationTemplate#replication_server_instance_type}.
	ReplicationServerInstanceType *string `field:"required" json:"replicationServerInstanceType" yaml:"replicationServerInstanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#replication_servers_security_groups_ids DrsReplicationConfigurationTemplate#replication_servers_security_groups_ids}.
	ReplicationServersSecurityGroupsIds *[]*string `field:"required" json:"replicationServersSecurityGroupsIds" yaml:"replicationServersSecurityGroupsIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#staging_area_subnet_id DrsReplicationConfigurationTemplate#staging_area_subnet_id}.
	StagingAreaSubnetId *string `field:"required" json:"stagingAreaSubnetId" yaml:"stagingAreaSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#staging_area_tags DrsReplicationConfigurationTemplate#staging_area_tags}.
	StagingAreaTags *map[string]*string `field:"required" json:"stagingAreaTags" yaml:"stagingAreaTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#use_dedicated_replication_server DrsReplicationConfigurationTemplate#use_dedicated_replication_server}.
	UseDedicatedReplicationServer interface{} `field:"required" json:"useDedicatedReplicationServer" yaml:"useDedicatedReplicationServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#auto_replicate_new_disks DrsReplicationConfigurationTemplate#auto_replicate_new_disks}.
	AutoReplicateNewDisks interface{} `field:"optional" json:"autoReplicateNewDisks" yaml:"autoReplicateNewDisks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#ebs_encryption_key_arn DrsReplicationConfigurationTemplate#ebs_encryption_key_arn}.
	EbsEncryptionKeyArn *string `field:"optional" json:"ebsEncryptionKeyArn" yaml:"ebsEncryptionKeyArn"`
	// pit_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#pit_policy DrsReplicationConfigurationTemplate#pit_policy}
	PitPolicy interface{} `field:"optional" json:"pitPolicy" yaml:"pitPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#tags DrsReplicationConfigurationTemplate#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/drs_replication_configuration_template#timeouts DrsReplicationConfigurationTemplate#timeouts}
	Timeouts *DrsReplicationConfigurationTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

