// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chimesdkmediapipelinesmediainsightspipelineconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/chimesdkmediapipelinesmediainsightspipelineconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference interface {
	cdktf.ComplexObject
	AmazonTranscribeCallAnalyticsProcessorConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference
	AmazonTranscribeCallAnalyticsProcessorConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfiguration
	AmazonTranscribeProcessorConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference
	AmazonTranscribeProcessorConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfiguration
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KinesisDataStreamSinkConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsKinesisDataStreamSinkConfigurationOutputReference
	KinesisDataStreamSinkConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsKinesisDataStreamSinkConfiguration
	LambdaFunctionSinkConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsLambdaFunctionSinkConfigurationOutputReference
	LambdaFunctionSinkConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsLambdaFunctionSinkConfiguration
	S3RecordingSinkConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference
	S3RecordingSinkConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfiguration
	SnsTopicSinkConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSnsTopicSinkConfigurationOutputReference
	SnsTopicSinkConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSnsTopicSinkConfiguration
	SqsQueueSinkConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSqsQueueSinkConfigurationOutputReference
	SqsQueueSinkConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSqsQueueSinkConfiguration
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VoiceAnalyticsProcessorConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsVoiceAnalyticsProcessorConfigurationOutputReference
	VoiceAnalyticsProcessorConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsVoiceAnalyticsProcessorConfiguration
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
	PutAmazonTranscribeCallAnalyticsProcessorConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfiguration)
	PutAmazonTranscribeProcessorConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfiguration)
	PutKinesisDataStreamSinkConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsKinesisDataStreamSinkConfiguration)
	PutLambdaFunctionSinkConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsLambdaFunctionSinkConfiguration)
	PutS3RecordingSinkConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfiguration)
	PutSnsTopicSinkConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSnsTopicSinkConfiguration)
	PutSqsQueueSinkConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSqsQueueSinkConfiguration)
	PutVoiceAnalyticsProcessorConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsVoiceAnalyticsProcessorConfiguration)
	ResetAmazonTranscribeCallAnalyticsProcessorConfiguration()
	ResetAmazonTranscribeProcessorConfiguration()
	ResetKinesisDataStreamSinkConfiguration()
	ResetLambdaFunctionSinkConfiguration()
	ResetS3RecordingSinkConfiguration()
	ResetSnsTopicSinkConfiguration()
	ResetSqsQueueSinkConfiguration()
	ResetVoiceAnalyticsProcessorConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference
type jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) AmazonTranscribeCallAnalyticsProcessorConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference {
	var returns ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationOutputReference
	_jsii_.Get(
		j,
		"amazonTranscribeCallAnalyticsProcessorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) AmazonTranscribeCallAnalyticsProcessorConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfiguration {
	var returns *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfiguration
	_jsii_.Get(
		j,
		"amazonTranscribeCallAnalyticsProcessorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) AmazonTranscribeProcessorConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference {
	var returns ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference
	_jsii_.Get(
		j,
		"amazonTranscribeProcessorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) AmazonTranscribeProcessorConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfiguration {
	var returns *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfiguration
	_jsii_.Get(
		j,
		"amazonTranscribeProcessorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) KinesisDataStreamSinkConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsKinesisDataStreamSinkConfigurationOutputReference {
	var returns ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsKinesisDataStreamSinkConfigurationOutputReference
	_jsii_.Get(
		j,
		"kinesisDataStreamSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) KinesisDataStreamSinkConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsKinesisDataStreamSinkConfiguration {
	var returns *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsKinesisDataStreamSinkConfiguration
	_jsii_.Get(
		j,
		"kinesisDataStreamSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) LambdaFunctionSinkConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsLambdaFunctionSinkConfigurationOutputReference {
	var returns ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsLambdaFunctionSinkConfigurationOutputReference
	_jsii_.Get(
		j,
		"lambdaFunctionSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) LambdaFunctionSinkConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsLambdaFunctionSinkConfiguration {
	var returns *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsLambdaFunctionSinkConfiguration
	_jsii_.Get(
		j,
		"lambdaFunctionSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) S3RecordingSinkConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference {
	var returns ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3RecordingSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) S3RecordingSinkConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfiguration {
	var returns *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfiguration
	_jsii_.Get(
		j,
		"s3RecordingSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) SnsTopicSinkConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSnsTopicSinkConfigurationOutputReference {
	var returns ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSnsTopicSinkConfigurationOutputReference
	_jsii_.Get(
		j,
		"snsTopicSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) SnsTopicSinkConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSnsTopicSinkConfiguration {
	var returns *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSnsTopicSinkConfiguration
	_jsii_.Get(
		j,
		"snsTopicSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) SqsQueueSinkConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSqsQueueSinkConfigurationOutputReference {
	var returns ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSqsQueueSinkConfigurationOutputReference
	_jsii_.Get(
		j,
		"sqsQueueSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) SqsQueueSinkConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSqsQueueSinkConfiguration {
	var returns *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSqsQueueSinkConfiguration
	_jsii_.Get(
		j,
		"sqsQueueSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) VoiceAnalyticsProcessorConfiguration() ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsVoiceAnalyticsProcessorConfigurationOutputReference {
	var returns ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsVoiceAnalyticsProcessorConfigurationOutputReference
	_jsii_.Get(
		j,
		"voiceAnalyticsProcessorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) VoiceAnalyticsProcessorConfigurationInput() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsVoiceAnalyticsProcessorConfiguration {
	var returns *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsVoiceAnalyticsProcessorConfiguration
	_jsii_.Get(
		j,
		"voiceAnalyticsProcessorConfigurationInput",
		&returns,
	)
	return returns
}


func NewChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference {
	_init_.Initialize()

	if err := validateNewChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.chimesdkmediapipelinesMediaInsightsPipelineConfiguration.ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference_Override(c ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.chimesdkmediapipelinesMediaInsightsPipelineConfiguration.ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) PutAmazonTranscribeCallAnalyticsProcessorConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfiguration) {
	if err := c.validatePutAmazonTranscribeCallAnalyticsProcessorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAmazonTranscribeCallAnalyticsProcessorConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) PutAmazonTranscribeProcessorConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfiguration) {
	if err := c.validatePutAmazonTranscribeProcessorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAmazonTranscribeProcessorConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) PutKinesisDataStreamSinkConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsKinesisDataStreamSinkConfiguration) {
	if err := c.validatePutKinesisDataStreamSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKinesisDataStreamSinkConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) PutLambdaFunctionSinkConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsLambdaFunctionSinkConfiguration) {
	if err := c.validatePutLambdaFunctionSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLambdaFunctionSinkConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) PutS3RecordingSinkConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfiguration) {
	if err := c.validatePutS3RecordingSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putS3RecordingSinkConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) PutSnsTopicSinkConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSnsTopicSinkConfiguration) {
	if err := c.validatePutSnsTopicSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSnsTopicSinkConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) PutSqsQueueSinkConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsSqsQueueSinkConfiguration) {
	if err := c.validatePutSqsQueueSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSqsQueueSinkConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) PutVoiceAnalyticsProcessorConfiguration(value *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsVoiceAnalyticsProcessorConfiguration) {
	if err := c.validatePutVoiceAnalyticsProcessorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVoiceAnalyticsProcessorConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) ResetAmazonTranscribeCallAnalyticsProcessorConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetAmazonTranscribeCallAnalyticsProcessorConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) ResetAmazonTranscribeProcessorConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetAmazonTranscribeProcessorConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) ResetKinesisDataStreamSinkConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetKinesisDataStreamSinkConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) ResetLambdaFunctionSinkConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetLambdaFunctionSinkConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) ResetS3RecordingSinkConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetS3RecordingSinkConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) ResetSnsTopicSinkConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetSnsTopicSinkConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) ResetSqsQueueSinkConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetSqsQueueSinkConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) ResetVoiceAnalyticsProcessorConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetVoiceAnalyticsProcessorConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

