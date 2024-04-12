// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2host


type DataAwsEc2HostFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/ec2_host#name DataAwsEc2Host#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/ec2_host#values DataAwsEc2Host#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

