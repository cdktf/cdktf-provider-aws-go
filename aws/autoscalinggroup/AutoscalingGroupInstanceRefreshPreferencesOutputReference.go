// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/autoscalinggroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingGroupInstanceRefreshPreferencesOutputReference interface {
	cdktf.ComplexObject
	AlarmSpecification() AutoscalingGroupInstanceRefreshPreferencesAlarmSpecificationOutputReference
	AlarmSpecificationInput() *AutoscalingGroupInstanceRefreshPreferencesAlarmSpecification
	AutoRollback() interface{}
	SetAutoRollback(val interface{})
	AutoRollbackInput() interface{}
	CheckpointDelay() *string
	SetCheckpointDelay(val *string)
	CheckpointDelayInput() *string
	CheckpointPercentages() *[]*float64
	SetCheckpointPercentages(val *[]*float64)
	CheckpointPercentagesInput() *[]*float64
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InstanceWarmup() *string
	SetInstanceWarmup(val *string)
	InstanceWarmupInput() *string
	InternalValue() *AutoscalingGroupInstanceRefreshPreferences
	SetInternalValue(val *AutoscalingGroupInstanceRefreshPreferences)
	MaxHealthyPercentage() *float64
	SetMaxHealthyPercentage(val *float64)
	MaxHealthyPercentageInput() *float64
	MinHealthyPercentage() *float64
	SetMinHealthyPercentage(val *float64)
	MinHealthyPercentageInput() *float64
	ScaleInProtectedInstances() *string
	SetScaleInProtectedInstances(val *string)
	ScaleInProtectedInstancesInput() *string
	SkipMatching() interface{}
	SetSkipMatching(val interface{})
	SkipMatchingInput() interface{}
	StandbyInstances() *string
	SetStandbyInstances(val *string)
	StandbyInstancesInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutAlarmSpecification(value *AutoscalingGroupInstanceRefreshPreferencesAlarmSpecification)
	ResetAlarmSpecification()
	ResetAutoRollback()
	ResetCheckpointDelay()
	ResetCheckpointPercentages()
	ResetInstanceWarmup()
	ResetMaxHealthyPercentage()
	ResetMinHealthyPercentage()
	ResetScaleInProtectedInstances()
	ResetSkipMatching()
	ResetStandbyInstances()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingGroupInstanceRefreshPreferencesOutputReference
type jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) AlarmSpecification() AutoscalingGroupInstanceRefreshPreferencesAlarmSpecificationOutputReference {
	var returns AutoscalingGroupInstanceRefreshPreferencesAlarmSpecificationOutputReference
	_jsii_.Get(
		j,
		"alarmSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) AlarmSpecificationInput() *AutoscalingGroupInstanceRefreshPreferencesAlarmSpecification {
	var returns *AutoscalingGroupInstanceRefreshPreferencesAlarmSpecification
	_jsii_.Get(
		j,
		"alarmSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) AutoRollback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRollback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) AutoRollbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRollbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) CheckpointDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkpointDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) CheckpointDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkpointDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) CheckpointPercentages() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"checkpointPercentages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) CheckpointPercentagesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"checkpointPercentagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) InstanceWarmup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceWarmup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) InstanceWarmupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceWarmupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) InternalValue() *AutoscalingGroupInstanceRefreshPreferences {
	var returns *AutoscalingGroupInstanceRefreshPreferences
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) MaxHealthyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHealthyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) MaxHealthyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHealthyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) MinHealthyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minHealthyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) MinHealthyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minHealthyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ScaleInProtectedInstances() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleInProtectedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ScaleInProtectedInstancesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleInProtectedInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) SkipMatching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) SkipMatchingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipMatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) StandbyInstances() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standbyInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) StandbyInstancesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standbyInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutoscalingGroupInstanceRefreshPreferencesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutoscalingGroupInstanceRefreshPreferencesOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingGroupInstanceRefreshPreferencesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroupInstanceRefreshPreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutoscalingGroupInstanceRefreshPreferencesOutputReference_Override(a AutoscalingGroupInstanceRefreshPreferencesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroupInstanceRefreshPreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference)SetAutoRollback(val interface{}) {
	if err := j.validateSetAutoRollbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRollback",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference)SetCheckpointDelay(val *string) {
	if err := j.validateSetCheckpointDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkpointDelay",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference)SetCheckpointPercentages(val *[]*float64) {
	if err := j.validateSetCheckpointPercentagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkpointPercentages",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference)SetInstanceWarmup(val *string) {
	if err := j.validateSetInstanceWarmupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceWarmup",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference)SetInternalValue(val *AutoscalingGroupInstanceRefreshPreferences) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference)SetMaxHealthyPercentage(val *float64) {
	if err := j.validateSetMaxHealthyPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxHealthyPercentage",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference)SetMinHealthyPercentage(val *float64) {
	if err := j.validateSetMinHealthyPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minHealthyPercentage",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference)SetScaleInProtectedInstances(val *string) {
	if err := j.validateSetScaleInProtectedInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleInProtectedInstances",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference)SetSkipMatching(val interface{}) {
	if err := j.validateSetSkipMatchingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipMatching",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference)SetStandbyInstances(val *string) {
	if err := j.validateSetStandbyInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"standbyInstances",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) PutAlarmSpecification(value *AutoscalingGroupInstanceRefreshPreferencesAlarmSpecification) {
	if err := a.validatePutAlarmSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlarmSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ResetAlarmSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetAlarmSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ResetAutoRollback() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoRollback",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ResetCheckpointDelay() {
	_jsii_.InvokeVoid(
		a,
		"resetCheckpointDelay",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ResetCheckpointPercentages() {
	_jsii_.InvokeVoid(
		a,
		"resetCheckpointPercentages",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ResetInstanceWarmup() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceWarmup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ResetMaxHealthyPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxHealthyPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ResetMinHealthyPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetMinHealthyPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ResetScaleInProtectedInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetScaleInProtectedInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ResetSkipMatching() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipMatching",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ResetStandbyInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetStandbyInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInstanceRefreshPreferencesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

