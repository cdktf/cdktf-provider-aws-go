// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2instancetypeoffering


type DataAwsEc2InstanceTypeOfferingFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/data-sources/ec2_instance_type_offering#name DataAwsEc2InstanceTypeOffering#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/data-sources/ec2_instance_type_offering#values DataAwsEc2InstanceTypeOffering#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

