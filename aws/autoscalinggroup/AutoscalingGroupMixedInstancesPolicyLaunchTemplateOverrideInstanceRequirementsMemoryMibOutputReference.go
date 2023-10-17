// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/autoscalinggroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference interface {
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
	InternalValue() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMib
	SetInternalValue(val *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMib)
	Max() *float64
	SetMax(val *float64)
	MaxInput() *float64
	Min() *float64
	SetMin(val *float64)
	MinInput() *float64
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetMax()
	ResetMin()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference
type jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) InternalValue() *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMib {
	var returns *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMib
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) Max() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"max",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) MaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) Min() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"min",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) MinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference_Override(a AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference)SetInternalValue(val *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMib) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference)SetMax(val *float64) {
	if err := j.validateSetMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"max",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference)SetMin(val *float64) {
	if err := j.validateSetMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"min",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) ResetMax() {
	_jsii_.InvokeVoid(
		a,
		"resetMax",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) ResetMin() {
	_jsii_.InvokeVoid(
		a,
		"resetMin",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMibOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

