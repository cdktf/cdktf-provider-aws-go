// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/autoscalinggroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/autoscaling_group aws_autoscaling_group}.
type AutoscalingGroup interface {
	cdktf.TerraformResource
	Arn() *string
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	AvailabilityZonesInput() *[]*string
	CapacityRebalance() interface{}
	SetCapacityRebalance(val interface{})
	CapacityRebalanceInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DefaultCooldown() *float64
	SetDefaultCooldown(val *float64)
	DefaultCooldownInput() *float64
	DefaultInstanceWarmup() *float64
	SetDefaultInstanceWarmup(val *float64)
	DefaultInstanceWarmupInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesiredCapacity() *float64
	SetDesiredCapacity(val *float64)
	DesiredCapacityInput() *float64
	DesiredCapacityType() *string
	SetDesiredCapacityType(val *string)
	DesiredCapacityTypeInput() *string
	EnabledMetrics() *[]*string
	SetEnabledMetrics(val *[]*string)
	EnabledMetricsInput() *[]*string
	ForceDelete() interface{}
	SetForceDelete(val interface{})
	ForceDeleteInput() interface{}
	ForceDeleteWarmPool() interface{}
	SetForceDeleteWarmPool(val interface{})
	ForceDeleteWarmPoolInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthCheckGracePeriod() *float64
	SetHealthCheckGracePeriod(val *float64)
	HealthCheckGracePeriodInput() *float64
	HealthCheckType() *string
	SetHealthCheckType(val *string)
	HealthCheckTypeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoreFailedScalingActivities() interface{}
	SetIgnoreFailedScalingActivities(val interface{})
	IgnoreFailedScalingActivitiesInput() interface{}
	InitialLifecycleHook() AutoscalingGroupInitialLifecycleHookList
	InitialLifecycleHookInput() interface{}
	InstanceMaintenancePolicy() AutoscalingGroupInstanceMaintenancePolicyOutputReference
	InstanceMaintenancePolicyInput() *AutoscalingGroupInstanceMaintenancePolicy
	InstanceRefresh() AutoscalingGroupInstanceRefreshOutputReference
	InstanceRefreshInput() *AutoscalingGroupInstanceRefresh
	LaunchConfiguration() *string
	SetLaunchConfiguration(val *string)
	LaunchConfigurationInput() *string
	LaunchTemplate() AutoscalingGroupLaunchTemplateOutputReference
	LaunchTemplateInput() *AutoscalingGroupLaunchTemplate
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancers() *[]*string
	SetLoadBalancers(val *[]*string)
	LoadBalancersInput() *[]*string
	MaxInstanceLifetime() *float64
	SetMaxInstanceLifetime(val *float64)
	MaxInstanceLifetimeInput() *float64
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	MetricsGranularity() *string
	SetMetricsGranularity(val *string)
	MetricsGranularityInput() *string
	MinElbCapacity() *float64
	SetMinElbCapacity(val *float64)
	MinElbCapacityInput() *float64
	MinSize() *float64
	SetMinSize(val *float64)
	MinSizeInput() *float64
	MixedInstancesPolicy() AutoscalingGroupMixedInstancesPolicyOutputReference
	MixedInstancesPolicyInput() *AutoscalingGroupMixedInstancesPolicy
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	// The tree node.
	Node() constructs.Node
	PlacementGroup() *string
	SetPlacementGroup(val *string)
	PlacementGroupInput() *string
	PredictedCapacity() *float64
	ProtectFromScaleIn() interface{}
	SetProtectFromScaleIn(val interface{})
	ProtectFromScaleInInput() interface{}
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
	ServiceLinkedRoleArn() *string
	SetServiceLinkedRoleArn(val *string)
	ServiceLinkedRoleArnInput() *string
	SuspendedProcesses() *[]*string
	SetSuspendedProcesses(val *[]*string)
	SuspendedProcessesInput() *[]*string
	Tag() AutoscalingGroupTagList
	TagInput() interface{}
	TargetGroupArns() *[]*string
	SetTargetGroupArns(val *[]*string)
	TargetGroupArnsInput() *[]*string
	TerminationPolicies() *[]*string
	SetTerminationPolicies(val *[]*string)
	TerminationPoliciesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AutoscalingGroupTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrafficSource() AutoscalingGroupTrafficSourceList
	TrafficSourceInput() interface{}
	VpcZoneIdentifier() *[]*string
	SetVpcZoneIdentifier(val *[]*string)
	VpcZoneIdentifierInput() *[]*string
	WaitForCapacityTimeout() *string
	SetWaitForCapacityTimeout(val *string)
	WaitForCapacityTimeoutInput() *string
	WaitForElbCapacity() *float64
	SetWaitForElbCapacity(val *float64)
	WaitForElbCapacityInput() *float64
	WarmPool() AutoscalingGroupWarmPoolOutputReference
	WarmPoolInput() *AutoscalingGroupWarmPool
	WarmPoolSize() *float64
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
	PutInitialLifecycleHook(value interface{})
	PutInstanceMaintenancePolicy(value *AutoscalingGroupInstanceMaintenancePolicy)
	PutInstanceRefresh(value *AutoscalingGroupInstanceRefresh)
	PutLaunchTemplate(value *AutoscalingGroupLaunchTemplate)
	PutMixedInstancesPolicy(value *AutoscalingGroupMixedInstancesPolicy)
	PutTag(value interface{})
	PutTimeouts(value *AutoscalingGroupTimeouts)
	PutTrafficSource(value interface{})
	PutWarmPool(value *AutoscalingGroupWarmPool)
	ResetAvailabilityZones()
	ResetCapacityRebalance()
	ResetContext()
	ResetDefaultCooldown()
	ResetDefaultInstanceWarmup()
	ResetDesiredCapacity()
	ResetDesiredCapacityType()
	ResetEnabledMetrics()
	ResetForceDelete()
	ResetForceDeleteWarmPool()
	ResetHealthCheckGracePeriod()
	ResetHealthCheckType()
	ResetId()
	ResetIgnoreFailedScalingActivities()
	ResetInitialLifecycleHook()
	ResetInstanceMaintenancePolicy()
	ResetInstanceRefresh()
	ResetLaunchConfiguration()
	ResetLaunchTemplate()
	ResetLoadBalancers()
	ResetMaxInstanceLifetime()
	ResetMetricsGranularity()
	ResetMinElbCapacity()
	ResetMixedInstancesPolicy()
	ResetName()
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacementGroup()
	ResetProtectFromScaleIn()
	ResetServiceLinkedRoleArn()
	ResetSuspendedProcesses()
	ResetTag()
	ResetTargetGroupArns()
	ResetTerminationPolicies()
	ResetTimeouts()
	ResetTrafficSource()
	ResetVpcZoneIdentifier()
	ResetWaitForCapacityTimeout()
	ResetWaitForElbCapacity()
	ResetWarmPool()
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

// The jsii proxy struct for AutoscalingGroup
type jsiiProxy_AutoscalingGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AutoscalingGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) CapacityRebalance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityRebalance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) CapacityRebalanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityRebalanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) DefaultCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) DefaultCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) DefaultInstanceWarmup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultInstanceWarmup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) DefaultInstanceWarmupInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultInstanceWarmupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) DesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) DesiredCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) DesiredCapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) DesiredCapacityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) EnabledMetrics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) EnabledMetricsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) ForceDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) ForceDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) ForceDeleteWarmPool() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteWarmPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) ForceDeleteWarmPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteWarmPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) HealthCheckGracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) HealthCheckGracePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) HealthCheckTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) IgnoreFailedScalingActivities() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreFailedScalingActivities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) IgnoreFailedScalingActivitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreFailedScalingActivitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) InitialLifecycleHook() AutoscalingGroupInitialLifecycleHookList {
	var returns AutoscalingGroupInitialLifecycleHookList
	_jsii_.Get(
		j,
		"initialLifecycleHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) InitialLifecycleHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialLifecycleHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) InstanceMaintenancePolicy() AutoscalingGroupInstanceMaintenancePolicyOutputReference {
	var returns AutoscalingGroupInstanceMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"instanceMaintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) InstanceMaintenancePolicyInput() *AutoscalingGroupInstanceMaintenancePolicy {
	var returns *AutoscalingGroupInstanceMaintenancePolicy
	_jsii_.Get(
		j,
		"instanceMaintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) InstanceRefresh() AutoscalingGroupInstanceRefreshOutputReference {
	var returns AutoscalingGroupInstanceRefreshOutputReference
	_jsii_.Get(
		j,
		"instanceRefresh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) InstanceRefreshInput() *AutoscalingGroupInstanceRefresh {
	var returns *AutoscalingGroupInstanceRefresh
	_jsii_.Get(
		j,
		"instanceRefreshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) LaunchConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) LaunchConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) LaunchTemplate() AutoscalingGroupLaunchTemplateOutputReference {
	var returns AutoscalingGroupLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) LaunchTemplateInput() *AutoscalingGroupLaunchTemplate {
	var returns *AutoscalingGroupLaunchTemplate
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) LoadBalancers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) LoadBalancersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) MaxInstanceLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) MaxInstanceLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) MetricsGranularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsGranularity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) MetricsGranularityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsGranularityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) MinElbCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minElbCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) MinElbCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minElbCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) MixedInstancesPolicy() AutoscalingGroupMixedInstancesPolicyOutputReference {
	var returns AutoscalingGroupMixedInstancesPolicyOutputReference
	_jsii_.Get(
		j,
		"mixedInstancesPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) MixedInstancesPolicyInput() *AutoscalingGroupMixedInstancesPolicy {
	var returns *AutoscalingGroupMixedInstancesPolicy
	_jsii_.Get(
		j,
		"mixedInstancesPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) PlacementGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) PredictedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) ProtectFromScaleIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectFromScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) ProtectFromScaleInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectFromScaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) ServiceLinkedRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLinkedRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) ServiceLinkedRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLinkedRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) SuspendedProcesses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"suspendedProcesses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) SuspendedProcessesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"suspendedProcessesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) Tag() AutoscalingGroupTagList {
	var returns AutoscalingGroupTagList
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) TagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) TargetGroupArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) TargetGroupArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) TerminationPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"terminationPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) TerminationPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"terminationPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) Timeouts() AutoscalingGroupTimeoutsOutputReference {
	var returns AutoscalingGroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) TrafficSource() AutoscalingGroupTrafficSourceList {
	var returns AutoscalingGroupTrafficSourceList
	_jsii_.Get(
		j,
		"trafficSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) TrafficSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) VpcZoneIdentifier() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcZoneIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) VpcZoneIdentifierInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcZoneIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) WaitForCapacityTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForCapacityTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) WaitForCapacityTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForCapacityTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) WaitForElbCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForElbCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) WaitForElbCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForElbCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) WarmPool() AutoscalingGroupWarmPoolOutputReference {
	var returns AutoscalingGroupWarmPoolOutputReference
	_jsii_.Get(
		j,
		"warmPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) WarmPoolInput() *AutoscalingGroupWarmPool {
	var returns *AutoscalingGroupWarmPool
	_jsii_.Get(
		j,
		"warmPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroup) WarmPoolSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmPoolSize",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/autoscaling_group aws_autoscaling_group} Resource.
func NewAutoscalingGroup(scope constructs.Construct, id *string, config *AutoscalingGroupConfig) AutoscalingGroup {
	_init_.Initialize()

	if err := validateNewAutoscalingGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingGroup{}

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/autoscaling_group aws_autoscaling_group} Resource.
func NewAutoscalingGroup_Override(a AutoscalingGroup, scope constructs.Construct, id *string, config *AutoscalingGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroup",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetCapacityRebalance(val interface{}) {
	if err := j.validateSetCapacityRebalanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityRebalance",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetDefaultCooldown(val *float64) {
	if err := j.validateSetDefaultCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultCooldown",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetDefaultInstanceWarmup(val *float64) {
	if err := j.validateSetDefaultInstanceWarmupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultInstanceWarmup",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetDesiredCapacity(val *float64) {
	if err := j.validateSetDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetDesiredCapacityType(val *string) {
	if err := j.validateSetDesiredCapacityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacityType",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetEnabledMetrics(val *[]*string) {
	if err := j.validateSetEnabledMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledMetrics",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetForceDelete(val interface{}) {
	if err := j.validateSetForceDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDelete",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetForceDeleteWarmPool(val interface{}) {
	if err := j.validateSetForceDeleteWarmPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDeleteWarmPool",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetHealthCheckGracePeriod(val *float64) {
	if err := j.validateSetHealthCheckGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckGracePeriod",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetHealthCheckType(val *string) {
	if err := j.validateSetHealthCheckTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckType",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetIgnoreFailedScalingActivities(val interface{}) {
	if err := j.validateSetIgnoreFailedScalingActivitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreFailedScalingActivities",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetLaunchConfiguration(val *string) {
	if err := j.validateSetLaunchConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchConfiguration",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetLoadBalancers(val *[]*string) {
	if err := j.validateSetLoadBalancersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetMaxInstanceLifetime(val *float64) {
	if err := j.validateSetMaxInstanceLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstanceLifetime",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetMetricsGranularity(val *string) {
	if err := j.validateSetMetricsGranularityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsGranularity",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetMinElbCapacity(val *float64) {
	if err := j.validateSetMinElbCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minElbCapacity",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetPlacementGroup(val *string) {
	if err := j.validateSetPlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroup",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetProtectFromScaleIn(val interface{}) {
	if err := j.validateSetProtectFromScaleInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectFromScaleIn",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetServiceLinkedRoleArn(val *string) {
	if err := j.validateSetServiceLinkedRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceLinkedRoleArn",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetSuspendedProcesses(val *[]*string) {
	if err := j.validateSetSuspendedProcessesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendedProcesses",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetTargetGroupArns(val *[]*string) {
	if err := j.validateSetTargetGroupArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupArns",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetTerminationPolicies(val *[]*string) {
	if err := j.validateSetTerminationPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationPolicies",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetVpcZoneIdentifier(val *[]*string) {
	if err := j.validateSetVpcZoneIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcZoneIdentifier",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetWaitForCapacityTimeout(val *string) {
	if err := j.validateSetWaitForCapacityTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForCapacityTimeout",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroup)SetWaitForElbCapacity(val *float64) {
	if err := j.validateSetWaitForElbCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForElbCapacity",
		val,
	)
}

// Generates CDKTF code for importing a AutoscalingGroup resource upon running "cdktf plan <stack-name>".
func AutoscalingGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAutoscalingGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroup",
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
func AutoscalingGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoscalingGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutoscalingGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoscalingGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutoscalingGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoscalingGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AutoscalingGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AutoscalingGroup) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AutoscalingGroup) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AutoscalingGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AutoscalingGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AutoscalingGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AutoscalingGroup) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AutoscalingGroup) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AutoscalingGroup) PutInitialLifecycleHook(value interface{}) {
	if err := a.validatePutInitialLifecycleHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInitialLifecycleHook",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroup) PutInstanceMaintenancePolicy(value *AutoscalingGroupInstanceMaintenancePolicy) {
	if err := a.validatePutInstanceMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceMaintenancePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroup) PutInstanceRefresh(value *AutoscalingGroupInstanceRefresh) {
	if err := a.validatePutInstanceRefreshParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceRefresh",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroup) PutLaunchTemplate(value *AutoscalingGroupLaunchTemplate) {
	if err := a.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroup) PutMixedInstancesPolicy(value *AutoscalingGroupMixedInstancesPolicy) {
	if err := a.validatePutMixedInstancesPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMixedInstancesPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroup) PutTag(value interface{}) {
	if err := a.validatePutTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTag",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroup) PutTimeouts(value *AutoscalingGroupTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroup) PutTrafficSource(value interface{}) {
	if err := a.validatePutTrafficSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrafficSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroup) PutWarmPool(value *AutoscalingGroupWarmPool) {
	if err := a.validatePutWarmPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWarmPool",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetCapacityRebalance() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityRebalance",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetContext() {
	_jsii_.InvokeVoid(
		a,
		"resetContext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetDefaultCooldown() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultCooldown",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetDefaultInstanceWarmup() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultInstanceWarmup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetDesiredCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetDesiredCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetDesiredCapacityType() {
	_jsii_.InvokeVoid(
		a,
		"resetDesiredCapacityType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetEnabledMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabledMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetForceDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetForceDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetForceDeleteWarmPool() {
	_jsii_.InvokeVoid(
		a,
		"resetForceDeleteWarmPool",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetHealthCheckGracePeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheckGracePeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetHealthCheckType() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheckType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetIgnoreFailedScalingActivities() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreFailedScalingActivities",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetInitialLifecycleHook() {
	_jsii_.InvokeVoid(
		a,
		"resetInitialLifecycleHook",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetInstanceMaintenancePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceMaintenancePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetInstanceRefresh() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceRefresh",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetLaunchConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetLoadBalancers() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetMaxInstanceLifetime() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxInstanceLifetime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetMetricsGranularity() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricsGranularity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetMinElbCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMinElbCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetMixedInstancesPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetMixedInstancesPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetPlacementGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetProtectFromScaleIn() {
	_jsii_.InvokeVoid(
		a,
		"resetProtectFromScaleIn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetServiceLinkedRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceLinkedRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetSuspendedProcesses() {
	_jsii_.InvokeVoid(
		a,
		"resetSuspendedProcesses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetTag() {
	_jsii_.InvokeVoid(
		a,
		"resetTag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetTargetGroupArns() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetGroupArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetTerminationPolicies() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminationPolicies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetTrafficSource() {
	_jsii_.InvokeVoid(
		a,
		"resetTrafficSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetVpcZoneIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcZoneIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetWaitForCapacityTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetWaitForCapacityTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetWaitForElbCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetWaitForElbCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) ResetWarmPool() {
	_jsii_.InvokeVoid(
		a,
		"resetWarmPool",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

