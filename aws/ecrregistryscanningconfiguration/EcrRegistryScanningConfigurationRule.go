// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecrregistryscanningconfiguration


type EcrRegistryScanningConfigurationRule struct {
	// repository_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ecr_registry_scanning_configuration#repository_filter EcrRegistryScanningConfiguration#repository_filter}
	RepositoryFilter interface{} `field:"required" json:"repositoryFilter" yaml:"repositoryFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ecr_registry_scanning_configuration#scan_frequency EcrRegistryScanningConfiguration#scan_frequency}.
	ScanFrequency *string `field:"required" json:"scanFrequency" yaml:"scanFrequency"`
}

