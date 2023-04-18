package connectinstancestorageconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/connectinstancestorageconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectInstanceStorageConfigStorageConfigOutputReference interface {
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
	InternalValue() *ConnectInstanceStorageConfigStorageConfig
	SetInternalValue(val *ConnectInstanceStorageConfigStorageConfig)
	KinesisFirehoseConfig() ConnectInstanceStorageConfigStorageConfigKinesisFirehoseConfigOutputReference
	KinesisFirehoseConfigInput() *ConnectInstanceStorageConfigStorageConfigKinesisFirehoseConfig
	KinesisStreamConfig() ConnectInstanceStorageConfigStorageConfigKinesisStreamConfigOutputReference
	KinesisStreamConfigInput() *ConnectInstanceStorageConfigStorageConfigKinesisStreamConfig
	KinesisVideoStreamConfig() ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference
	KinesisVideoStreamConfigInput() *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfig
	S3Config() ConnectInstanceStorageConfigStorageConfigS3ConfigOutputReference
	S3ConfigInput() *ConnectInstanceStorageConfigStorageConfigS3Config
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
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
	PutKinesisFirehoseConfig(value *ConnectInstanceStorageConfigStorageConfigKinesisFirehoseConfig)
	PutKinesisStreamConfig(value *ConnectInstanceStorageConfigStorageConfigKinesisStreamConfig)
	PutKinesisVideoStreamConfig(value *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfig)
	PutS3Config(value *ConnectInstanceStorageConfigStorageConfigS3Config)
	ResetKinesisFirehoseConfig()
	ResetKinesisStreamConfig()
	ResetKinesisVideoStreamConfig()
	ResetS3Config()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectInstanceStorageConfigStorageConfigOutputReference
type jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) InternalValue() *ConnectInstanceStorageConfigStorageConfig {
	var returns *ConnectInstanceStorageConfigStorageConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) KinesisFirehoseConfig() ConnectInstanceStorageConfigStorageConfigKinesisFirehoseConfigOutputReference {
	var returns ConnectInstanceStorageConfigStorageConfigKinesisFirehoseConfigOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehoseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) KinesisFirehoseConfigInput() *ConnectInstanceStorageConfigStorageConfigKinesisFirehoseConfig {
	var returns *ConnectInstanceStorageConfigStorageConfigKinesisFirehoseConfig
	_jsii_.Get(
		j,
		"kinesisFirehoseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) KinesisStreamConfig() ConnectInstanceStorageConfigStorageConfigKinesisStreamConfigOutputReference {
	var returns ConnectInstanceStorageConfigStorageConfigKinesisStreamConfigOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) KinesisStreamConfigInput() *ConnectInstanceStorageConfigStorageConfigKinesisStreamConfig {
	var returns *ConnectInstanceStorageConfigStorageConfigKinesisStreamConfig
	_jsii_.Get(
		j,
		"kinesisStreamConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) KinesisVideoStreamConfig() ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference {
	var returns ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference
	_jsii_.Get(
		j,
		"kinesisVideoStreamConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) KinesisVideoStreamConfigInput() *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfig {
	var returns *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfig
	_jsii_.Get(
		j,
		"kinesisVideoStreamConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) S3Config() ConnectInstanceStorageConfigStorageConfigS3ConfigOutputReference {
	var returns ConnectInstanceStorageConfigStorageConfigS3ConfigOutputReference
	_jsii_.Get(
		j,
		"s3Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) S3ConfigInput() *ConnectInstanceStorageConfigStorageConfigS3Config {
	var returns *ConnectInstanceStorageConfigStorageConfigS3Config
	_jsii_.Get(
		j,
		"s3ConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectInstanceStorageConfigStorageConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConnectInstanceStorageConfigStorageConfigOutputReference {
	_init_.Initialize()

	if err := validateNewConnectInstanceStorageConfigStorageConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.connectInstanceStorageConfig.ConnectInstanceStorageConfigStorageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectInstanceStorageConfigStorageConfigOutputReference_Override(c ConnectInstanceStorageConfigStorageConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.connectInstanceStorageConfig.ConnectInstanceStorageConfigStorageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference)SetInternalValue(val *ConnectInstanceStorageConfigStorageConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) PutKinesisFirehoseConfig(value *ConnectInstanceStorageConfigStorageConfigKinesisFirehoseConfig) {
	if err := c.validatePutKinesisFirehoseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKinesisFirehoseConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) PutKinesisStreamConfig(value *ConnectInstanceStorageConfigStorageConfigKinesisStreamConfig) {
	if err := c.validatePutKinesisStreamConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKinesisStreamConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) PutKinesisVideoStreamConfig(value *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfig) {
	if err := c.validatePutKinesisVideoStreamConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKinesisVideoStreamConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) PutS3Config(value *ConnectInstanceStorageConfigStorageConfigS3Config) {
	if err := c.validatePutS3ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putS3Config",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) ResetKinesisFirehoseConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetKinesisFirehoseConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) ResetKinesisStreamConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetKinesisStreamConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) ResetKinesisVideoStreamConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetKinesisVideoStreamConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) ResetS3Config() {
	_jsii_.InvokeVoid(
		c,
		"resetS3Config",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

