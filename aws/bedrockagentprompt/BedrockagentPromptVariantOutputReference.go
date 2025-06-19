// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentprompt

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/bedrockagentprompt/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockagentPromptVariantOutputReference interface {
	cdktf.ComplexObject
	AdditionalModelRequestFields() *string
	SetAdditionalModelRequestFields(val *string)
	AdditionalModelRequestFieldsInput() *string
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
	GenAiResource() BedrockagentPromptVariantGenAiResourceList
	GenAiResourceInput() interface{}
	InferenceConfiguration() BedrockagentPromptVariantInferenceConfigurationList
	InferenceConfigurationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Metadata() BedrockagentPromptVariantMetadataList
	MetadataInput() interface{}
	ModelId() *string
	SetModelId(val *string)
	ModelIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	TemplateConfiguration() BedrockagentPromptVariantTemplateConfigurationList
	TemplateConfigurationInput() interface{}
	TemplateType() *string
	SetTemplateType(val *string)
	TemplateTypeInput() *string
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
	PutGenAiResource(value interface{})
	PutInferenceConfiguration(value interface{})
	PutMetadata(value interface{})
	PutTemplateConfiguration(value interface{})
	ResetAdditionalModelRequestFields()
	ResetGenAiResource()
	ResetInferenceConfiguration()
	ResetMetadata()
	ResetModelId()
	ResetTemplateConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentPromptVariantOutputReference
type jsiiProxy_BedrockagentPromptVariantOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) AdditionalModelRequestFields() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalModelRequestFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) AdditionalModelRequestFieldsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalModelRequestFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) GenAiResource() BedrockagentPromptVariantGenAiResourceList {
	var returns BedrockagentPromptVariantGenAiResourceList
	_jsii_.Get(
		j,
		"genAiResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) GenAiResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"genAiResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) InferenceConfiguration() BedrockagentPromptVariantInferenceConfigurationList {
	var returns BedrockagentPromptVariantInferenceConfigurationList
	_jsii_.Get(
		j,
		"inferenceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) InferenceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) Metadata() BedrockagentPromptVariantMetadataList {
	var returns BedrockagentPromptVariantMetadataList
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) ModelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) TemplateConfiguration() BedrockagentPromptVariantTemplateConfigurationList {
	var returns BedrockagentPromptVariantTemplateConfigurationList
	_jsii_.Get(
		j,
		"templateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) TemplateConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) TemplateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) TemplateTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentPromptVariantOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentPromptVariantOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentPromptVariantOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentPromptVariantOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentPrompt.BedrockagentPromptVariantOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentPromptVariantOutputReference_Override(b BedrockagentPromptVariantOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentPrompt.BedrockagentPromptVariantOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference)SetAdditionalModelRequestFields(val *string) {
	if err := j.validateSetAdditionalModelRequestFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalModelRequestFields",
		val,
	)
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference)SetModelId(val *string) {
	if err := j.validateSetModelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelId",
		val,
	)
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference)SetTemplateType(val *string) {
	if err := j.validateSetTemplateTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateType",
		val,
	)
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentPromptVariantOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) PutGenAiResource(value interface{}) {
	if err := b.validatePutGenAiResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putGenAiResource",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) PutInferenceConfiguration(value interface{}) {
	if err := b.validatePutInferenceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putInferenceConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) PutMetadata(value interface{}) {
	if err := b.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMetadata",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) PutTemplateConfiguration(value interface{}) {
	if err := b.validatePutTemplateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTemplateConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) ResetAdditionalModelRequestFields() {
	_jsii_.InvokeVoid(
		b,
		"resetAdditionalModelRequestFields",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) ResetGenAiResource() {
	_jsii_.InvokeVoid(
		b,
		"resetGenAiResource",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) ResetInferenceConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetInferenceConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		b,
		"resetMetadata",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) ResetModelId() {
	_jsii_.InvokeVoid(
		b,
		"resetModelId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) ResetTemplateConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetTemplateConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentPromptVariantOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

