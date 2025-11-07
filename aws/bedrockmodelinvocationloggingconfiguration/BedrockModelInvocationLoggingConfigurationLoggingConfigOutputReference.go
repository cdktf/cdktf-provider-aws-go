// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockmodelinvocationloggingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/bedrockmodelinvocationloggingconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference interface {
	cdktf.ComplexObject
	CloudwatchConfig() BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigList
	CloudwatchConfigInput() interface{}
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
	EmbeddingDataDeliveryEnabled() interface{}
	SetEmbeddingDataDeliveryEnabled(val interface{})
	EmbeddingDataDeliveryEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	ImageDataDeliveryEnabled() interface{}
	SetImageDataDeliveryEnabled(val interface{})
	ImageDataDeliveryEnabledInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3Config() BedrockModelInvocationLoggingConfigurationLoggingConfigS3ConfigList
	S3ConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextDataDeliveryEnabled() interface{}
	SetTextDataDeliveryEnabled(val interface{})
	TextDataDeliveryEnabledInput() interface{}
	VideoDataDeliveryEnabled() interface{}
	SetVideoDataDeliveryEnabled(val interface{})
	VideoDataDeliveryEnabledInput() interface{}
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
	PutCloudwatchConfig(value interface{})
	PutS3Config(value interface{})
	ResetCloudwatchConfig()
	ResetEmbeddingDataDeliveryEnabled()
	ResetImageDataDeliveryEnabled()
	ResetS3Config()
	ResetTextDataDeliveryEnabled()
	ResetVideoDataDeliveryEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference
type jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) CloudwatchConfig() BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigList {
	var returns BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigList
	_jsii_.Get(
		j,
		"cloudwatchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) CloudwatchConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) EmbeddingDataDeliveryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingDataDeliveryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) EmbeddingDataDeliveryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingDataDeliveryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) ImageDataDeliveryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageDataDeliveryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) ImageDataDeliveryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageDataDeliveryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) S3Config() BedrockModelInvocationLoggingConfigurationLoggingConfigS3ConfigList {
	var returns BedrockModelInvocationLoggingConfigurationLoggingConfigS3ConfigList
	_jsii_.Get(
		j,
		"s3Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) S3ConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) TextDataDeliveryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"textDataDeliveryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) TextDataDeliveryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"textDataDeliveryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) VideoDataDeliveryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoDataDeliveryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) VideoDataDeliveryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoDataDeliveryEnabledInput",
		&returns,
	)
	return returns
}


func NewBedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockModelInvocationLoggingConfigurationLoggingConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockModelInvocationLoggingConfiguration.BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference_Override(b BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockModelInvocationLoggingConfiguration.BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference)SetEmbeddingDataDeliveryEnabled(val interface{}) {
	if err := j.validateSetEmbeddingDataDeliveryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embeddingDataDeliveryEnabled",
		val,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference)SetImageDataDeliveryEnabled(val interface{}) {
	if err := j.validateSetImageDataDeliveryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageDataDeliveryEnabled",
		val,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference)SetTextDataDeliveryEnabled(val interface{}) {
	if err := j.validateSetTextDataDeliveryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textDataDeliveryEnabled",
		val,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference)SetVideoDataDeliveryEnabled(val interface{}) {
	if err := j.validateSetVideoDataDeliveryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"videoDataDeliveryEnabled",
		val,
	)
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) PutCloudwatchConfig(value interface{}) {
	if err := b.validatePutCloudwatchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCloudwatchConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) PutS3Config(value interface{}) {
	if err := b.validatePutS3ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putS3Config",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) ResetCloudwatchConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetCloudwatchConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) ResetEmbeddingDataDeliveryEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetEmbeddingDataDeliveryEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) ResetImageDataDeliveryEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetImageDataDeliveryEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) ResetS3Config() {
	_jsii_.InvokeVoid(
		b,
		"resetS3Config",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) ResetTextDataDeliveryEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetTextDataDeliveryEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) ResetVideoDataDeliveryEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetVideoDataDeliveryEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

