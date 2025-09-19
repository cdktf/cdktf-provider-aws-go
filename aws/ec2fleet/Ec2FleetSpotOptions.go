// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2fleet


type Ec2FleetSpotOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ec2_fleet#allocation_strategy Ec2Fleet#allocation_strategy}.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ec2_fleet#instance_interruption_behavior Ec2Fleet#instance_interruption_behavior}.
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ec2_fleet#instance_pools_to_use_count Ec2Fleet#instance_pools_to_use_count}.
	InstancePoolsToUseCount *float64 `field:"optional" json:"instancePoolsToUseCount" yaml:"instancePoolsToUseCount"`
	// maintenance_strategies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ec2_fleet#maintenance_strategies Ec2Fleet#maintenance_strategies}
	MaintenanceStrategies *Ec2FleetSpotOptionsMaintenanceStrategies `field:"optional" json:"maintenanceStrategies" yaml:"maintenanceStrategies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ec2_fleet#max_total_price Ec2Fleet#max_total_price}.
	MaxTotalPrice *string `field:"optional" json:"maxTotalPrice" yaml:"maxTotalPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ec2_fleet#min_target_capacity Ec2Fleet#min_target_capacity}.
	MinTargetCapacity *float64 `field:"optional" json:"minTargetCapacity" yaml:"minTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ec2_fleet#single_availability_zone Ec2Fleet#single_availability_zone}.
	SingleAvailabilityZone interface{} `field:"optional" json:"singleAvailabilityZone" yaml:"singleAvailabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ec2_fleet#single_instance_type Ec2Fleet#single_instance_type}.
	SingleInstanceType interface{} `field:"optional" json:"singleInstanceType" yaml:"singleInstanceType"`
}

