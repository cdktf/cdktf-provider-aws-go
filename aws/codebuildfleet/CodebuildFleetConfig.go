// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildfleet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodebuildFleetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/codebuild_fleet#base_capacity CodebuildFleet#base_capacity}.
	BaseCapacity *float64 `field:"required" json:"baseCapacity" yaml:"baseCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/codebuild_fleet#compute_type CodebuildFleet#compute_type}.
	ComputeType *string `field:"required" json:"computeType" yaml:"computeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/codebuild_fleet#environment_type CodebuildFleet#environment_type}.
	EnvironmentType *string `field:"required" json:"environmentType" yaml:"environmentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/codebuild_fleet#name CodebuildFleet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/codebuild_fleet#fleet_service_role CodebuildFleet#fleet_service_role}.
	FleetServiceRole *string `field:"optional" json:"fleetServiceRole" yaml:"fleetServiceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/codebuild_fleet#image_id CodebuildFleet#image_id}.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/codebuild_fleet#overflow_behavior CodebuildFleet#overflow_behavior}.
	OverflowBehavior *string `field:"optional" json:"overflowBehavior" yaml:"overflowBehavior"`
	// scaling_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/codebuild_fleet#scaling_configuration CodebuildFleet#scaling_configuration}
	ScalingConfiguration *CodebuildFleetScalingConfiguration `field:"optional" json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/codebuild_fleet#tags CodebuildFleet#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/codebuild_fleet#tags_all CodebuildFleet#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/codebuild_fleet#vpc_config CodebuildFleet#vpc_config}
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

