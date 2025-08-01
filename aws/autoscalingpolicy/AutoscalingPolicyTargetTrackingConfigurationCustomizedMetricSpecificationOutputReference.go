// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/autoscalingpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference interface {
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
	InternalValue() *AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification
	SetInternalValue(val *AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification)
	MetricDimension() AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList
	MetricDimensionInput() interface{}
	MetricName() *string
	SetMetricName(val *string)
	MetricNameInput() *string
	Metrics() AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricsList
	MetricsInput() interface{}
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	Period() *float64
	SetPeriod(val *float64)
	PeriodInput() *float64
	Statistic() *string
	SetStatistic(val *string)
	StatisticInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutMetricDimension(value interface{})
	PutMetrics(value interface{})
	ResetMetricDimension()
	ResetMetricName()
	ResetMetrics()
	ResetNamespace()
	ResetPeriod()
	ResetStatistic()
	ResetUnit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference
type jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) InternalValue() *AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification {
	var returns *AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) MetricDimension() AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList {
	var returns AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionList
	_jsii_.Get(
		j,
		"metricDimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) MetricDimensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricDimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) Metrics() AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricsList {
	var returns AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricsList
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) MetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) StatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}


func NewAutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingPolicy.AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference_Override(a AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingPolicy.AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference)SetInternalValue(val *AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference)SetPeriod(val *float64) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) PutMetricDimension(value interface{}) {
	if err := a.validatePutMetricDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetricDimension",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) PutMetrics(value interface{}) {
	if err := a.validatePutMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) ResetMetricDimension() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricDimension",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) ResetMetricName() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) ResetMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) ResetPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) ResetStatistic() {
	_jsii_.InvokeVoid(
		a,
		"resetStatistic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) ResetUnit() {
	_jsii_.InvokeVoid(
		a,
		"resetUnit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

