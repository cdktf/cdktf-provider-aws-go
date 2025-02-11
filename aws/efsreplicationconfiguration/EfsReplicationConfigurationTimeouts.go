// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package efsreplicationconfiguration


type EfsReplicationConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/efs_replication_configuration#create EfsReplicationConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/efs_replication_configuration#delete EfsReplicationConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

