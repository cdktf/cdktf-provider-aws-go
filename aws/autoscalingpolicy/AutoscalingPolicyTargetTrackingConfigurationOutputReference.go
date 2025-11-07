// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/autoscalingpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingPolicyTargetTrackingConfigurationOutputReference interface {
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
	CustomizedMetricSpecification() AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference
	CustomizedMetricSpecificationInput() *AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification
	DisableScaleIn() interface{}
	SetDisableScaleIn(val interface{})
	DisableScaleInInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *AutoscalingPolicyTargetTrackingConfiguration
	SetInternalValue(val *AutoscalingPolicyTargetTrackingConfiguration)
	PredefinedMetricSpecification() AutoscalingPolicyTargetTrackingConfigurationPredefinedMetricSpecificationOutputReference
	PredefinedMetricSpecificationInput() *AutoscalingPolicyTargetTrackingConfigurationPredefinedMetricSpecification
	TargetValue() *float64
	SetTargetValue(val *float64)
	TargetValueInput() *float64
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
	PutCustomizedMetricSpecification(value *AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification)
	PutPredefinedMetricSpecification(value *AutoscalingPolicyTargetTrackingConfigurationPredefinedMetricSpecification)
	ResetCustomizedMetricSpecification()
	ResetDisableScaleIn()
	ResetPredefinedMetricSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingPolicyTargetTrackingConfigurationOutputReference
type jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) CustomizedMetricSpecification() AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference {
	var returns AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"customizedMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) CustomizedMetricSpecificationInput() *AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification {
	var returns *AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification
	_jsii_.Get(
		j,
		"customizedMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) DisableScaleIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) DisableScaleInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableScaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) InternalValue() *AutoscalingPolicyTargetTrackingConfiguration {
	var returns *AutoscalingPolicyTargetTrackingConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) PredefinedMetricSpecification() AutoscalingPolicyTargetTrackingConfigurationPredefinedMetricSpecificationOutputReference {
	var returns AutoscalingPolicyTargetTrackingConfigurationPredefinedMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"predefinedMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) PredefinedMetricSpecificationInput() *AutoscalingPolicyTargetTrackingConfigurationPredefinedMetricSpecification {
	var returns *AutoscalingPolicyTargetTrackingConfigurationPredefinedMetricSpecification
	_jsii_.Get(
		j,
		"predefinedMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) TargetValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) TargetValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutoscalingPolicyTargetTrackingConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutoscalingPolicyTargetTrackingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingPolicyTargetTrackingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingPolicy.AutoscalingPolicyTargetTrackingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutoscalingPolicyTargetTrackingConfigurationOutputReference_Override(a AutoscalingPolicyTargetTrackingConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingPolicy.AutoscalingPolicyTargetTrackingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference)SetDisableScaleIn(val interface{}) {
	if err := j.validateSetDisableScaleInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableScaleIn",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference)SetInternalValue(val *AutoscalingPolicyTargetTrackingConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference)SetTargetValue(val *float64) {
	if err := j.validateSetTargetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) PutCustomizedMetricSpecification(value *AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification) {
	if err := a.validatePutCustomizedMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) PutPredefinedMetricSpecification(value *AutoscalingPolicyTargetTrackingConfigurationPredefinedMetricSpecification) {
	if err := a.validatePutPredefinedMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) ResetCustomizedMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) ResetDisableScaleIn() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableScaleIn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) ResetPredefinedMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutoscalingPolicyTargetTrackingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

