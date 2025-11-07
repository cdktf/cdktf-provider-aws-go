// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/lexv2modelsintent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference interface {
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
	FailureConditional() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalList
	FailureConditionalInput() interface{}
	FailureNextStep() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureNextStepList
	FailureNextStepInput() interface{}
	FailureResponse() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureResponseList
	FailureResponseInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SuccessConditional() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessConditionalList
	SuccessConditionalInput() interface{}
	SuccessNextStep() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepList
	SuccessNextStepInput() interface{}
	SuccessResponse() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessResponseList
	SuccessResponseInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutConditional() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalList
	TimeoutConditionalInput() interface{}
	TimeoutNextStep() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutNextStepList
	TimeoutNextStepInput() interface{}
	TimeoutResponse() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutResponseList
	TimeoutResponseInput() interface{}
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
	PutFailureConditional(value interface{})
	PutFailureNextStep(value interface{})
	PutFailureResponse(value interface{})
	PutSuccessConditional(value interface{})
	PutSuccessNextStep(value interface{})
	PutSuccessResponse(value interface{})
	PutTimeoutConditional(value interface{})
	PutTimeoutNextStep(value interface{})
	PutTimeoutResponse(value interface{})
	ResetFailureConditional()
	ResetFailureNextStep()
	ResetFailureResponse()
	ResetSuccessConditional()
	ResetSuccessNextStep()
	ResetSuccessResponse()
	ResetTimeoutConditional()
	ResetTimeoutNextStep()
	ResetTimeoutResponse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference
type jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) FailureConditional() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalList {
	var returns Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalList
	_jsii_.Get(
		j,
		"failureConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) FailureConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) FailureNextStep() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureNextStepList {
	var returns Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureNextStepList
	_jsii_.Get(
		j,
		"failureNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) FailureNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) FailureResponse() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureResponseList {
	var returns Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureResponseList
	_jsii_.Get(
		j,
		"failureResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) FailureResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) SuccessConditional() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessConditionalList {
	var returns Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessConditionalList
	_jsii_.Get(
		j,
		"successConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) SuccessConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) SuccessNextStep() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepList {
	var returns Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepList
	_jsii_.Get(
		j,
		"successNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) SuccessNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) SuccessResponse() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessResponseList {
	var returns Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessResponseList
	_jsii_.Get(
		j,
		"successResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) SuccessResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) TimeoutConditional() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalList {
	var returns Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalList
	_jsii_.Get(
		j,
		"timeoutConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) TimeoutConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) TimeoutNextStep() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutNextStepList {
	var returns Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutNextStepList
	_jsii_.Get(
		j,
		"timeoutNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) TimeoutNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) TimeoutResponse() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutResponseList {
	var returns Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutResponseList
	_jsii_.Get(
		j,
		"timeoutResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) TimeoutResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutResponseInput",
		&returns,
	)
	return returns
}


func NewLexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewLexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference_Override(l Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) PutFailureConditional(value interface{}) {
	if err := l.validatePutFailureConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFailureConditional",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) PutFailureNextStep(value interface{}) {
	if err := l.validatePutFailureNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFailureNextStep",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) PutFailureResponse(value interface{}) {
	if err := l.validatePutFailureResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFailureResponse",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) PutSuccessConditional(value interface{}) {
	if err := l.validatePutSuccessConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSuccessConditional",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) PutSuccessNextStep(value interface{}) {
	if err := l.validatePutSuccessNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSuccessNextStep",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) PutSuccessResponse(value interface{}) {
	if err := l.validatePutSuccessResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSuccessResponse",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) PutTimeoutConditional(value interface{}) {
	if err := l.validatePutTimeoutConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeoutConditional",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) PutTimeoutNextStep(value interface{}) {
	if err := l.validatePutTimeoutNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeoutNextStep",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) PutTimeoutResponse(value interface{}) {
	if err := l.validatePutTimeoutResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeoutResponse",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) ResetFailureConditional() {
	_jsii_.InvokeVoid(
		l,
		"resetFailureConditional",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) ResetFailureNextStep() {
	_jsii_.InvokeVoid(
		l,
		"resetFailureNextStep",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) ResetFailureResponse() {
	_jsii_.InvokeVoid(
		l,
		"resetFailureResponse",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) ResetSuccessConditional() {
	_jsii_.InvokeVoid(
		l,
		"resetSuccessConditional",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) ResetSuccessNextStep() {
	_jsii_.InvokeVoid(
		l,
		"resetSuccessNextStep",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) ResetSuccessResponse() {
	_jsii_.InvokeVoid(
		l,
		"resetSuccessResponse",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) ResetTimeoutConditional() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeoutConditional",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) ResetTimeoutNextStep() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeoutNextStep",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) ResetTimeoutResponse() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeoutResponse",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

