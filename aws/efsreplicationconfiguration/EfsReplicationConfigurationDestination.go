// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package efsreplicationconfiguration


type EfsReplicationConfigurationDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/efs_replication_configuration#availability_zone_name EfsReplicationConfiguration#availability_zone_name}.
	AvailabilityZoneName *string `field:"optional" json:"availabilityZoneName" yaml:"availabilityZoneName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/efs_replication_configuration#file_system_id EfsReplicationConfiguration#file_system_id}.
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/efs_replication_configuration#kms_key_id EfsReplicationConfiguration#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/efs_replication_configuration#region EfsReplicationConfiguration#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

