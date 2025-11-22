// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/bedrockguardrail/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference interface {
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
	InputAction() *string
	SetInputAction(val *string)
	InputActionInput() *string
	InputEnabled() interface{}
	SetInputEnabled(val interface{})
	InputEnabledInput() interface{}
	InputModalities() *[]*string
	SetInputModalities(val *[]*string)
	InputModalitiesInput() *[]*string
	InputStrength() *string
	SetInputStrength(val *string)
	InputStrengthInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OutputAction() *string
	SetOutputAction(val *string)
	OutputActionInput() *string
	OutputEnabled() interface{}
	SetOutputEnabled(val interface{})
	OutputEnabledInput() interface{}
	OutputModalities() *[]*string
	SetOutputModalities(val *[]*string)
	OutputModalitiesInput() *[]*string
	OutputStrength() *string
	SetOutputStrength(val *string)
	OutputStrengthInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetInputAction()
	ResetInputEnabled()
	ResetInputModalities()
	ResetOutputAction()
	ResetOutputEnabled()
	ResetOutputModalities()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference
type jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) InputAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) InputActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) InputEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) InputEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) InputModalities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputModalities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) InputModalitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputModalitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) InputStrength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputStrength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) InputStrengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputStrengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) OutputAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) OutputActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) OutputEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) OutputEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) OutputModalities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputModalities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) OutputModalitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputModalitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) OutputStrength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputStrength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) OutputStrengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputStrengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewBedrockGuardrailContentPolicyConfigFiltersConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockGuardrailContentPolicyConfigFiltersConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockGuardrail.BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockGuardrailContentPolicyConfigFiltersConfigOutputReference_Override(b BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockGuardrail.BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference)SetInputAction(val *string) {
	if err := j.validateSetInputActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputAction",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference)SetInputEnabled(val interface{}) {
	if err := j.validateSetInputEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputEnabled",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference)SetInputModalities(val *[]*string) {
	if err := j.validateSetInputModalitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputModalities",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference)SetInputStrength(val *string) {
	if err := j.validateSetInputStrengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputStrength",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference)SetOutputAction(val *string) {
	if err := j.validateSetOutputActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputAction",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference)SetOutputEnabled(val interface{}) {
	if err := j.validateSetOutputEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputEnabled",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference)SetOutputModalities(val *[]*string) {
	if err := j.validateSetOutputModalitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputModalities",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference)SetOutputStrength(val *string) {
	if err := j.validateSetOutputStrengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputStrength",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) ResetInputAction() {
	_jsii_.InvokeVoid(
		b,
		"resetInputAction",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) ResetInputEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetInputEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) ResetInputModalities() {
	_jsii_.InvokeVoid(
		b,
		"resetInputModalities",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) ResetOutputAction() {
	_jsii_.InvokeVoid(
		b,
		"resetOutputAction",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) ResetOutputEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetOutputEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) ResetOutputModalities() {
	_jsii_.InvokeVoid(
		b,
		"resetOutputModalities",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrailContentPolicyConfigFiltersConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

