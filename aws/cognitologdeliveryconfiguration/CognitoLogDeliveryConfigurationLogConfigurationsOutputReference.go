// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitologdeliveryconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/cognitologdeliveryconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitoLogDeliveryConfigurationLogConfigurationsOutputReference interface {
	cdktf.ComplexObject
	CloudWatchLogsConfiguration() CognitoLogDeliveryConfigurationLogConfigurationsCloudWatchLogsConfigurationList
	CloudWatchLogsConfigurationInput() interface{}
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
	EventSource() *string
	SetEventSource(val *string)
	EventSourceInput() *string
	FirehoseConfiguration() CognitoLogDeliveryConfigurationLogConfigurationsFirehoseConfigurationList
	FirehoseConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	S3Configuration() CognitoLogDeliveryConfigurationLogConfigurationsS3ConfigurationList
	S3ConfigurationInput() interface{}
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
	PutCloudWatchLogsConfiguration(value interface{})
	PutFirehoseConfiguration(value interface{})
	PutS3Configuration(value interface{})
	ResetCloudWatchLogsConfiguration()
	ResetFirehoseConfiguration()
	ResetS3Configuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoLogDeliveryConfigurationLogConfigurationsOutputReference
type jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) CloudWatchLogsConfiguration() CognitoLogDeliveryConfigurationLogConfigurationsCloudWatchLogsConfigurationList {
	var returns CognitoLogDeliveryConfigurationLogConfigurationsCloudWatchLogsConfigurationList
	_jsii_.Get(
		j,
		"cloudWatchLogsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) CloudWatchLogsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudWatchLogsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) EventSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) EventSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) FirehoseConfiguration() CognitoLogDeliveryConfigurationLogConfigurationsFirehoseConfigurationList {
	var returns CognitoLogDeliveryConfigurationLogConfigurationsFirehoseConfigurationList
	_jsii_.Get(
		j,
		"firehoseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) FirehoseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firehoseConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) S3Configuration() CognitoLogDeliveryConfigurationLogConfigurationsS3ConfigurationList {
	var returns CognitoLogDeliveryConfigurationLogConfigurationsS3ConfigurationList
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) S3ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCognitoLogDeliveryConfigurationLogConfigurationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CognitoLogDeliveryConfigurationLogConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoLogDeliveryConfigurationLogConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoLogDeliveryConfiguration.CognitoLogDeliveryConfigurationLogConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCognitoLogDeliveryConfigurationLogConfigurationsOutputReference_Override(c CognitoLogDeliveryConfigurationLogConfigurationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoLogDeliveryConfiguration.CognitoLogDeliveryConfigurationLogConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference)SetEventSource(val *string) {
	if err := j.validateSetEventSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSource",
		val,
	)
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) PutCloudWatchLogsConfiguration(value interface{}) {
	if err := c.validatePutCloudWatchLogsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCloudWatchLogsConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) PutFirehoseConfiguration(value interface{}) {
	if err := c.validatePutFirehoseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFirehoseConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) PutS3Configuration(value interface{}) {
	if err := c.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) ResetCloudWatchLogsConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudWatchLogsConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) ResetFirehoseConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetFirehoseConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) ResetS3Configuration() {
	_jsii_.InvokeVoid(
		c,
		"resetS3Configuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoLogDeliveryConfigurationLogConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

