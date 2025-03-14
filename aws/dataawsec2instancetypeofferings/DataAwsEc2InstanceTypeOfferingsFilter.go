// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2instancetypeofferings


type DataAwsEc2InstanceTypeOfferingsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/data-sources/ec2_instance_type_offerings#name DataAwsEc2InstanceTypeOfferings#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/data-sources/ec2_instance_type_offerings#values DataAwsEc2InstanceTypeOfferings#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

