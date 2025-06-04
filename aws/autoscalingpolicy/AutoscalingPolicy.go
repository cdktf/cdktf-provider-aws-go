// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/autoscalingpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/autoscaling_policy aws_autoscaling_policy}.
type AutoscalingPolicy interface {
	cdktf.TerraformResource
	AdjustmentType() *string
	SetAdjustmentType(val *string)
	AdjustmentTypeInput() *string
	Arn() *string
	AutoscalingGroupName() *string
	SetAutoscalingGroupName(val *string)
	AutoscalingGroupNameInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Cooldown() *float64
	SetCooldown(val *float64)
	CooldownInput() *float64
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EstimatedInstanceWarmup() *float64
	SetEstimatedInstanceWarmup(val *float64)
	EstimatedInstanceWarmupInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MetricAggregationType() *string
	SetMetricAggregationType(val *string)
	MetricAggregationTypeInput() *string
	MinAdjustmentMagnitude() *float64
	SetMinAdjustmentMagnitude(val *float64)
	MinAdjustmentMagnitudeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PolicyType() *string
	SetPolicyType(val *string)
	PolicyTypeInput() *string
	PredictiveScalingConfiguration() AutoscalingPolicyPredictiveScalingConfigurationOutputReference
	PredictiveScalingConfigurationInput() *AutoscalingPolicyPredictiveScalingConfiguration
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
	ScalingAdjustment() *float64
	SetScalingAdjustment(val *float64)
	ScalingAdjustmentInput() *float64
	StepAdjustment() AutoscalingPolicyStepAdjustmentList
	StepAdjustmentInput() interface{}
	TargetTrackingConfiguration() AutoscalingPolicyTargetTrackingConfigurationOutputReference
	TargetTrackingConfigurationInput() *AutoscalingPolicyTargetTrackingConfiguration
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutPredictiveScalingConfiguration(value *AutoscalingPolicyPredictiveScalingConfiguration)
	PutStepAdjustment(value interface{})
	PutTargetTrackingConfiguration(value *AutoscalingPolicyTargetTrackingConfiguration)
	ResetAdjustmentType()
	ResetCooldown()
	ResetEnabled()
	ResetEstimatedInstanceWarmup()
	ResetId()
	ResetMetricAggregationType()
	ResetMinAdjustmentMagnitude()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyType()
	ResetPredictiveScalingConfiguration()
	ResetScalingAdjustment()
	ResetStepAdjustment()
	ResetTargetTrackingConfiguration()
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

// The jsii proxy struct for AutoscalingPolicy
type jsiiProxy_AutoscalingPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AutoscalingPolicy) AdjustmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) AdjustmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) AutoscalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) AutoscalingGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) Cooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) CooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) EstimatedInstanceWarmup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"estimatedInstanceWarmup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) EstimatedInstanceWarmupInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"estimatedInstanceWarmupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) MetricAggregationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricAggregationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) MetricAggregationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricAggregationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) MinAdjustmentMagnitude() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAdjustmentMagnitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) MinAdjustmentMagnitudeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAdjustmentMagnitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) PolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) PredictiveScalingConfiguration() AutoscalingPolicyPredictiveScalingConfigurationOutputReference {
	var returns AutoscalingPolicyPredictiveScalingConfigurationOutputReference
	_jsii_.Get(
		j,
		"predictiveScalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) PredictiveScalingConfigurationInput() *AutoscalingPolicyPredictiveScalingConfiguration {
	var returns *AutoscalingPolicyPredictiveScalingConfiguration
	_jsii_.Get(
		j,
		"predictiveScalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) ScalingAdjustment() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scalingAdjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) ScalingAdjustmentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scalingAdjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) StepAdjustment() AutoscalingPolicyStepAdjustmentList {
	var returns AutoscalingPolicyStepAdjustmentList
	_jsii_.Get(
		j,
		"stepAdjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) StepAdjustmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepAdjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) TargetTrackingConfiguration() AutoscalingPolicyTargetTrackingConfigurationOutputReference {
	var returns AutoscalingPolicyTargetTrackingConfigurationOutputReference
	_jsii_.Get(
		j,
		"targetTrackingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) TargetTrackingConfigurationInput() *AutoscalingPolicyTargetTrackingConfiguration {
	var returns *AutoscalingPolicyTargetTrackingConfiguration
	_jsii_.Get(
		j,
		"targetTrackingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/autoscaling_policy aws_autoscaling_policy} Resource.
func NewAutoscalingPolicy(scope constructs.Construct, id *string, config *AutoscalingPolicyConfig) AutoscalingPolicy {
	_init_.Initialize()

	if err := validateNewAutoscalingPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingPolicy{}

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingPolicy.AutoscalingPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/autoscaling_policy aws_autoscaling_policy} Resource.
func NewAutoscalingPolicy_Override(a AutoscalingPolicy, scope constructs.Construct, id *string, config *AutoscalingPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingPolicy.AutoscalingPolicy",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetAdjustmentType(val *string) {
	if err := j.validateSetAdjustmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adjustmentType",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetAutoscalingGroupName(val *string) {
	if err := j.validateSetAutoscalingGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingGroupName",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetCooldown(val *float64) {
	if err := j.validateSetCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cooldown",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetEstimatedInstanceWarmup(val *float64) {
	if err := j.validateSetEstimatedInstanceWarmupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"estimatedInstanceWarmup",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetMetricAggregationType(val *string) {
	if err := j.validateSetMetricAggregationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricAggregationType",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetMinAdjustmentMagnitude(val *float64) {
	if err := j.validateSetMinAdjustmentMagnitudeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minAdjustmentMagnitude",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetPolicyType(val *string) {
	if err := j.validateSetPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicy)SetScalingAdjustment(val *float64) {
	if err := j.validateSetScalingAdjustmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingAdjustment",
		val,
	)
}

// Generates CDKTF code for importing a AutoscalingPolicy resource upon running "cdktf plan <stack-name>".
func AutoscalingPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAutoscalingPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.autoscalingPolicy.AutoscalingPolicy",
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
func AutoscalingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoscalingPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.autoscalingPolicy.AutoscalingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutoscalingPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoscalingPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.autoscalingPolicy.AutoscalingPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutoscalingPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoscalingPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.autoscalingPolicy.AutoscalingPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AutoscalingPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.autoscalingPolicy.AutoscalingPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AutoscalingPolicy) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AutoscalingPolicy) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AutoscalingPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutoscalingPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutoscalingPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutoscalingPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutoscalingPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutoscalingPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutoscalingPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutoscalingPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutoscalingPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AutoscalingPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingPolicy) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AutoscalingPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AutoscalingPolicy) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AutoscalingPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AutoscalingPolicy) PutPredictiveScalingConfiguration(value *AutoscalingPolicyPredictiveScalingConfiguration) {
	if err := a.validatePutPredictiveScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredictiveScalingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingPolicy) PutStepAdjustment(value interface{}) {
	if err := a.validatePutStepAdjustmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStepAdjustment",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingPolicy) PutTargetTrackingConfiguration(value *AutoscalingPolicyTargetTrackingConfiguration) {
	if err := a.validatePutTargetTrackingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargetTrackingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingPolicy) ResetAdjustmentType() {
	_jsii_.InvokeVoid(
		a,
		"resetAdjustmentType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicy) ResetCooldown() {
	_jsii_.InvokeVoid(
		a,
		"resetCooldown",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicy) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicy) ResetEstimatedInstanceWarmup() {
	_jsii_.InvokeVoid(
		a,
		"resetEstimatedInstanceWarmup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicy) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicy) ResetMetricAggregationType() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricAggregationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicy) ResetMinAdjustmentMagnitude() {
	_jsii_.InvokeVoid(
		a,
		"resetMinAdjustmentMagnitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicy) ResetPolicyType() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicy) ResetPredictiveScalingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetPredictiveScalingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicy) ResetScalingAdjustment() {
	_jsii_.InvokeVoid(
		a,
		"resetScalingAdjustment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicy) ResetStepAdjustment() {
	_jsii_.InvokeVoid(
		a,
		"resetStepAdjustment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicy) ResetTargetTrackingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetTrackingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

