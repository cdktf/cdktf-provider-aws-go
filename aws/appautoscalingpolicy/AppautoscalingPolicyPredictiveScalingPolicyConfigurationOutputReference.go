// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appautoscalingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appautoscalingpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *AppautoscalingPolicyPredictiveScalingPolicyConfiguration
	SetInternalValue(val *AppautoscalingPolicyPredictiveScalingPolicyConfiguration)
	MaxCapacityBreachBehavior() *string
	SetMaxCapacityBreachBehavior(val *string)
	MaxCapacityBreachBehaviorInput() *string
	MaxCapacityBuffer() *float64
	SetMaxCapacityBuffer(val *float64)
	MaxCapacityBufferInput() *float64
	MetricSpecification() AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationList
	MetricSpecificationInput() interface{}
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	SchedulingBufferTime() *float64
	SetSchedulingBufferTime(val *float64)
	SchedulingBufferTimeInput() *float64
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
	PutMetricSpecification(value interface{})
	ResetMaxCapacityBreachBehavior()
	ResetMaxCapacityBuffer()
	ResetMode()
	ResetSchedulingBufferTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference
type jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) InternalValue() *AppautoscalingPolicyPredictiveScalingPolicyConfiguration {
	var returns *AppautoscalingPolicyPredictiveScalingPolicyConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) MaxCapacityBreachBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCapacityBreachBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) MaxCapacityBreachBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCapacityBreachBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) MaxCapacityBuffer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) MaxCapacityBufferInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityBufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) MetricSpecification() AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationList {
	var returns AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationList
	_jsii_.Get(
		j,
		"metricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) MetricSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) SchedulingBufferTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulingBufferTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) SchedulingBufferTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulingBufferTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewAppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appautoscalingPolicy.AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference_Override(a AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appautoscalingPolicy.AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference)SetInternalValue(val *AppautoscalingPolicyPredictiveScalingPolicyConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference)SetMaxCapacityBreachBehavior(val *string) {
	if err := j.validateSetMaxCapacityBreachBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacityBreachBehavior",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference)SetMaxCapacityBuffer(val *float64) {
	if err := j.validateSetMaxCapacityBufferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacityBuffer",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference)SetSchedulingBufferTime(val *float64) {
	if err := j.validateSetSchedulingBufferTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulingBufferTime",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) PutMetricSpecification(value interface{}) {
	if err := a.validatePutMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) ResetMaxCapacityBreachBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxCapacityBreachBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) ResetMaxCapacityBuffer() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxCapacityBuffer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		a,
		"resetMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) ResetSchedulingBufferTime() {
	_jsii_.InvokeVoid(
		a,
		"resetSchedulingBufferTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppautoscalingPolicyPredictiveScalingPolicyConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

