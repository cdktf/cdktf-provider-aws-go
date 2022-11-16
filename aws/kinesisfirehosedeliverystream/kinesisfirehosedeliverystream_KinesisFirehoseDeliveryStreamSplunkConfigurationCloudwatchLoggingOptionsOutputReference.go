package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/kinesisfirehosedeliverystream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptions
	SetInternalValue(val *KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptions)
	LogGroupName() *string
	SetLogGroupName(val *string)
	LogGroupNameInput() *string
	LogStreamName() *string
	SetLogStreamName(val *string)
	LogStreamNameInput() *string
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
	ResetEnabled()
	ResetLogGroupName()
	ResetLogStreamName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference
type jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) InternalValue() *KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptions {
	var returns *KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) LogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) LogStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) LogStreamNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference_Override(k KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference)SetInternalValue(val *KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference)SetLogGroupName(val *string) {
	if err := j.validateSetLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference)SetLogStreamName(val *string) {
	if err := j.validateSetLogStreamNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStreamName",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) ResetLogGroupName() {
	_jsii_.InvokeVoid(
		k,
		"resetLogGroupName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) ResetLogStreamName() {
	_jsii_.InvokeVoid(
		k,
		"resetLogStreamName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

