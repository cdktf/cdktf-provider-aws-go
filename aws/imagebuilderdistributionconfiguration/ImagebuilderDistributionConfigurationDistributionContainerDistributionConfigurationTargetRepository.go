// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributionContainerDistributionConfigurationTargetRepository struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/imagebuilder_distribution_configuration#repository_name ImagebuilderDistributionConfiguration#repository_name}.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/imagebuilder_distribution_configuration#service ImagebuilderDistributionConfiguration#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
}

