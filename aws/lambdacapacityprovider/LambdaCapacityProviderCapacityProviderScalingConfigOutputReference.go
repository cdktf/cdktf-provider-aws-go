// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdacapacityprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/lambdacapacityprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LambdaCapacityProviderCapacityProviderScalingConfigOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxVcpuCount() *float64
	SetMaxVcpuCount(val *float64)
	MaxVcpuCountInput() *float64
	ScalingMode() *string
	SetScalingMode(val *string)
	ScalingModeInput() *string
	ScalingPolicies() LambdaCapacityProviderCapacityProviderScalingConfigScalingPoliciesList
	ScalingPoliciesInput() interface{}
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
	PutScalingPolicies(value interface{})
	ResetMaxVcpuCount()
	ResetScalingMode()
	ResetScalingPolicies()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LambdaCapacityProviderCapacityProviderScalingConfigOutputReference
type jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) MaxVcpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVcpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) MaxVcpuCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVcpuCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ScalingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ScalingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ScalingPolicies() LambdaCapacityProviderCapacityProviderScalingConfigScalingPoliciesList {
	var returns LambdaCapacityProviderCapacityProviderScalingConfigScalingPoliciesList
	_jsii_.Get(
		j,
		"scalingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ScalingPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLambdaCapacityProviderCapacityProviderScalingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LambdaCapacityProviderCapacityProviderScalingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLambdaCapacityProviderCapacityProviderScalingConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lambdaCapacityProvider.LambdaCapacityProviderCapacityProviderScalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLambdaCapacityProviderCapacityProviderScalingConfigOutputReference_Override(l LambdaCapacityProviderCapacityProviderScalingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lambdaCapacityProvider.LambdaCapacityProviderCapacityProviderScalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference)SetMaxVcpuCount(val *float64) {
	if err := j.validateSetMaxVcpuCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxVcpuCount",
		val,
	)
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference)SetScalingMode(val *string) {
	if err := j.validateSetScalingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingMode",
		val,
	)
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) PutScalingPolicies(value interface{}) {
	if err := l.validatePutScalingPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putScalingPolicies",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ResetMaxVcpuCount() {
	_jsii_.InvokeVoid(
		l,
		"resetMaxVcpuCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ResetScalingMode() {
	_jsii_.InvokeVoid(
		l,
		"resetScalingMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ResetScalingPolicies() {
	_jsii_.InvokeVoid(
		l,
		"resetScalingPolicies",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaCapacityProviderCapacityProviderScalingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

