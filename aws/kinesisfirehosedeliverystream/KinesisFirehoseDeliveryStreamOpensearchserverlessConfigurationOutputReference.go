// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/kinesisfirehosedeliverystream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference interface {
	cdktf.ComplexObject
	BufferingInterval() *float64
	SetBufferingInterval(val *float64)
	BufferingIntervalInput() *float64
	BufferingSize() *float64
	SetBufferingSize(val *float64)
	BufferingSizeInput() *float64
	CloudwatchLoggingOptions() KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationCloudwatchLoggingOptionsOutputReference
	CloudwatchLoggingOptionsInput() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationCloudwatchLoggingOptions
	CollectionEndpoint() *string
	SetCollectionEndpoint(val *string)
	CollectionEndpointInput() *string
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
	IndexName() *string
	SetIndexName(val *string)
	IndexNameInput() *string
	InternalValue() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfiguration
	SetInternalValue(val *KinesisFirehoseDeliveryStreamOpensearchserverlessConfiguration)
	ProcessingConfiguration() KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationOutputReference
	ProcessingConfigurationInput() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfiguration
	RetryDuration() *float64
	SetRetryDuration(val *float64)
	RetryDurationInput() *float64
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	S3BackupMode() *string
	SetS3BackupMode(val *string)
	S3BackupModeInput() *string
	S3Configuration() KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference
	S3ConfigurationInput() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3Configuration
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcConfig() KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationVpcConfigOutputReference
	VpcConfigInput() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationVpcConfig
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
	PutCloudwatchLoggingOptions(value *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationCloudwatchLoggingOptions)
	PutProcessingConfiguration(value *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfiguration)
	PutS3Configuration(value *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3Configuration)
	PutVpcConfig(value *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationVpcConfig)
	ResetBufferingInterval()
	ResetBufferingSize()
	ResetCloudwatchLoggingOptions()
	ResetProcessingConfiguration()
	ResetRetryDuration()
	ResetS3BackupMode()
	ResetVpcConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference
type jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) BufferingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) BufferingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) BufferingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) BufferingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) CloudwatchLoggingOptions() KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationCloudwatchLoggingOptionsOutputReference {
	var returns KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationCloudwatchLoggingOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) CloudwatchLoggingOptionsInput() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationCloudwatchLoggingOptions {
	var returns *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationCloudwatchLoggingOptions
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) CollectionEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) CollectionEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) IndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) InternalValue() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfiguration {
	var returns *KinesisFirehoseDeliveryStreamOpensearchserverlessConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) ProcessingConfiguration() KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) ProcessingConfigurationInput() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfiguration {
	var returns *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfiguration
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) RetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) RetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) S3Configuration() KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) S3ConfigurationInput() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3Configuration {
	var returns *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3Configuration
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) VpcConfig() KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationVpcConfigOutputReference {
	var returns KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) VpcConfigInput() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationVpcConfig {
	var returns *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


func NewKinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference_Override(k KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference)SetBufferingInterval(val *float64) {
	if err := j.validateSetBufferingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingInterval",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference)SetBufferingSize(val *float64) {
	if err := j.validateSetBufferingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingSize",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference)SetCollectionEndpoint(val *string) {
	if err := j.validateSetCollectionEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionEndpoint",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference)SetIndexName(val *string) {
	if err := j.validateSetIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference)SetInternalValue(val *KinesisFirehoseDeliveryStreamOpensearchserverlessConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference)SetRetryDuration(val *float64) {
	if err := j.validateSetRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryDuration",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) PutCloudwatchLoggingOptions(value *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationCloudwatchLoggingOptions) {
	if err := k.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) PutProcessingConfiguration(value *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfiguration) {
	if err := k.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) PutS3Configuration(value *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationS3Configuration) {
	if err := k.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) PutVpcConfig(value *KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationVpcConfig) {
	if err := k.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) ResetBufferingInterval() {
	_jsii_.InvokeVoid(
		k,
		"resetBufferingInterval",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) ResetBufferingSize() {
	_jsii_.InvokeVoid(
		k,
		"resetBufferingSize",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) ResetRetryDuration() {
	_jsii_.InvokeVoid(
		k,
		"resetRetryDuration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		k,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

