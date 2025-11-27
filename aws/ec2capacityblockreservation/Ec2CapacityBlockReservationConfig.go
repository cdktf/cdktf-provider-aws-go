// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2capacityblockreservation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2CapacityBlockReservationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ec2_capacity_block_reservation#capacity_block_offering_id Ec2CapacityBlockReservation#capacity_block_offering_id}.
	CapacityBlockOfferingId *string `field:"required" json:"capacityBlockOfferingId" yaml:"capacityBlockOfferingId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ec2_capacity_block_reservation#instance_platform Ec2CapacityBlockReservation#instance_platform}.
	InstancePlatform *string `field:"required" json:"instancePlatform" yaml:"instancePlatform"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ec2_capacity_block_reservation#region Ec2CapacityBlockReservation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ec2_capacity_block_reservation#tags Ec2CapacityBlockReservation#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ec2_capacity_block_reservation#timeouts Ec2CapacityBlockReservation#timeouts}
	Timeouts *Ec2CapacityBlockReservationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

