// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsautoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/dataawsautoscalinggroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/data-sources/autoscaling_group aws_autoscaling_group}.
type DataAwsAutoscalingGroup interface {
	cdktf.TerraformDataSource
	Arn() *string
	AvailabilityZones() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultCooldown() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesiredCapacity() *float64
	DesiredCapacityType() *string
	EnabledMetrics() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthCheckGracePeriod() *float64
	HealthCheckType() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceMaintenancePolicy() DataAwsAutoscalingGroupInstanceMaintenancePolicyList
	LaunchConfiguration() *string
	LaunchTemplate() DataAwsAutoscalingGroupLaunchTemplateList
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancers() *[]*string
	MaxInstanceLifetime() *float64
	MaxSize() *float64
	MinSize() *float64
	MixedInstancesPolicy() DataAwsAutoscalingGroupMixedInstancesPolicyList
	Name() *string
	SetName(val *string)
	NameInput() *string
	NewInstancesProtectedFromScaleIn() cdktf.IResolvable
	// The tree node.
	Node() constructs.Node
	PlacementGroup() *string
	PredictedCapacity() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ServiceLinkedRoleArn() *string
	Status() *string
	SuspendedProcesses() *[]*string
	Tag() DataAwsAutoscalingGroupTagList
	TargetGroupArns() *[]*string
	TerminationPolicies() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrafficSource() DataAwsAutoscalingGroupTrafficSourceList
	VpcZoneIdentifier() *string
	WarmPool() DataAwsAutoscalingGroupWarmPoolList
	WarmPoolSize() *float64
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsAutoscalingGroup
type jsiiProxy_DataAwsAutoscalingGroup struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) DefaultCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) DesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) DesiredCapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) EnabledMetrics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) HealthCheckGracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) InstanceMaintenancePolicy() DataAwsAutoscalingGroupInstanceMaintenancePolicyList {
	var returns DataAwsAutoscalingGroupInstanceMaintenancePolicyList
	_jsii_.Get(
		j,
		"instanceMaintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) LaunchConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) LaunchTemplate() DataAwsAutoscalingGroupLaunchTemplateList {
	var returns DataAwsAutoscalingGroupLaunchTemplateList
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) LoadBalancers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) MaxInstanceLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) MixedInstancesPolicy() DataAwsAutoscalingGroupMixedInstancesPolicyList {
	var returns DataAwsAutoscalingGroupMixedInstancesPolicyList
	_jsii_.Get(
		j,
		"mixedInstancesPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) NewInstancesProtectedFromScaleIn() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"newInstancesProtectedFromScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) PredictedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) ServiceLinkedRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLinkedRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) SuspendedProcesses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"suspendedProcesses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) Tag() DataAwsAutoscalingGroupTagList {
	var returns DataAwsAutoscalingGroupTagList
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) TargetGroupArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) TerminationPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"terminationPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) TrafficSource() DataAwsAutoscalingGroupTrafficSourceList {
	var returns DataAwsAutoscalingGroupTrafficSourceList
	_jsii_.Get(
		j,
		"trafficSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) VpcZoneIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcZoneIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) WarmPool() DataAwsAutoscalingGroupWarmPoolList {
	var returns DataAwsAutoscalingGroupWarmPoolList
	_jsii_.Get(
		j,
		"warmPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAutoscalingGroup) WarmPoolSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmPoolSize",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/data-sources/autoscaling_group aws_autoscaling_group} Data Source.
func NewDataAwsAutoscalingGroup(scope constructs.Construct, id *string, config *DataAwsAutoscalingGroupConfig) DataAwsAutoscalingGroup {
	_init_.Initialize()

	if err := validateNewDataAwsAutoscalingGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsAutoscalingGroup{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsAutoscalingGroup.DataAwsAutoscalingGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/data-sources/autoscaling_group aws_autoscaling_group} Data Source.
func NewDataAwsAutoscalingGroup_Override(d DataAwsAutoscalingGroup, scope constructs.Construct, id *string, config *DataAwsAutoscalingGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsAutoscalingGroup.DataAwsAutoscalingGroup",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsAutoscalingGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsAutoscalingGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsAutoscalingGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsAutoscalingGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsAutoscalingGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsAutoscalingGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsAutoscalingGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsAutoscalingGroup resource upon running "cdktf plan <stack-name>".
func DataAwsAutoscalingGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsAutoscalingGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsAutoscalingGroup.DataAwsAutoscalingGroup",
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
func DataAwsAutoscalingGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsAutoscalingGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsAutoscalingGroup.DataAwsAutoscalingGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsAutoscalingGroup_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsAutoscalingGroup_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsAutoscalingGroup.DataAwsAutoscalingGroup",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsAutoscalingGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsAutoscalingGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsAutoscalingGroup.DataAwsAutoscalingGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsAutoscalingGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsAutoscalingGroup.DataAwsAutoscalingGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsAutoscalingGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

