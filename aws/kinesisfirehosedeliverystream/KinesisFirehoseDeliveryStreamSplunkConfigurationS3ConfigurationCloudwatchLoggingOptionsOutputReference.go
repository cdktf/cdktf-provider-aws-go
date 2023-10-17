// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/kinesisfirehosedeliverystream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference interface {
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
	InternalValue() *KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptions
	SetInternalValue(val *KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptions)
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

// The jsii proxy struct for KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference
type jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) InternalValue() *KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptions {
	var returns *KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) LogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) LogStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) LogStreamNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference_Override(k KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)SetInternalValue(val *KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)SetLogGroupName(val *string) {
	if err := j.validateSetLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)SetLogStreamName(val *string) {
	if err := j.validateSetLogStreamNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStreamName",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) ResetLogGroupName() {
	_jsii_.InvokeVoid(
		k,
		"resetLogGroupName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) ResetLogStreamName() {
	_jsii_.InvokeVoid(
		k,
		"resetLogStreamName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamSplunkConfigurationS3ConfigurationCloudwatchLoggingOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

