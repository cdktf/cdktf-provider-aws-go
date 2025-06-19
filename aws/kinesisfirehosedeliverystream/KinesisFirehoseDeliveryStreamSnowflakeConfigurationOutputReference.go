// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/kinesisfirehosedeliverystream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference interface {
	cdktf.ComplexObject
	AccountUrl() *string
	SetAccountUrl(val *string)
	AccountUrlInput() *string
	BufferingInterval() *float64
	SetBufferingInterval(val *float64)
	BufferingIntervalInput() *float64
	BufferingSize() *float64
	SetBufferingSize(val *float64)
	BufferingSizeInput() *float64
	CloudwatchLoggingOptions() KinesisFirehoseDeliveryStreamSnowflakeConfigurationCloudwatchLoggingOptionsOutputReference
	CloudwatchLoggingOptionsInput() *KinesisFirehoseDeliveryStreamSnowflakeConfigurationCloudwatchLoggingOptions
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
	ContentColumnName() *string
	SetContentColumnName(val *string)
	ContentColumnNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	DataLoadingOption() *string
	SetDataLoadingOption(val *string)
	DataLoadingOptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KinesisFirehoseDeliveryStreamSnowflakeConfiguration
	SetInternalValue(val *KinesisFirehoseDeliveryStreamSnowflakeConfiguration)
	KeyPassphrase() *string
	SetKeyPassphrase(val *string)
	KeyPassphraseInput() *string
	MetadataColumnName() *string
	SetMetadataColumnName(val *string)
	MetadataColumnNameInput() *string
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	ProcessingConfiguration() KinesisFirehoseDeliveryStreamSnowflakeConfigurationProcessingConfigurationOutputReference
	ProcessingConfigurationInput() *KinesisFirehoseDeliveryStreamSnowflakeConfigurationProcessingConfiguration
	RetryDuration() *float64
	SetRetryDuration(val *float64)
	RetryDurationInput() *float64
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	S3BackupMode() *string
	SetS3BackupMode(val *string)
	S3BackupModeInput() *string
	S3Configuration() KinesisFirehoseDeliveryStreamSnowflakeConfigurationS3ConfigurationOutputReference
	S3ConfigurationInput() *KinesisFirehoseDeliveryStreamSnowflakeConfigurationS3Configuration
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	SecretsManagerConfiguration() KinesisFirehoseDeliveryStreamSnowflakeConfigurationSecretsManagerConfigurationOutputReference
	SecretsManagerConfigurationInput() *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSecretsManagerConfiguration
	SnowflakeRoleConfiguration() KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeRoleConfigurationOutputReference
	SnowflakeRoleConfigurationInput() *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeRoleConfiguration
	SnowflakeVpcConfiguration() KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeVpcConfigurationOutputReference
	SnowflakeVpcConfigurationInput() *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeVpcConfiguration
	Table() *string
	SetTable(val *string)
	TableInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	User() *string
	SetUser(val *string)
	UserInput() *string
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
	PutCloudwatchLoggingOptions(value *KinesisFirehoseDeliveryStreamSnowflakeConfigurationCloudwatchLoggingOptions)
	PutProcessingConfiguration(value *KinesisFirehoseDeliveryStreamSnowflakeConfigurationProcessingConfiguration)
	PutS3Configuration(value *KinesisFirehoseDeliveryStreamSnowflakeConfigurationS3Configuration)
	PutSecretsManagerConfiguration(value *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSecretsManagerConfiguration)
	PutSnowflakeRoleConfiguration(value *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeRoleConfiguration)
	PutSnowflakeVpcConfiguration(value *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeVpcConfiguration)
	ResetBufferingInterval()
	ResetBufferingSize()
	ResetCloudwatchLoggingOptions()
	ResetContentColumnName()
	ResetDataLoadingOption()
	ResetKeyPassphrase()
	ResetMetadataColumnName()
	ResetPrivateKey()
	ResetProcessingConfiguration()
	ResetRetryDuration()
	ResetS3BackupMode()
	ResetSecretsManagerConfiguration()
	ResetSnowflakeRoleConfiguration()
	ResetSnowflakeVpcConfiguration()
	ResetUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference
type jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) AccountUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) AccountUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) BufferingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) BufferingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) BufferingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) BufferingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) CloudwatchLoggingOptions() KinesisFirehoseDeliveryStreamSnowflakeConfigurationCloudwatchLoggingOptionsOutputReference {
	var returns KinesisFirehoseDeliveryStreamSnowflakeConfigurationCloudwatchLoggingOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) CloudwatchLoggingOptionsInput() *KinesisFirehoseDeliveryStreamSnowflakeConfigurationCloudwatchLoggingOptions {
	var returns *KinesisFirehoseDeliveryStreamSnowflakeConfigurationCloudwatchLoggingOptions
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ContentColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ContentColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) DataLoadingOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLoadingOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) DataLoadingOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLoadingOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) InternalValue() *KinesisFirehoseDeliveryStreamSnowflakeConfiguration {
	var returns *KinesisFirehoseDeliveryStreamSnowflakeConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) KeyPassphrase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPassphrase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) KeyPassphraseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPassphraseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) MetadataColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) MetadataColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ProcessingConfiguration() KinesisFirehoseDeliveryStreamSnowflakeConfigurationProcessingConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamSnowflakeConfigurationProcessingConfigurationOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ProcessingConfigurationInput() *KinesisFirehoseDeliveryStreamSnowflakeConfigurationProcessingConfiguration {
	var returns *KinesisFirehoseDeliveryStreamSnowflakeConfigurationProcessingConfiguration
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) RetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) RetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) S3Configuration() KinesisFirehoseDeliveryStreamSnowflakeConfigurationS3ConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamSnowflakeConfigurationS3ConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) S3ConfigurationInput() *KinesisFirehoseDeliveryStreamSnowflakeConfigurationS3Configuration {
	var returns *KinesisFirehoseDeliveryStreamSnowflakeConfigurationS3Configuration
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) SecretsManagerConfiguration() KinesisFirehoseDeliveryStreamSnowflakeConfigurationSecretsManagerConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamSnowflakeConfigurationSecretsManagerConfigurationOutputReference
	_jsii_.Get(
		j,
		"secretsManagerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) SecretsManagerConfigurationInput() *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSecretsManagerConfiguration {
	var returns *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSecretsManagerConfiguration
	_jsii_.Get(
		j,
		"secretsManagerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) SnowflakeRoleConfiguration() KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeRoleConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeRoleConfigurationOutputReference
	_jsii_.Get(
		j,
		"snowflakeRoleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) SnowflakeRoleConfigurationInput() *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeRoleConfiguration {
	var returns *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeRoleConfiguration
	_jsii_.Get(
		j,
		"snowflakeRoleConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) SnowflakeVpcConfiguration() KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeVpcConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeVpcConfigurationOutputReference
	_jsii_.Get(
		j,
		"snowflakeVpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) SnowflakeVpcConfigurationInput() *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeVpcConfiguration {
	var returns *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeVpcConfiguration
	_jsii_.Get(
		j,
		"snowflakeVpcConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) Table() *string {
	var returns *string
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) TableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewKinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference_Override(k KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetAccountUrl(val *string) {
	if err := j.validateSetAccountUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountUrl",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetBufferingInterval(val *float64) {
	if err := j.validateSetBufferingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingInterval",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetBufferingSize(val *float64) {
	if err := j.validateSetBufferingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingSize",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetContentColumnName(val *string) {
	if err := j.validateSetContentColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentColumnName",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetDataLoadingOption(val *string) {
	if err := j.validateSetDataLoadingOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataLoadingOption",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetInternalValue(val *KinesisFirehoseDeliveryStreamSnowflakeConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetKeyPassphrase(val *string) {
	if err := j.validateSetKeyPassphraseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPassphrase",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetMetadataColumnName(val *string) {
	if err := j.validateSetMetadataColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataColumnName",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetPrivateKey(val *string) {
	if err := j.validateSetPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetRetryDuration(val *float64) {
	if err := j.validateSetRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryDuration",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetTable(val *string) {
	if err := j.validateSetTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"table",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) PutCloudwatchLoggingOptions(value *KinesisFirehoseDeliveryStreamSnowflakeConfigurationCloudwatchLoggingOptions) {
	if err := k.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) PutProcessingConfiguration(value *KinesisFirehoseDeliveryStreamSnowflakeConfigurationProcessingConfiguration) {
	if err := k.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) PutS3Configuration(value *KinesisFirehoseDeliveryStreamSnowflakeConfigurationS3Configuration) {
	if err := k.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) PutSecretsManagerConfiguration(value *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSecretsManagerConfiguration) {
	if err := k.validatePutSecretsManagerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSecretsManagerConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) PutSnowflakeRoleConfiguration(value *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeRoleConfiguration) {
	if err := k.validatePutSnowflakeRoleConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSnowflakeRoleConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) PutSnowflakeVpcConfiguration(value *KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeVpcConfiguration) {
	if err := k.validatePutSnowflakeVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSnowflakeVpcConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetBufferingInterval() {
	_jsii_.InvokeVoid(
		k,
		"resetBufferingInterval",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetBufferingSize() {
	_jsii_.InvokeVoid(
		k,
		"resetBufferingSize",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetContentColumnName() {
	_jsii_.InvokeVoid(
		k,
		"resetContentColumnName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetDataLoadingOption() {
	_jsii_.InvokeVoid(
		k,
		"resetDataLoadingOption",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetKeyPassphrase() {
	_jsii_.InvokeVoid(
		k,
		"resetKeyPassphrase",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetMetadataColumnName() {
	_jsii_.InvokeVoid(
		k,
		"resetMetadataColumnName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		k,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetRetryDuration() {
	_jsii_.InvokeVoid(
		k,
		"resetRetryDuration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		k,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetSecretsManagerConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSecretsManagerConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetSnowflakeRoleConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSnowflakeRoleConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetSnowflakeVpcConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSnowflakeVpcConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		k,
		"resetUser",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

