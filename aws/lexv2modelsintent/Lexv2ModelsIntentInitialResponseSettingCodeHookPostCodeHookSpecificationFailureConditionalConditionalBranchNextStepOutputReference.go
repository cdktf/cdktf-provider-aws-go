// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/lexv2modelsintent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference interface {
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
	DialogAction() Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepDialogActionList
	DialogActionInput() interface{}
	// Experimental.
	Fqn() *string
	Intent() Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepIntentList
	IntentInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SessionAttributes() *map[string]*string
	SetSessionAttributes(val *map[string]*string)
	SessionAttributesInput() *map[string]*string
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
	PutDialogAction(value interface{})
	PutIntent(value interface{})
	ResetDialogAction()
	ResetIntent()
	ResetSessionAttributes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference
type jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) DialogAction() Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepDialogActionList {
	var returns Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepDialogActionList
	_jsii_.Get(
		j,
		"dialogAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) DialogActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dialogActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) Intent() Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepIntentList {
	var returns Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepIntentList
	_jsii_.Get(
		j,
		"intent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) IntentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) SessionAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sessionAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) SessionAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sessionAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference {
	_init_.Initialize()

	if err := validateNewLexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference_Override(l Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference)SetSessionAttributes(val *map[string]*string) {
	if err := j.validateSetSessionAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAttributes",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) PutDialogAction(value interface{}) {
	if err := l.validatePutDialogActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDialogAction",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) PutIntent(value interface{}) {
	if err := l.validatePutIntentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIntent",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) ResetDialogAction() {
	_jsii_.InvokeVoid(
		l,
		"resetDialogAction",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) ResetIntent() {
	_jsii_.InvokeVoid(
		l,
		"resetIntent",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) ResetSessionAttributes() {
	_jsii_.InvokeVoid(
		l,
		"resetSessionAttributes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

