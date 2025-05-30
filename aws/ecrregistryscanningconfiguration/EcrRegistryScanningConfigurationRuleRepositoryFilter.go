// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecrregistryscanningconfiguration


type EcrRegistryScanningConfigurationRuleRepositoryFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/ecr_registry_scanning_configuration#filter EcrRegistryScanningConfiguration#filter}.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/ecr_registry_scanning_configuration#filter_type EcrRegistryScanningConfiguration#filter_type}.
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
}

