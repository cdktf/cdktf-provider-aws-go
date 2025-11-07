// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/bedrockagentagent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference interface {
	cdktf.ComplexObject
	BasePromptTemplate() *string
	SetBasePromptTemplate(val *string)
	BasePromptTemplateInput() *string
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
	InferenceConfiguration() BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsInferenceConfigurationList
	InferenceConfigurationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ParserMode() *string
	SetParserMode(val *string)
	ParserModeInput() *string
	PromptCreationMode() *string
	SetPromptCreationMode(val *string)
	PromptCreationModeInput() *string
	PromptState() *string
	SetPromptState(val *string)
	PromptStateInput() *string
	PromptType() *string
	SetPromptType(val *string)
	PromptTypeInput() *string
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
	PutInferenceConfiguration(value interface{})
	ResetBasePromptTemplate()
	ResetInferenceConfiguration()
	ResetParserMode()
	ResetPromptCreationMode()
	ResetPromptState()
	ResetPromptType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference
type jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) BasePromptTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basePromptTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) BasePromptTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basePromptTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) InferenceConfiguration() BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsInferenceConfigurationList {
	var returns BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsInferenceConfigurationList
	_jsii_.Get(
		j,
		"inferenceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) InferenceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ParserMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parserMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ParserModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parserModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) PromptCreationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptCreationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) PromptCreationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptCreationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) PromptState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) PromptStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) PromptType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) PromptTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentAgent.BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference_Override(b BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentAgent.BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetBasePromptTemplate(val *string) {
	if err := j.validateSetBasePromptTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"basePromptTemplate",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetParserMode(val *string) {
	if err := j.validateSetParserModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parserMode",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetPromptCreationMode(val *string) {
	if err := j.validateSetPromptCreationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promptCreationMode",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetPromptState(val *string) {
	if err := j.validateSetPromptStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promptState",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetPromptType(val *string) {
	if err := j.validateSetPromptTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promptType",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) PutInferenceConfiguration(value interface{}) {
	if err := b.validatePutInferenceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putInferenceConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ResetBasePromptTemplate() {
	_jsii_.InvokeVoid(
		b,
		"resetBasePromptTemplate",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ResetInferenceConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetInferenceConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ResetParserMode() {
	_jsii_.InvokeVoid(
		b,
		"resetParserMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ResetPromptCreationMode() {
	_jsii_.InvokeVoid(
		b,
		"resetPromptCreationMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ResetPromptState() {
	_jsii_.InvokeVoid(
		b,
		"resetPromptState",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ResetPromptType() {
	_jsii_.InvokeVoid(
		b,
		"resetPromptType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

