// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/lexv2modelsintent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsIntentConfirmationSettingOutputReference interface {
	cdktf.ComplexObject
	Active() interface{}
	SetActive(val interface{})
	ActiveInput() interface{}
	CodeHook() Lexv2ModelsIntentConfirmationSettingCodeHookList
	CodeHookInput() interface{}
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
	ConfirmationConditional() Lexv2ModelsIntentConfirmationSettingConfirmationConditionalList
	ConfirmationConditionalInput() interface{}
	ConfirmationNextStep() Lexv2ModelsIntentConfirmationSettingConfirmationNextStepList
	ConfirmationNextStepInput() interface{}
	ConfirmationResponse() Lexv2ModelsIntentConfirmationSettingConfirmationResponseList
	ConfirmationResponseInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeclinationConditional() Lexv2ModelsIntentConfirmationSettingDeclinationConditionalList
	DeclinationConditionalInput() interface{}
	DeclinationNextStep() Lexv2ModelsIntentConfirmationSettingDeclinationNextStepList
	DeclinationNextStepInput() interface{}
	DeclinationResponse() Lexv2ModelsIntentConfirmationSettingDeclinationResponseList
	DeclinationResponseInput() interface{}
	ElicitationCodeHook() Lexv2ModelsIntentConfirmationSettingElicitationCodeHookList
	ElicitationCodeHookInput() interface{}
	FailureConditional() Lexv2ModelsIntentConfirmationSettingFailureConditionalList
	FailureConditionalInput() interface{}
	FailureNextStep() Lexv2ModelsIntentConfirmationSettingFailureNextStepList
	FailureNextStepInput() interface{}
	FailureResponse() Lexv2ModelsIntentConfirmationSettingFailureResponseList
	FailureResponseInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PromptSpecification() Lexv2ModelsIntentConfirmationSettingPromptSpecificationList
	PromptSpecificationInput() interface{}
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
	PutCodeHook(value interface{})
	PutConfirmationConditional(value interface{})
	PutConfirmationNextStep(value interface{})
	PutConfirmationResponse(value interface{})
	PutDeclinationConditional(value interface{})
	PutDeclinationNextStep(value interface{})
	PutDeclinationResponse(value interface{})
	PutElicitationCodeHook(value interface{})
	PutFailureConditional(value interface{})
	PutFailureNextStep(value interface{})
	PutFailureResponse(value interface{})
	PutPromptSpecification(value interface{})
	ResetActive()
	ResetCodeHook()
	ResetConfirmationConditional()
	ResetConfirmationNextStep()
	ResetConfirmationResponse()
	ResetDeclinationConditional()
	ResetDeclinationNextStep()
	ResetDeclinationResponse()
	ResetElicitationCodeHook()
	ResetFailureConditional()
	ResetFailureNextStep()
	ResetFailureResponse()
	ResetPromptSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Lexv2ModelsIntentConfirmationSettingOutputReference
type jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) Active() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) CodeHook() Lexv2ModelsIntentConfirmationSettingCodeHookList {
	var returns Lexv2ModelsIntentConfirmationSettingCodeHookList
	_jsii_.Get(
		j,
		"codeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) CodeHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ConfirmationConditional() Lexv2ModelsIntentConfirmationSettingConfirmationConditionalList {
	var returns Lexv2ModelsIntentConfirmationSettingConfirmationConditionalList
	_jsii_.Get(
		j,
		"confirmationConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ConfirmationConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confirmationConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ConfirmationNextStep() Lexv2ModelsIntentConfirmationSettingConfirmationNextStepList {
	var returns Lexv2ModelsIntentConfirmationSettingConfirmationNextStepList
	_jsii_.Get(
		j,
		"confirmationNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ConfirmationNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confirmationNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ConfirmationResponse() Lexv2ModelsIntentConfirmationSettingConfirmationResponseList {
	var returns Lexv2ModelsIntentConfirmationSettingConfirmationResponseList
	_jsii_.Get(
		j,
		"confirmationResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ConfirmationResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confirmationResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) DeclinationConditional() Lexv2ModelsIntentConfirmationSettingDeclinationConditionalList {
	var returns Lexv2ModelsIntentConfirmationSettingDeclinationConditionalList
	_jsii_.Get(
		j,
		"declinationConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) DeclinationConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"declinationConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) DeclinationNextStep() Lexv2ModelsIntentConfirmationSettingDeclinationNextStepList {
	var returns Lexv2ModelsIntentConfirmationSettingDeclinationNextStepList
	_jsii_.Get(
		j,
		"declinationNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) DeclinationNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"declinationNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) DeclinationResponse() Lexv2ModelsIntentConfirmationSettingDeclinationResponseList {
	var returns Lexv2ModelsIntentConfirmationSettingDeclinationResponseList
	_jsii_.Get(
		j,
		"declinationResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) DeclinationResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"declinationResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ElicitationCodeHook() Lexv2ModelsIntentConfirmationSettingElicitationCodeHookList {
	var returns Lexv2ModelsIntentConfirmationSettingElicitationCodeHookList
	_jsii_.Get(
		j,
		"elicitationCodeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ElicitationCodeHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elicitationCodeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) FailureConditional() Lexv2ModelsIntentConfirmationSettingFailureConditionalList {
	var returns Lexv2ModelsIntentConfirmationSettingFailureConditionalList
	_jsii_.Get(
		j,
		"failureConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) FailureConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) FailureNextStep() Lexv2ModelsIntentConfirmationSettingFailureNextStepList {
	var returns Lexv2ModelsIntentConfirmationSettingFailureNextStepList
	_jsii_.Get(
		j,
		"failureNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) FailureNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) FailureResponse() Lexv2ModelsIntentConfirmationSettingFailureResponseList {
	var returns Lexv2ModelsIntentConfirmationSettingFailureResponseList
	_jsii_.Get(
		j,
		"failureResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) FailureResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) PromptSpecification() Lexv2ModelsIntentConfirmationSettingPromptSpecificationList {
	var returns Lexv2ModelsIntentConfirmationSettingPromptSpecificationList
	_jsii_.Get(
		j,
		"promptSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) PromptSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLexv2ModelsIntentConfirmationSettingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Lexv2ModelsIntentConfirmationSettingOutputReference {
	_init_.Initialize()

	if err := validateNewLexv2ModelsIntentConfirmationSettingOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntentConfirmationSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLexv2ModelsIntentConfirmationSettingOutputReference_Override(l Lexv2ModelsIntentConfirmationSettingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntentConfirmationSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference)SetActive(val interface{}) {
	if err := j.validateSetActiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"active",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) PutCodeHook(value interface{}) {
	if err := l.validatePutCodeHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCodeHook",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) PutConfirmationConditional(value interface{}) {
	if err := l.validatePutConfirmationConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putConfirmationConditional",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) PutConfirmationNextStep(value interface{}) {
	if err := l.validatePutConfirmationNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putConfirmationNextStep",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) PutConfirmationResponse(value interface{}) {
	if err := l.validatePutConfirmationResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putConfirmationResponse",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) PutDeclinationConditional(value interface{}) {
	if err := l.validatePutDeclinationConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDeclinationConditional",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) PutDeclinationNextStep(value interface{}) {
	if err := l.validatePutDeclinationNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDeclinationNextStep",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) PutDeclinationResponse(value interface{}) {
	if err := l.validatePutDeclinationResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDeclinationResponse",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) PutElicitationCodeHook(value interface{}) {
	if err := l.validatePutElicitationCodeHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putElicitationCodeHook",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) PutFailureConditional(value interface{}) {
	if err := l.validatePutFailureConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFailureConditional",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) PutFailureNextStep(value interface{}) {
	if err := l.validatePutFailureNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFailureNextStep",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) PutFailureResponse(value interface{}) {
	if err := l.validatePutFailureResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFailureResponse",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) PutPromptSpecification(value interface{}) {
	if err := l.validatePutPromptSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPromptSpecification",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ResetActive() {
	_jsii_.InvokeVoid(
		l,
		"resetActive",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ResetCodeHook() {
	_jsii_.InvokeVoid(
		l,
		"resetCodeHook",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ResetConfirmationConditional() {
	_jsii_.InvokeVoid(
		l,
		"resetConfirmationConditional",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ResetConfirmationNextStep() {
	_jsii_.InvokeVoid(
		l,
		"resetConfirmationNextStep",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ResetConfirmationResponse() {
	_jsii_.InvokeVoid(
		l,
		"resetConfirmationResponse",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ResetDeclinationConditional() {
	_jsii_.InvokeVoid(
		l,
		"resetDeclinationConditional",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ResetDeclinationNextStep() {
	_jsii_.InvokeVoid(
		l,
		"resetDeclinationNextStep",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ResetDeclinationResponse() {
	_jsii_.InvokeVoid(
		l,
		"resetDeclinationResponse",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ResetElicitationCodeHook() {
	_jsii_.InvokeVoid(
		l,
		"resetElicitationCodeHook",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ResetFailureConditional() {
	_jsii_.InvokeVoid(
		l,
		"resetFailureConditional",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ResetFailureNextStep() {
	_jsii_.InvokeVoid(
		l,
		"resetFailureNextStep",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ResetFailureResponse() {
	_jsii_.InvokeVoid(
		l,
		"resetFailureResponse",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ResetPromptSpecification() {
	_jsii_.InvokeVoid(
		l,
		"resetPromptSpecification",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

