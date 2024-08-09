// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/lexv2modelsslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference interface {
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
	DefaultValueSpecification() Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecificationList
	DefaultValueSpecificationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PromptSpecification() Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationList
	PromptSpecificationInput() interface{}
	SampleUtterance() Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingSampleUtteranceList
	SampleUtteranceInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WaitAndContinueSpecification() Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecificationList
	WaitAndContinueSpecificationInput() interface{}
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
	PutDefaultValueSpecification(value interface{})
	PutPromptSpecification(value interface{})
	PutSampleUtterance(value interface{})
	PutWaitAndContinueSpecification(value interface{})
	ResetDefaultValueSpecification()
	ResetPromptSpecification()
	ResetSampleUtterance()
	ResetWaitAndContinueSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference
type jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) DefaultValueSpecification() Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecificationList {
	var returns Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecificationList
	_jsii_.Get(
		j,
		"defaultValueSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) DefaultValueSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultValueSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) PromptSpecification() Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationList {
	var returns Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationList
	_jsii_.Get(
		j,
		"promptSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) PromptSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) SampleUtterance() Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingSampleUtteranceList {
	var returns Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingSampleUtteranceList
	_jsii_.Get(
		j,
		"sampleUtterance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) SampleUtteranceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sampleUtteranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) WaitAndContinueSpecification() Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecificationList {
	var returns Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecificationList
	_jsii_.Get(
		j,
		"waitAndContinueSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) WaitAndContinueSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitAndContinueSpecificationInput",
		&returns,
	)
	return returns
}


func NewLexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference {
	_init_.Initialize()

	if err := validateNewLexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsSlot.Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference_Override(l Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsSlot.Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) PutDefaultValueSpecification(value interface{}) {
	if err := l.validatePutDefaultValueSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDefaultValueSpecification",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) PutPromptSpecification(value interface{}) {
	if err := l.validatePutPromptSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPromptSpecification",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) PutSampleUtterance(value interface{}) {
	if err := l.validatePutSampleUtteranceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSampleUtterance",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) PutWaitAndContinueSpecification(value interface{}) {
	if err := l.validatePutWaitAndContinueSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putWaitAndContinueSpecification",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) ResetDefaultValueSpecification() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultValueSpecification",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) ResetPromptSpecification() {
	_jsii_.InvokeVoid(
		l,
		"resetPromptSpecification",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) ResetSampleUtterance() {
	_jsii_.InvokeVoid(
		l,
		"resetSampleUtterance",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) ResetWaitAndContinueSpecification() {
	_jsii_.InvokeVoid(
		l,
		"resetWaitAndContinueSpecification",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

