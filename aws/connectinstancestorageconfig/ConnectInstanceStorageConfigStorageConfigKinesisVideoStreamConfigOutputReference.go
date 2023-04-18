package connectinstancestorageconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/connectinstancestorageconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference interface {
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
	EncryptionConfig() ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfigOutputReference
	EncryptionConfigInput() *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfig
	// Experimental.
	Fqn() *string
	InternalValue() *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfig
	SetInternalValue(val *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfig)
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	RetentionPeriodHours() *float64
	SetRetentionPeriodHours(val *float64)
	RetentionPeriodHoursInput() *float64
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
	PutEncryptionConfig(value *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfig)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference
type jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) EncryptionConfig() ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfigOutputReference {
	var returns ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) EncryptionConfigInput() *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfig {
	var returns *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) InternalValue() *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfig {
	var returns *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) RetentionPeriodHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) RetentionPeriodHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference {
	_init_.Initialize()

	if err := validateNewConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.connectInstanceStorageConfig.ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference_Override(c ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.connectInstanceStorageConfig.ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference)SetInternalValue(val *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference)SetRetentionPeriodHours(val *float64) {
	if err := j.validateSetRetentionPeriodHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionPeriodHours",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) PutEncryptionConfig(value *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfig) {
	if err := c.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

