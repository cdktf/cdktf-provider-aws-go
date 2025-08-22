// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2fleet


type Ec2FleetLaunchTemplateConfigOverride struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/ec2_fleet#availability_zone Ec2Fleet#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// instance_requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/ec2_fleet#instance_requirements Ec2Fleet#instance_requirements}
	InstanceRequirements *Ec2FleetLaunchTemplateConfigOverrideInstanceRequirements `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/ec2_fleet#instance_type Ec2Fleet#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/ec2_fleet#max_price Ec2Fleet#max_price}.
	MaxPrice *string `field:"optional" json:"maxPrice" yaml:"maxPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/ec2_fleet#priority Ec2Fleet#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/ec2_fleet#subnet_id Ec2Fleet#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/ec2_fleet#weighted_capacity Ec2Fleet#weighted_capacity}.
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

