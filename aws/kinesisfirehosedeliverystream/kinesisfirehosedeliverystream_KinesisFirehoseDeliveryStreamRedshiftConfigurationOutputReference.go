package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/kinesisfirehosedeliverystream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference interface {
	cdktf.ComplexObject
	CloudwatchLoggingOptions() KinesisFirehoseDeliveryStreamRedshiftConfigurationCloudwatchLoggingOptionsOutputReference
	CloudwatchLoggingOptionsInput() *KinesisFirehoseDeliveryStreamRedshiftConfigurationCloudwatchLoggingOptions
	ClusterJdbcurl() *string
	SetClusterJdbcurl(val *string)
	ClusterJdbcurlInput() *string
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
	CopyOptions() *string
	SetCopyOptions(val *string)
	CopyOptionsInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataTableColumns() *string
	SetDataTableColumns(val *string)
	DataTableColumnsInput() *string
	DataTableName() *string
	SetDataTableName(val *string)
	DataTableNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KinesisFirehoseDeliveryStreamRedshiftConfiguration
	SetInternalValue(val *KinesisFirehoseDeliveryStreamRedshiftConfiguration)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	ProcessingConfiguration() KinesisFirehoseDeliveryStreamRedshiftConfigurationProcessingConfigurationOutputReference
	ProcessingConfigurationInput() *KinesisFirehoseDeliveryStreamRedshiftConfigurationProcessingConfiguration
	RetryDuration() *float64
	SetRetryDuration(val *float64)
	RetryDurationInput() *float64
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	S3BackupConfiguration() KinesisFirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationOutputReference
	S3BackupConfigurationInput() *KinesisFirehoseDeliveryStreamRedshiftConfigurationS3BackupConfiguration
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
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	PutCloudwatchLoggingOptions(value *KinesisFirehoseDeliveryStreamRedshiftConfigurationCloudwatchLoggingOptions)
	PutProcessingConfiguration(value *KinesisFirehoseDeliveryStreamRedshiftConfigurationProcessingConfiguration)
	PutS3BackupConfiguration(value *KinesisFirehoseDeliveryStreamRedshiftConfigurationS3BackupConfiguration)
	ResetCloudwatchLoggingOptions()
	ResetCopyOptions()
	ResetDataTableColumns()
	ResetProcessingConfiguration()
	ResetRetryDuration()
	ResetS3BackupConfiguration()
	ResetS3BackupMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference
type jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) CloudwatchLoggingOptions() KinesisFirehoseDeliveryStreamRedshiftConfigurationCloudwatchLoggingOptionsOutputReference {
	var returns KinesisFirehoseDeliveryStreamRedshiftConfigurationCloudwatchLoggingOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) CloudwatchLoggingOptionsInput() *KinesisFirehoseDeliveryStreamRedshiftConfigurationCloudwatchLoggingOptions {
	var returns *KinesisFirehoseDeliveryStreamRedshiftConfigurationCloudwatchLoggingOptions
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ClusterJdbcurl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterJdbcurl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ClusterJdbcurlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterJdbcurlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) CopyOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) CopyOptionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) DataTableColumns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTableColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) DataTableColumnsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTableColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) DataTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) DataTableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) InternalValue() *KinesisFirehoseDeliveryStreamRedshiftConfiguration {
	var returns *KinesisFirehoseDeliveryStreamRedshiftConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ProcessingConfiguration() KinesisFirehoseDeliveryStreamRedshiftConfigurationProcessingConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamRedshiftConfigurationProcessingConfigurationOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ProcessingConfigurationInput() *KinesisFirehoseDeliveryStreamRedshiftConfigurationProcessingConfiguration {
	var returns *KinesisFirehoseDeliveryStreamRedshiftConfigurationProcessingConfiguration
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) RetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) RetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) S3BackupConfiguration() KinesisFirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3BackupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) S3BackupConfigurationInput() *KinesisFirehoseDeliveryStreamRedshiftConfigurationS3BackupConfiguration {
	var returns *KinesisFirehoseDeliveryStreamRedshiftConfigurationS3BackupConfiguration
	_jsii_.Get(
		j,
		"s3BackupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewKinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference_Override(k KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference)SetClusterJdbcurl(val *string) {
	if err := j.validateSetClusterJdbcurlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterJdbcurl",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference)SetCopyOptions(val *string) {
	if err := j.validateSetCopyOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyOptions",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference)SetDataTableColumns(val *string) {
	if err := j.validateSetDataTableColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTableColumns",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference)SetDataTableName(val *string) {
	if err := j.validateSetDataTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTableName",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference)SetInternalValue(val *KinesisFirehoseDeliveryStreamRedshiftConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference)SetRetryDuration(val *float64) {
	if err := j.validateSetRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryDuration",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) PutCloudwatchLoggingOptions(value *KinesisFirehoseDeliveryStreamRedshiftConfigurationCloudwatchLoggingOptions) {
	if err := k.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) PutProcessingConfiguration(value *KinesisFirehoseDeliveryStreamRedshiftConfigurationProcessingConfiguration) {
	if err := k.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) PutS3BackupConfiguration(value *KinesisFirehoseDeliveryStreamRedshiftConfigurationS3BackupConfiguration) {
	if err := k.validatePutS3BackupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putS3BackupConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ResetCopyOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetCopyOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ResetDataTableColumns() {
	_jsii_.InvokeVoid(
		k,
		"resetDataTableColumns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ResetRetryDuration() {
	_jsii_.InvokeVoid(
		k,
		"resetRetryDuration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ResetS3BackupConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetS3BackupConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		k,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

