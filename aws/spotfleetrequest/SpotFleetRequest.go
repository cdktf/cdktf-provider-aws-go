// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotfleetrequest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/spotfleetrequest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/spot_fleet_request aws_spot_fleet_request}.
type SpotFleetRequest interface {
	cdktf.TerraformResource
	AllocationStrategy() *string
	SetAllocationStrategy(val *string)
	AllocationStrategyInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientToken() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Context() *string
	SetContext(val *string)
	ContextInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExcessCapacityTerminationPolicy() *string
	SetExcessCapacityTerminationPolicy(val *string)
	ExcessCapacityTerminationPolicyInput() *string
	FleetType() *string
	SetFleetType(val *string)
	FleetTypeInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamFleetRole() *string
	SetIamFleetRole(val *string)
	IamFleetRoleInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceInterruptionBehaviour() *string
	SetInstanceInterruptionBehaviour(val *string)
	InstanceInterruptionBehaviourInput() *string
	InstancePoolsToUseCount() *float64
	SetInstancePoolsToUseCount(val *float64)
	InstancePoolsToUseCountInput() *float64
	LaunchSpecification() SpotFleetRequestLaunchSpecificationList
	LaunchSpecificationInput() interface{}
	LaunchTemplateConfig() SpotFleetRequestLaunchTemplateConfigList
	LaunchTemplateConfigInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancers() *[]*string
	SetLoadBalancers(val *[]*string)
	LoadBalancersInput() *[]*string
	// The tree node.
	Node() constructs.Node
	OnDemandAllocationStrategy() *string
	SetOnDemandAllocationStrategy(val *string)
	OnDemandAllocationStrategyInput() *string
	OnDemandMaxTotalPrice() *string
	SetOnDemandMaxTotalPrice(val *string)
	OnDemandMaxTotalPriceInput() *string
	OnDemandTargetCapacity() *float64
	SetOnDemandTargetCapacity(val *float64)
	OnDemandTargetCapacityInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReplaceUnhealthyInstances() interface{}
	SetReplaceUnhealthyInstances(val interface{})
	ReplaceUnhealthyInstancesInput() interface{}
	SpotMaintenanceStrategies() SpotFleetRequestSpotMaintenanceStrategiesOutputReference
	SpotMaintenanceStrategiesInput() *SpotFleetRequestSpotMaintenanceStrategies
	SpotPrice() *string
	SetSpotPrice(val *string)
	SpotPriceInput() *string
	SpotRequestState() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TargetCapacity() *float64
	SetTargetCapacity(val *float64)
	TargetCapacityInput() *float64
	TargetCapacityUnitType() *string
	SetTargetCapacityUnitType(val *string)
	TargetCapacityUnitTypeInput() *string
	TargetGroupArns() *[]*string
	SetTargetGroupArns(val *[]*string)
	TargetGroupArnsInput() *[]*string
	TerminateInstancesOnDelete() *string
	SetTerminateInstancesOnDelete(val *string)
	TerminateInstancesOnDeleteInput() *string
	TerminateInstancesWithExpiration() interface{}
	SetTerminateInstancesWithExpiration(val interface{})
	TerminateInstancesWithExpirationInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SpotFleetRequestTimeoutsOutputReference
	TimeoutsInput() interface{}
	ValidFrom() *string
	SetValidFrom(val *string)
	ValidFromInput() *string
	ValidUntil() *string
	SetValidUntil(val *string)
	ValidUntilInput() *string
	WaitForFulfillment() interface{}
	SetWaitForFulfillment(val interface{})
	WaitForFulfillmentInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutLaunchSpecification(value interface{})
	PutLaunchTemplateConfig(value interface{})
	PutSpotMaintenanceStrategies(value *SpotFleetRequestSpotMaintenanceStrategies)
	PutTimeouts(value *SpotFleetRequestTimeouts)
	ResetAllocationStrategy()
	ResetContext()
	ResetExcessCapacityTerminationPolicy()
	ResetFleetType()
	ResetId()
	ResetInstanceInterruptionBehaviour()
	ResetInstancePoolsToUseCount()
	ResetLaunchSpecification()
	ResetLaunchTemplateConfig()
	ResetLoadBalancers()
	ResetOnDemandAllocationStrategy()
	ResetOnDemandMaxTotalPrice()
	ResetOnDemandTargetCapacity()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReplaceUnhealthyInstances()
	ResetSpotMaintenanceStrategies()
	ResetSpotPrice()
	ResetTags()
	ResetTagsAll()
	ResetTargetCapacityUnitType()
	ResetTargetGroupArns()
	ResetTerminateInstancesOnDelete()
	ResetTerminateInstancesWithExpiration()
	ResetTimeouts()
	ResetValidFrom()
	ResetValidUntil()
	ResetWaitForFulfillment()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SpotFleetRequest
type jsiiProxy_SpotFleetRequest struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SpotFleetRequest) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) ExcessCapacityTerminationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excessCapacityTerminationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) ExcessCapacityTerminationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excessCapacityTerminationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) FleetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) FleetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) IamFleetRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamFleetRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) IamFleetRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamFleetRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) InstanceInterruptionBehaviour() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehaviour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) InstanceInterruptionBehaviourInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehaviourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) InstancePoolsToUseCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancePoolsToUseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) InstancePoolsToUseCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancePoolsToUseCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) LaunchSpecification() SpotFleetRequestLaunchSpecificationList {
	var returns SpotFleetRequestLaunchSpecificationList
	_jsii_.Get(
		j,
		"launchSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) LaunchSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) LaunchTemplateConfig() SpotFleetRequestLaunchTemplateConfigList {
	var returns SpotFleetRequestLaunchTemplateConfigList
	_jsii_.Get(
		j,
		"launchTemplateConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) LaunchTemplateConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) LoadBalancers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) LoadBalancersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) OnDemandAllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandAllocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) OnDemandAllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandAllocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) OnDemandMaxTotalPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandMaxTotalPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) OnDemandMaxTotalPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandMaxTotalPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) OnDemandTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) OnDemandTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) ReplaceUnhealthyInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceUnhealthyInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) ReplaceUnhealthyInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceUnhealthyInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) SpotMaintenanceStrategies() SpotFleetRequestSpotMaintenanceStrategiesOutputReference {
	var returns SpotFleetRequestSpotMaintenanceStrategiesOutputReference
	_jsii_.Get(
		j,
		"spotMaintenanceStrategies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) SpotMaintenanceStrategiesInput() *SpotFleetRequestSpotMaintenanceStrategies {
	var returns *SpotFleetRequestSpotMaintenanceStrategies
	_jsii_.Get(
		j,
		"spotMaintenanceStrategiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) SpotPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) SpotPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) SpotRequestState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotRequestState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TargetCapacityUnitType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityUnitType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TargetCapacityUnitTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityUnitTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TargetGroupArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TargetGroupArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TerminateInstancesOnDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminateInstancesOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TerminateInstancesOnDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminateInstancesOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TerminateInstancesWithExpiration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesWithExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TerminateInstancesWithExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesWithExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) Timeouts() SpotFleetRequestTimeoutsOutputReference {
	var returns SpotFleetRequestTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) ValidFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) ValidFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) ValidUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) ValidUntilInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntilInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) WaitForFulfillment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForFulfillment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotFleetRequest) WaitForFulfillmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForFulfillmentInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/spot_fleet_request aws_spot_fleet_request} Resource.
