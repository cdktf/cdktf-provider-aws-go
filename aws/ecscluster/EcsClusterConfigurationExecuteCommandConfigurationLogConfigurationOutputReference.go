package ecscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v15/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v15/ecscluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference interface {
	cdktf.ComplexObject
	CloudWatchEncryptionEnabled() interface{}
	SetCloudWatchEncryptionEnabled(val interface{})
	CloudWatchEncryptionEnabledInput() interface{}
	CloudWatchLogGroupName() *string
	SetCloudWatchLogGroupName(val *string)
	CloudWatchLogGroupNameInput() *string
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
	InternalValue() *EcsClusterConfigurationExecuteCommandConfigurationLogConfiguration
	SetInternalValue(val *EcsClusterConfigurationExecuteCommandConfigurationLogConfiguration)
	S3BucketEncryptionEnabled() interface{}
	SetS3BucketEncryptionEnabled(val interface{})
	S3BucketEncryptionEnabledInput() interface{}
	S3BucketName() *string
	SetS3BucketName(val *string)
	S3BucketNameInput() *string
	S3KeyPrefix() *string
	SetS3KeyPrefix(val *string)
	S3KeyPrefixInput() *string
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
	ResetCloudWatchEncryptionEnabled()
	ResetCloudWatchLogGroupName()
	ResetS3BucketEncryptionEnabled()
	ResetS3BucketName()
	ResetS3KeyPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference
type jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) CloudWatchEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudWatchEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) CloudWatchEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudWatchEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) CloudWatchLogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) CloudWatchLogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) InternalValue() *EcsClusterConfigurationExecuteCommandConfigurationLogConfiguration {
	var returns *EcsClusterConfigurationExecuteCommandConfigurationLogConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) S3BucketEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3BucketEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) S3BucketEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3BucketEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) S3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) S3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) S3KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewEcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecsCluster.EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference_Override(e EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecsCluster.EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetCloudWatchEncryptionEnabled(val interface{}) {
	if err := j.validateSetCloudWatchEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudWatchEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetCloudWatchLogGroupName(val *string) {
	if err := j.validateSetCloudWatchLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudWatchLogGroupName",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetInternalValue(val *EcsClusterConfigurationExecuteCommandConfigurationLogConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetS3BucketEncryptionEnabled(val interface{}) {
	if err := j.validateSetS3BucketEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetS3BucketName(val *string) {
	if err := j.validateSetS3BucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetS3KeyPrefix(val *string) {
	if err := j.validateSetS3KeyPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ResetCloudWatchEncryptionEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetCloudWatchEncryptionEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ResetCloudWatchLogGroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetCloudWatchLogGroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ResetS3BucketEncryptionEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetS3BucketEncryptionEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ResetS3BucketName() {
	_jsii_.InvokeVoid(
		e,
		"resetS3BucketName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ResetS3KeyPrefix() {
	_jsii_.InvokeVoid(
		e,
		"resetS3KeyPrefix",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsClusterConfigurationExecuteCommandConfigurationLogConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

