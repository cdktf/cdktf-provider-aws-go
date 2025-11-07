// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/lexv2modelsslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsSlotValueElicitationSettingOutputReference interface {
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
	DefaultValueSpecification() Lexv2ModelsSlotValueElicitationSettingDefaultValueSpecificationList
	DefaultValueSpecificationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PromptSpecification() Lexv2ModelsSlotValueElicitationSettingPromptSpecificationList
	PromptSpecificationInput() interface{}
	SampleUtterance() Lexv2ModelsSlotValueElicitationSettingSampleUtteranceList
	SampleUtteranceInput() interface{}
	SlotConstraint() *string
	SetSlotConstraint(val *string)
	SlotConstraintInput() *string
	SlotResolutionSetting() Lexv2ModelsSlotValueElicitationSettingSlotResolutionSettingList
	SlotResolutionSettingInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WaitAndContinueSpecification() Lexv2ModelsSlotValueElicitationSettingWaitAndContinueSpecificationList
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutDefaultValueSpecification(value interface{})
	PutPromptSpecification(value interface{})
	PutSampleUtterance(value interface{})
	PutSlotResolutionSetting(value interface{})
	PutWaitAndContinueSpecification(value interface{})
	ResetDefaultValueSpecification()
	ResetPromptSpecification()
	ResetSampleUtterance()
	ResetSlotResolutionSetting()
	ResetWaitAndContinueSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Lexv2ModelsSlotValueElicitationSettingOutputReference
type jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) DefaultValueSpecification() Lexv2ModelsSlotValueElicitationSettingDefaultValueSpecificationList {
	var returns Lexv2ModelsSlotValueElicitationSettingDefaultValueSpecificationList
	_jsii_.Get(
		j,
		"defaultValueSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) DefaultValueSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultValueSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) PromptSpecification() Lexv2ModelsSlotValueElicitationSettingPromptSpecificationList {
	var returns Lexv2ModelsSlotValueElicitationSettingPromptSpecificationList
	_jsii_.Get(
		j,
		"promptSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) PromptSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) SampleUtterance() Lexv2ModelsSlotValueElicitationSettingSampleUtteranceList {
	var returns Lexv2ModelsSlotValueElicitationSettingSampleUtteranceList
	_jsii_.Get(
		j,
		"sampleUtterance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) SampleUtteranceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sampleUtteranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) SlotConstraint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) SlotConstraintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) SlotResolutionSetting() Lexv2ModelsSlotValueElicitationSettingSlotResolutionSettingList {
	var returns Lexv2ModelsSlotValueElicitationSettingSlotResolutionSettingList
	_jsii_.Get(
		j,
		"slotResolutionSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) SlotResolutionSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slotResolutionSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) WaitAndContinueSpecification() Lexv2ModelsSlotValueElicitationSettingWaitAndContinueSpecificationList {
	var returns Lexv2ModelsSlotValueElicitationSettingWaitAndContinueSpecificationList
	_jsii_.Get(
		j,
		"waitAndContinueSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) WaitAndContinueSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitAndContinueSpecificationInput",
		&returns,
	)
	return returns
}


func NewLexv2ModelsSlotValueElicitationSettingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Lexv2ModelsSlotValueElicitationSettingOutputReference {
	_init_.Initialize()

	if err := validateNewLexv2ModelsSlotValueElicitationSettingOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsSlot.Lexv2ModelsSlotValueElicitationSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLexv2ModelsSlotValueElicitationSettingOutputReference_Override(l Lexv2ModelsSlotValueElicitationSettingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsSlot.Lexv2ModelsSlotValueElicitationSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference)SetSlotConstraint(val *string) {
	if err := j.validateSetSlotConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slotConstraint",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) PutDefaultValueSpecification(value interface{}) {
	if err := l.validatePutDefaultValueSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDefaultValueSpecification",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) PutPromptSpecification(value interface{}) {
	if err := l.validatePutPromptSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPromptSpecification",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) PutSampleUtterance(value interface{}) {
	if err := l.validatePutSampleUtteranceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSampleUtterance",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) PutSlotResolutionSetting(value interface{}) {
	if err := l.validatePutSlotResolutionSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSlotResolutionSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) PutWaitAndContinueSpecification(value interface{}) {
	if err := l.validatePutWaitAndContinueSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putWaitAndContinueSpecification",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) ResetDefaultValueSpecification() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultValueSpecification",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) ResetPromptSpecification() {
	_jsii_.InvokeVoid(
		l,
		"resetPromptSpecification",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) ResetSampleUtterance() {
	_jsii_.InvokeVoid(
		l,
		"resetSampleUtterance",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) ResetSlotResolutionSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetSlotResolutionSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) ResetWaitAndContinueSpecification() {
	_jsii_.InvokeVoid(
		l,
		"resetWaitAndContinueSpecification",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_Lexv2ModelsSlotValueElicitationSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

