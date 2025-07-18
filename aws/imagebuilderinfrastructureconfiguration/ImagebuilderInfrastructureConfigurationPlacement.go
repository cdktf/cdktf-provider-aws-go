// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderinfrastructureconfiguration


type ImagebuilderInfrastructureConfigurationPlacement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/imagebuilder_infrastructure_configuration#availability_zone ImagebuilderInfrastructureConfiguration#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/imagebuilder_infrastructure_configuration#host_id ImagebuilderInfrastructureConfiguration#host_id}.
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/imagebuilder_infrastructure_configuration#host_resource_group_arn ImagebuilderInfrastructureConfiguration#host_resource_group_arn}.
	HostResourceGroupArn *string `field:"optional" json:"hostResourceGroupArn" yaml:"hostResourceGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/imagebuilder_infrastructure_configuration#tenancy ImagebuilderInfrastructureConfiguration#tenancy}.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

