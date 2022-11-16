package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/kinesisfirehosedeliverystream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference interface {
	cdktf.ComplexObject
	BufferingInterval() *float64
	SetBufferingInterval(val *float64)
	BufferingIntervalInput() *float64
	BufferingSize() *float64
	SetBufferingSize(val *float64)
	BufferingSizeInput() *float64
	CloudwatchLoggingOptions() KinesisFirehoseDeliveryStreamElasticsearchConfigurationCloudwatchLoggingOptionsOutputReference
	CloudwatchLoggingOptionsInput() *KinesisFirehoseDeliveryStreamElasticsearchConfigurationCloudwatchLoggingOptions
	ClusterEndpoint() *string
	SetClusterEndpoint(val *string)
	ClusterEndpointInput() *string
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
	DomainArn() *string
	SetDomainArn(val *string)
	DomainArnInput() *string
	// Experimental.
	Fqn() *string
	IndexName() *string
	SetIndexName(val *string)
	IndexNameInput() *string
	IndexRotationPeriod() *string
	SetIndexRotationPeriod(val *string)
	IndexRotationPeriodInput() *string
	InternalValue() *KinesisFirehoseDeliveryStreamElasticsearchConfiguration
	SetInternalValue(val *KinesisFirehoseDeliveryStreamElasticsearchConfiguration)
	ProcessingConfiguration() KinesisFirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationOutputReference
	ProcessingConfigurationInput() *KinesisFirehoseDeliveryStreamElasticsearchConfigurationProcessingConfiguration
	RetryDuration() *float64
	SetRetryDuration(val *float64)
	RetryDurationInput() *float64
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	S3BackupMode() *string
	SetS3BackupMode(val *string)
	S3BackupModeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TypeName() *string
	SetTypeName(val *string)
	TypeNameInput() *string
	VpcConfig() KinesisFirehoseDeliveryStreamElasticsearchConfigurationVpcConfigOutputReference
	VpcConfigInput() *KinesisFirehoseDeliveryStreamElasticsearchConfigurationVpcConfig
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
	PutCloudwatchLoggingOptions(value *KinesisFirehoseDeliveryStreamElasticsearchConfigurationCloudwatchLoggingOptions)
	PutProcessingConfiguration(value *KinesisFirehoseDeliveryStreamElasticsearchConfigurationProcessingConfiguration)
	PutVpcConfig(value *KinesisFirehoseDeliveryStreamElasticsearchConfigurationVpcConfig)
	ResetBufferingInterval()
	ResetBufferingSize()
	ResetCloudwatchLoggingOptions()
	ResetClusterEndpoint()
	ResetDomainArn()
	ResetIndexRotationPeriod()
	ResetProcessingConfiguration()
	ResetRetryDuration()
	ResetS3BackupMode()
	ResetTypeName()
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

// The jsii proxy struct for KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference
type jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) BufferingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) BufferingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) BufferingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) BufferingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) CloudwatchLoggingOptions() KinesisFirehoseDeliveryStreamElasticsearchConfigurationCloudwatchLoggingOptionsOutputReference {
	var returns KinesisFirehoseDeliveryStreamElasticsearchConfigurationCloudwatchLoggingOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) CloudwatchLoggingOptionsInput() *KinesisFirehoseDeliveryStreamElasticsearchConfigurationCloudwatchLoggingOptions {
	var returns *KinesisFirehoseDeliveryStreamElasticsearchConfigurationCloudwatchLoggingOptions
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ClusterEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ClusterEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) DomainArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) DomainArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) IndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) IndexRotationPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexRotationPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) IndexRotationPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexRotationPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) InternalValue() *KinesisFirehoseDeliveryStreamElasticsearchConfiguration {
	var returns *KinesisFirehoseDeliveryStreamElasticsearchConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ProcessingConfiguration() KinesisFirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ProcessingConfigurationInput() *KinesisFirehoseDeliveryStreamElasticsearchConfigurationProcessingConfiguration {
	var returns *KinesisFirehoseDeliveryStreamElasticsearchConfigurationProcessingConfiguration
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) RetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) RetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) TypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) VpcConfig() KinesisFirehoseDeliveryStreamElasticsearchConfigurationVpcConfigOutputReference {
	var returns KinesisFirehoseDeliveryStreamElasticsearchConfigurationVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) VpcConfigInput() *KinesisFirehoseDeliveryStreamElasticsearchConfigurationVpcConfig {
	var returns *KinesisFirehoseDeliveryStreamElasticsearchConfigurationVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


func NewKinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference_Override(k KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetBufferingInterval(val *float64) {
	if err := j.validateSetBufferingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingInterval",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetBufferingSize(val *float64) {
	if err := j.validateSetBufferingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingSize",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetClusterEndpoint(val *string) {
	if err := j.validateSetClusterEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterEndpoint",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetDomainArn(val *string) {
	if err := j.validateSetDomainArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainArn",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetIndexName(val *string) {
	if err := j.validateSetIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetIndexRotationPeriod(val *string) {
	if err := j.validateSetIndexRotationPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexRotationPeriod",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetInternalValue(val *KinesisFirehoseDeliveryStreamElasticsearchConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetRetryDuration(val *float64) {
	if err := j.validateSetRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryDuration",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference)SetTypeName(val *string) {
	if err := j.validateSetTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) PutCloudwatchLoggingOptions(value *KinesisFirehoseDeliveryStreamElasticsearchConfigurationCloudwatchLoggingOptions) {
	if err := k.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) PutProcessingConfiguration(value *KinesisFirehoseDeliveryStreamElasticsearchConfigurationProcessingConfiguration) {
	if err := k.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) PutVpcConfig(value *KinesisFirehoseDeliveryStreamElasticsearchConfigurationVpcConfig) {
	if err := k.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ResetBufferingInterval() {
	_jsii_.InvokeVoid(
		k,
		"resetBufferingInterval",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ResetBufferingSize() {
	_jsii_.InvokeVoid(
		k,
		"resetBufferingSize",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ResetClusterEndpoint() {
	_jsii_.InvokeVoid(
		k,
		"resetClusterEndpoint",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ResetDomainArn() {
	_jsii_.InvokeVoid(
		k,
		"resetDomainArn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ResetIndexRotationPeriod() {
	_jsii_.InvokeVoid(
		k,
		"resetIndexRotationPeriod",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ResetRetryDuration() {
	_jsii_.InvokeVoid(
		k,
		"resetRetryDuration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		k,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ResetTypeName() {
	_jsii_.InvokeVoid(
		k,
		"resetTypeName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

