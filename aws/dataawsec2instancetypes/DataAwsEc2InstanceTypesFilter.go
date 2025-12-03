// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2instancetypes


type DataAwsEc2InstanceTypesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/data-sources/ec2_instance_types#name DataAwsEc2InstanceTypes#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/data-sources/ec2_instance_types#values DataAwsEc2InstanceTypes#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

