// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2fleet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2FleetConfig struct {
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
	// launch_template_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#launch_template_config Ec2Fleet#launch_template_config}
	LaunchTemplateConfig interface{} `field:"required" json:"launchTemplateConfig" yaml:"launchTemplateConfig"`
	// target_capacity_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#target_capacity_specification Ec2Fleet#target_capacity_specification}
	TargetCapacitySpecification *Ec2FleetTargetCapacitySpecification `field:"required" json:"targetCapacitySpecification" yaml:"targetCapacitySpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#context Ec2Fleet#context}.
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#excess_capacity_termination_policy Ec2Fleet#excess_capacity_termination_policy}.
	ExcessCapacityTerminationPolicy *string `field:"optional" json:"excessCapacityTerminationPolicy" yaml:"excessCapacityTerminationPolicy"`
	// fleet_instance_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#fleet_instance_set Ec2Fleet#fleet_instance_set}
	FleetInstanceSet interface{} `field:"optional" json:"fleetInstanceSet" yaml:"fleetInstanceSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#fleet_state Ec2Fleet#fleet_state}.
	FleetState *string `field:"optional" json:"fleetState" yaml:"fleetState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#fulfilled_capacity Ec2Fleet#fulfilled_capacity}.
	FulfilledCapacity *float64 `field:"optional" json:"fulfilledCapacity" yaml:"fulfilledCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#fulfilled_on_demand_capacity Ec2Fleet#fulfilled_on_demand_capacity}.
	FulfilledOnDemandCapacity *float64 `field:"optional" json:"fulfilledOnDemandCapacity" yaml:"fulfilledOnDemandCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#id Ec2Fleet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// on_demand_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#on_demand_options Ec2Fleet#on_demand_options}
	OnDemandOptions *Ec2FleetOnDemandOptions `field:"optional" json:"onDemandOptions" yaml:"onDemandOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#replace_unhealthy_instances Ec2Fleet#replace_unhealthy_instances}.
	ReplaceUnhealthyInstances interface{} `field:"optional" json:"replaceUnhealthyInstances" yaml:"replaceUnhealthyInstances"`
	// spot_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#spot_options Ec2Fleet#spot_options}
	SpotOptions *Ec2FleetSpotOptions `field:"optional" json:"spotOptions" yaml:"spotOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#tags Ec2Fleet#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#tags_all Ec2Fleet#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#terminate_instances Ec2Fleet#terminate_instances}.
	TerminateInstances interface{} `field:"optional" json:"terminateInstances" yaml:"terminateInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#terminate_instances_with_expiration Ec2Fleet#terminate_instances_with_expiration}.
	TerminateInstancesWithExpiration interface{} `field:"optional" json:"terminateInstancesWithExpiration" yaml:"terminateInstancesWithExpiration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#timeouts Ec2Fleet#timeouts}
	Timeouts *Ec2FleetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#type Ec2Fleet#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#valid_from Ec2Fleet#valid_from}.
	ValidFrom *string `field:"optional" json:"validFrom" yaml:"validFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ec2_fleet#valid_until Ec2Fleet#valid_until}.
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}

