// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentprompt

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/bedrockagentprompt/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockagentPromptVariantTemplateConfigurationChatOutputReference interface {
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
	InputVariable() BedrockagentPromptVariantTemplateConfigurationChatInputVariableList
	InputVariableInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Message() BedrockagentPromptVariantTemplateConfigurationChatMessageList
	MessageInput() interface{}
	SystemAttribute() BedrockagentPromptVariantTemplateConfigurationChatSystemList
	SystemAttributeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ToolConfiguration() BedrockagentPromptVariantTemplateConfigurationChatToolConfigurationList
	ToolConfigurationInput() interface{}
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
	PutInputVariable(value interface{})
	PutMessage(value interface{})
	PutSystemAttribute(value interface{})
	PutToolConfiguration(value interface{})
	ResetInputVariable()
	ResetMessage()
	ResetSystemAttribute()
	ResetToolConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentPromptVariantTemplateConfigurationChatOutputReference
type jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) InputVariable() BedrockagentPromptVariantTemplateConfigurationChatInputVariableList {
	var returns BedrockagentPromptVariantTemplateConfigurationChatInputVariableList
	_jsii_.Get(
		j,
		"inputVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) InputVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) Message() BedrockagentPromptVariantTemplateConfigurationChatMessageList {
	var returns BedrockagentPromptVariantTemplateConfigurationChatMessageList
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) MessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) SystemAttribute() BedrockagentPromptVariantTemplateConfigurationChatSystemList {
	var returns BedrockagentPromptVariantTemplateConfigurationChatSystemList
	_jsii_.Get(
		j,
		"systemAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) SystemAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) ToolConfiguration() BedrockagentPromptVariantTemplateConfigurationChatToolConfigurationList {
	var returns BedrockagentPromptVariantTemplateConfigurationChatToolConfigurationList
	_jsii_.Get(
		j,
		"toolConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) ToolConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"toolConfigurationInput",
		&returns,
	)
	return returns
}


func NewBedrockagentPromptVariantTemplateConfigurationChatOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentPromptVariantTemplateConfigurationChatOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentPromptVariantTemplateConfigurationChatOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentPrompt.BedrockagentPromptVariantTemplateConfigurationChatOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentPromptVariantTemplateConfigurationChatOutputReference_Override(b BedrockagentPromptVariantTemplateConfigurationChatOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentPrompt.BedrockagentPromptVariantTemplateConfigurationChatOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) PutInputVariable(value interface{}) {
	if err := b.validatePutInputVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putInputVariable",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) PutMessage(value interface{}) {
	if err := b.validatePutMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMessage",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) PutSystemAttribute(value interface{}) {
	if err := b.validatePutSystemAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSystemAttribute",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) PutToolConfiguration(value interface{}) {
	if err := b.validatePutToolConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putToolConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) ResetInputVariable() {
	_jsii_.InvokeVoid(
		b,
		"resetInputVariable",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) ResetMessage() {
	_jsii_.InvokeVoid(
		b,
		"resetMessage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) ResetSystemAttribute() {
	_jsii_.InvokeVoid(
		b,
		"resetSystemAttribute",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) ResetToolConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetToolConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentPromptVariantTemplateConfigurationChatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

