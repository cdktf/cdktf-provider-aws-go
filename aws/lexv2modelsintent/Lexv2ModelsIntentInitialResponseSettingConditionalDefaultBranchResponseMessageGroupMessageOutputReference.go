// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/lexv2modelsintent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference interface {
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
	CustomPayload() Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageCustomPayloadList
	CustomPayloadInput() interface{}
	// Experimental.
	Fqn() *string
	ImageResponseCard() Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardList
	ImageResponseCardInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PlainTextMessage() Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessagePlainTextMessageList
	PlainTextMessageInput() interface{}
	SsmlMessage() Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageSsmlMessageList
	SsmlMessageInput() interface{}
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
	PutCustomPayload(value interface{})
	PutImageResponseCard(value interface{})
	PutPlainTextMessage(value interface{})
	PutSsmlMessage(value interface{})
	ResetCustomPayload()
	ResetImageResponseCard()
	ResetPlainTextMessage()
	ResetSsmlMessage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference
type jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) CustomPayload() Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageCustomPayloadList {
	var returns Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageCustomPayloadList
	_jsii_.Get(
		j,
		"customPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) CustomPayloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) ImageResponseCard() Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardList {
	var returns Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardList
	_jsii_.Get(
		j,
		"imageResponseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) ImageResponseCardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageResponseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) PlainTextMessage() Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessagePlainTextMessageList {
	var returns Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessagePlainTextMessageList
	_jsii_.Get(
		j,
		"plainTextMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) PlainTextMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plainTextMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) SsmlMessage() Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageSsmlMessageList {
	var returns Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageSsmlMessageList
	_jsii_.Get(
		j,
		"ssmlMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) SsmlMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssmlMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference {
	_init_.Initialize()

	if err := validateNewLexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference_Override(l Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) PutCustomPayload(value interface{}) {
	if err := l.validatePutCustomPayloadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCustomPayload",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) PutImageResponseCard(value interface{}) {
	if err := l.validatePutImageResponseCardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putImageResponseCard",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) PutPlainTextMessage(value interface{}) {
	if err := l.validatePutPlainTextMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPlainTextMessage",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) PutSsmlMessage(value interface{}) {
	if err := l.validatePutSsmlMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSsmlMessage",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) ResetCustomPayload() {
	_jsii_.InvokeVoid(
		l,
		"resetCustomPayload",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) ResetImageResponseCard() {
	_jsii_.InvokeVoid(
		l,
		"resetImageResponseCard",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) ResetPlainTextMessage() {
	_jsii_.InvokeVoid(
		l,
		"resetPlainTextMessage",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) ResetSsmlMessage() {
	_jsii_.InvokeVoid(
		l,
		"resetSsmlMessage",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntentInitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

