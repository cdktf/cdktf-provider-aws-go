// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotfleetrequest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpotFleetRequestConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#iam_fleet_role SpotFleetRequest#iam_fleet_role}.
	IamFleetRole *string `field:"required" json:"iamFleetRole" yaml:"iamFleetRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#target_capacity SpotFleetRequest#target_capacity}.
	TargetCapacity *float64 `field:"required" json:"targetCapacity" yaml:"targetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#allocation_strategy SpotFleetRequest#allocation_strategy}.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#context SpotFleetRequest#context}.
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#excess_capacity_termination_policy SpotFleetRequest#excess_capacity_termination_policy}.
	ExcessCapacityTerminationPolicy *string `field:"optional" json:"excessCapacityTerminationPolicy" yaml:"excessCapacityTerminationPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#fleet_type SpotFleetRequest#fleet_type}.
	FleetType *string `field:"optional" json:"fleetType" yaml:"fleetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#id SpotFleetRequest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#instance_interruption_behaviour SpotFleetRequest#instance_interruption_behaviour}.
	InstanceInterruptionBehaviour *string `field:"optional" json:"instanceInterruptionBehaviour" yaml:"instanceInterruptionBehaviour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#instance_pools_to_use_count SpotFleetRequest#instance_pools_to_use_count}.
	InstancePoolsToUseCount *float64 `field:"optional" json:"instancePoolsToUseCount" yaml:"instancePoolsToUseCount"`
	// launch_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#launch_specification SpotFleetRequest#launch_specification}
	LaunchSpecification interface{} `field:"optional" json:"launchSpecification" yaml:"launchSpecification"`
	// launch_template_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#launch_template_config SpotFleetRequest#launch_template_config}
	LaunchTemplateConfig interface{} `field:"optional" json:"launchTemplateConfig" yaml:"launchTemplateConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#load_balancers SpotFleetRequest#load_balancers}.
	LoadBalancers *[]*string `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#on_demand_allocation_strategy SpotFleetRequest#on_demand_allocation_strategy}.
	OnDemandAllocationStrategy *string `field:"optional" json:"onDemandAllocationStrategy" yaml:"onDemandAllocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#on_demand_max_total_price SpotFleetRequest#on_demand_max_total_price}.
	OnDemandMaxTotalPrice *string `field:"optional" json:"onDemandMaxTotalPrice" yaml:"onDemandMaxTotalPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#on_demand_target_capacity SpotFleetRequest#on_demand_target_capacity}.
	OnDemandTargetCapacity *float64 `field:"optional" json:"onDemandTargetCapacity" yaml:"onDemandTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#replace_unhealthy_instances SpotFleetRequest#replace_unhealthy_instances}.
	ReplaceUnhealthyInstances interface{} `field:"optional" json:"replaceUnhealthyInstances" yaml:"replaceUnhealthyInstances"`
	// spot_maintenance_strategies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#spot_maintenance_strategies SpotFleetRequest#spot_maintenance_strategies}
	SpotMaintenanceStrategies *SpotFleetRequestSpotMaintenanceStrategies `field:"optional" json:"spotMaintenanceStrategies" yaml:"spotMaintenanceStrategies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#spot_price SpotFleetRequest#spot_price}.
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#tags SpotFleetRequest#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#tags_all SpotFleetRequest#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#target_capacity_unit_type SpotFleetRequest#target_capacity_unit_type}.
	TargetCapacityUnitType *string `field:"optional" json:"targetCapacityUnitType" yaml:"targetCapacityUnitType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#target_group_arns SpotFleetRequest#target_group_arns}.
	TargetGroupArns *[]*string `field:"optional" json:"targetGroupArns" yaml:"targetGroupArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#terminate_instances_on_delete SpotFleetRequest#terminate_instances_on_delete}.
	TerminateInstancesOnDelete *string `field:"optional" json:"terminateInstancesOnDelete" yaml:"terminateInstancesOnDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#terminate_instances_with_expiration SpotFleetRequest#terminate_instances_with_expiration}.
	TerminateInstancesWithExpiration interface{} `field:"optional" json:"terminateInstancesWithExpiration" yaml:"terminateInstancesWithExpiration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#timeouts SpotFleetRequest#timeouts}
	Timeouts *SpotFleetRequestTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#valid_from SpotFleetRequest#valid_from}.
	ValidFrom *string `field:"optional" json:"validFrom" yaml:"validFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#valid_until SpotFleetRequest#valid_until}.
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#wait_for_fulfillment SpotFleetRequest#wait_for_fulfillment}.
	WaitForFulfillment interface{} `field:"optional" json:"waitForFulfillment" yaml:"waitForFulfillment"`
}

