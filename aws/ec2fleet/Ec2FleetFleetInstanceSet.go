// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2fleet


type Ec2FleetFleetInstanceSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/ec2_fleet#instance_ids Ec2Fleet#instance_ids}.
	InstanceIds *[]*string `field:"optional" json:"instanceIds" yaml:"instanceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/ec2_fleet#instance_type Ec2Fleet#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/ec2_fleet#lifecycle Ec2Fleet#lifecycle}.
	Lifecycle *string `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/ec2_fleet#platform Ec2Fleet#platform}.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
}

