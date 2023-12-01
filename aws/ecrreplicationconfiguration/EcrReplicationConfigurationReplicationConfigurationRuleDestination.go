// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecrreplicationconfiguration


type EcrReplicationConfigurationReplicationConfigurationRuleDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/ecr_replication_configuration#region EcrReplicationConfiguration#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/ecr_replication_configuration#registry_id EcrReplicationConfiguration#registry_id}.
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
}

