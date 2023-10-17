// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/pipespipe/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference interface {
	cdktf.ComplexObject
	Base() *float64
	SetBase(val *float64)
	BaseInput() *float64
	CapacityProvider() *string
	SetCapacityProvider(val *string)
	CapacityProviderInput() *string
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
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Weight() *float64
	SetWeight(val *float64)
	WeightInput() *float64
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
	ResetBase()
	ResetWeight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference
type jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) Base() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"base",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) BaseInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) CapacityProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) CapacityProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


func NewPipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference {
	_init_.Initialize()

	if err := validateNewPipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference_Override(p PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference)SetBase(val *float64) {
	if err := j.validateSetBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"base",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference)SetCapacityProvider(val *string) {
	if err := j.validateSetCapacityProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityProvider",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference)SetWeight(val *float64) {
	if err := j.validateSetWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) ResetBase() {
	_jsii_.InvokeVoid(
		p,
		"resetBase",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		p,
		"resetWeight",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersCapacityProviderStrategyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

