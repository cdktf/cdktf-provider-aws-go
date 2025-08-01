// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecrreplicationconfiguration


type EcrReplicationConfigurationReplicationConfigurationRuleRepositoryFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ecr_replication_configuration#filter EcrReplicationConfiguration#filter}.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ecr_replication_configuration#filter_type EcrReplicationConfiguration#filter_type}.
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
}

