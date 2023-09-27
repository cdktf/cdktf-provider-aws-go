// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsimagebuilderinfrastructureconfigurations


type DataAwsImagebuilderInfrastructureConfigurationsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/imagebuilder_infrastructure_configurations#name DataAwsImagebuilderInfrastructureConfigurations#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/imagebuilder_infrastructure_configurations#values DataAwsImagebuilderInfrastructureConfigurations#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

