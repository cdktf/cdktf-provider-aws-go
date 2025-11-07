// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/kinesisfirehosedeliverystream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference interface {
	cdktf.ComplexObject
	BucketArn() *string
	SetBucketArn(val *string)
	BucketArnInput() *string
	BufferingInterval() *float64
	SetBufferingInterval(val *float64)
	BufferingIntervalInput() *float64
	BufferingSize() *float64
	SetBufferingSize(val *float64)
	BufferingSizeInput() *float64
	CloudwatchLoggingOptions() KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference
	CloudwatchLoggingOptionsInput() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptions
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
	CompressionFormat() *string
	SetCompressionFormat(val *string)
	CompressionFormatInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ErrorOutputPrefix() *string
	SetErrorOutputPrefix(val *string)
	ErrorOutputPrefixInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3Configuration
	SetInternalValue(val *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3Configuration)
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	PutCloudwatchLoggingOptions(value *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptions)
	ResetBufferingInterval()
	ResetBufferingSize()
	ResetCloudwatchLoggingOptions()
	ResetCompressionFormat()
	ResetErrorOutputPrefix()
	ResetKmsKeyArn()
	ResetPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference
type jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) BucketArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) BufferingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) BufferingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) BufferingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) BufferingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) CloudwatchLoggingOptions() KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference {
	var returns KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) CloudwatchLoggingOptionsInput() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptions {
	var returns *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptions
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) CompressionFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) CompressionFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) ErrorOutputPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOutputPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) ErrorOutputPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOutputPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) InternalValue() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3Configuration {
	var returns *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3Configuration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference_Override(k KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference)SetBucketArn(val *string) {
	if err := j.validateSetBucketArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketArn",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference)SetBufferingInterval(val *float64) {
	if err := j.validateSetBufferingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingInterval",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference)SetBufferingSize(val *float64) {
	if err := j.validateSetBufferingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingSize",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference)SetCompressionFormat(val *string) {
	if err := j.validateSetCompressionFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionFormat",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference)SetErrorOutputPrefix(val *string) {
	if err := j.validateSetErrorOutputPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOutputPrefix",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference)SetInternalValue(val *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3Configuration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) PutCloudwatchLoggingOptions(value *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptions) {
	if err := k.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) ResetBufferingInterval() {
	_jsii_.InvokeVoid(
		k,
		"resetBufferingInterval",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) ResetBufferingSize() {
	_jsii_.InvokeVoid(
		k,
		"resetBufferingSize",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) ResetCompressionFormat() {
	_jsii_.InvokeVoid(
		k,
		"resetCompressionFormat",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) ResetErrorOutputPrefix() {
	_jsii_.InvokeVoid(
		k,
		"resetErrorOutputPrefix",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		k,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		k,
		"resetPrefix",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