func NewSpotFleetRequest(scope constructs.Construct, id *string, config *SpotFleetRequestConfig) SpotFleetRequest {
	_init_.Initialize()

	if err := validateNewSpotFleetRequestParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpotFleetRequest{}

	_jsii_.Create(
		"@cdktf/provider-aws.spotFleetRequest.SpotFleetRequest",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/spot_fleet_request aws_spot_fleet_request} Resource.
func NewSpotFleetRequest_Override(s SpotFleetRequest, scope constructs.Construct, id *string, config *SpotFleetRequestConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.spotFleetRequest.SpotFleetRequest",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetAllocationStrategy(val *string) {
	if err := j.validateSetAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetExcessCapacityTerminationPolicy(val *string) {
	if err := j.validateSetExcessCapacityTerminationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excessCapacityTerminationPolicy",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetFleetType(val *string) {
	if err := j.validateSetFleetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fleetType",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetIamFleetRole(val *string) {
	if err := j.validateSetIamFleetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamFleetRole",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetInstanceInterruptionBehaviour(val *string) {
	if err := j.validateSetInstanceInterruptionBehaviourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInterruptionBehaviour",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetInstancePoolsToUseCount(val *float64) {
	if err := j.validateSetInstancePoolsToUseCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolsToUseCount",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetLoadBalancers(val *[]*string) {
	if err := j.validateSetLoadBalancersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetOnDemandAllocationStrategy(val *string) {
	if err := j.validateSetOnDemandAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandAllocationStrategy",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetOnDemandMaxTotalPrice(val *string) {
	if err := j.validateSetOnDemandMaxTotalPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandMaxTotalPrice",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetOnDemandTargetCapacity(val *float64) {
	if err := j.validateSetOnDemandTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetReplaceUnhealthyInstances(val interface{}) {
	if err := j.validateSetReplaceUnhealthyInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceUnhealthyInstances",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetSpotPrice(val *string) {
	if err := j.validateSetSpotPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotPrice",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetTargetCapacity(val *float64) {
	if err := j.validateSetTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCapacity",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetTargetCapacityUnitType(val *string) {
	if err := j.validateSetTargetCapacityUnitTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCapacityUnitType",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetTargetGroupArns(val *[]*string) {
	if err := j.validateSetTargetGroupArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupArns",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetTerminateInstancesOnDelete(val *string) {
	if err := j.validateSetTerminateInstancesOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateInstancesOnDelete",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetTerminateInstancesWithExpiration(val interface{}) {
	if err := j.validateSetTerminateInstancesWithExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateInstancesWithExpiration",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetValidFrom(val *string) {
	if err := j.validateSetValidFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validFrom",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetValidUntil(val *string) {
	if err := j.validateSetValidUntilParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validUntil",
		val,
	)
}

func (j *jsiiProxy_SpotFleetRequest)SetWaitForFulfillment(val interface{}) {
	if err := j.validateSetWaitForFulfillmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForFulfillment",
		val,
	)
}

// Generates CDKTF code for importing a SpotFleetRequest resource upon running "cdktf plan <stack-name>".
func SpotFleetRequest_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSpotFleetRequest_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.spotFleetRequest.SpotFleetRequest",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SpotFleetRequest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpotFleetRequest_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.spotFleetRequest.SpotFleetRequest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpotFleetRequest_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpotFleetRequest_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.spotFleetRequest.SpotFleetRequest",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpotFleetRequest_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpotFleetRequest_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.spotFleetRequest.SpotFleetRequest",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SpotFleetRequest_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.spotFleetRequest.SpotFleetRequest",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SpotFleetRequest) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SpotFleetRequest) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SpotFleetRequest) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SpotFleetRequest) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpotFleetRequest) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SpotFleetRequest) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpotFleetRequest) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SpotFleetRequest) PutLaunchSpecification(value interface{}) {
	if err := s.validatePutLaunchSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLaunchSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequest) PutLaunchTemplateConfig(value interface{}) {
	if err := s.validatePutLaunchTemplateConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLaunchTemplateConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequest) PutSpotMaintenanceStrategies(value *SpotFleetRequestSpotMaintenanceStrategies) {
	if err := s.validatePutSpotMaintenanceStrategiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSpotMaintenanceStrategies",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequest) PutTimeouts(value *SpotFleetRequestTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetAllocationStrategy() {
	_jsii_.InvokeVoid(
		s,
		"resetAllocationStrategy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetContext() {
	_jsii_.InvokeVoid(
		s,
		"resetContext",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetExcessCapacityTerminationPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetExcessCapacityTerminationPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetFleetType() {
	_jsii_.InvokeVoid(
		s,
		"resetFleetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetInstanceInterruptionBehaviour() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceInterruptionBehaviour",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetInstancePoolsToUseCount() {
	_jsii_.InvokeVoid(
		s,
		"resetInstancePoolsToUseCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetLaunchSpecification() {
	_jsii_.InvokeVoid(
		s,
		"resetLaunchSpecification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetLaunchTemplateConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetLaunchTemplateConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetLoadBalancers() {
	_jsii_.InvokeVoid(
		s,
		"resetLoadBalancers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetOnDemandAllocationStrategy() {
	_jsii_.InvokeVoid(
		s,
		"resetOnDemandAllocationStrategy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetOnDemandMaxTotalPrice() {
	_jsii_.InvokeVoid(
		s,
		"resetOnDemandMaxTotalPrice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetOnDemandTargetCapacity() {
	_jsii_.InvokeVoid(
		s,
		"resetOnDemandTargetCapacity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetReplaceUnhealthyInstances() {
	_jsii_.InvokeVoid(
		s,
		"resetReplaceUnhealthyInstances",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetSpotMaintenanceStrategies() {
	_jsii_.InvokeVoid(
		s,
		"resetSpotMaintenanceStrategies",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetSpotPrice() {
	_jsii_.InvokeVoid(
		s,
		"resetSpotPrice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetTargetCapacityUnitType() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetCapacityUnitType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetTargetGroupArns() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetGroupArns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetTerminateInstancesOnDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetTerminateInstancesOnDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetTerminateInstancesWithExpiration() {
	_jsii_.InvokeVoid(
		s,
		"resetTerminateInstancesWithExpiration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetValidFrom() {
	_jsii_.InvokeVoid(
		s,
		"resetValidFrom",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetValidUntil() {
	_jsii_.InvokeVoid(
		s,
		"resetValidUntil",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) ResetWaitForFulfillment() {
	_jsii_.InvokeVoid(
		s,
		"resetWaitForFulfillment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotFleetRequest) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotFleetRequest) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

