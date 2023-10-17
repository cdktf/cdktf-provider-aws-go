// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/pipespipe/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference interface {
	cdktf.ComplexObject
	BatchSize() *float64
	SetBatchSize(val *float64)
	BatchSizeInput() *float64
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
	ConsumerGroupId() *string
	SetConsumerGroupId(val *string)
	ConsumerGroupIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Credentials() PipesPipeSourceParametersManagedStreamingKafkaParametersCredentialsOutputReference
	CredentialsInput() *PipesPipeSourceParametersManagedStreamingKafkaParametersCredentials
	// Experimental.
	Fqn() *string
	InternalValue() *PipesPipeSourceParametersManagedStreamingKafkaParameters
	SetInternalValue(val *PipesPipeSourceParametersManagedStreamingKafkaParameters)
	MaximumBatchingWindowInSeconds() *float64
	SetMaximumBatchingWindowInSeconds(val *float64)
	MaximumBatchingWindowInSecondsInput() *float64
	StartingPosition() *string
	SetStartingPosition(val *string)
	StartingPositionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TopicName() *string
	SetTopicName(val *string)
	TopicNameInput() *string
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
	PutCredentials(value *PipesPipeSourceParametersManagedStreamingKafkaParametersCredentials)
	ResetBatchSize()
	ResetConsumerGroupId()
	ResetCredentials()
	ResetMaximumBatchingWindowInSeconds()
	ResetStartingPosition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference
type jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) BatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) ConsumerGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) ConsumerGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) Credentials() PipesPipeSourceParametersManagedStreamingKafkaParametersCredentialsOutputReference {
	var returns PipesPipeSourceParametersManagedStreamingKafkaParametersCredentialsOutputReference
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) CredentialsInput() *PipesPipeSourceParametersManagedStreamingKafkaParametersCredentials {
	var returns *PipesPipeSourceParametersManagedStreamingKafkaParametersCredentials
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) InternalValue() *PipesPipeSourceParametersManagedStreamingKafkaParameters {
	var returns *PipesPipeSourceParametersManagedStreamingKafkaParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) MaximumBatchingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) MaximumBatchingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) StartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) StartingPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) TopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) TopicNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicNameInput",
		&returns,
	)
	return returns
}


func NewPipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference {
	_init_.Initialize()

	if err := validateNewPipesPipeSourceParametersManagedStreamingKafkaParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference_Override(p PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference)SetBatchSize(val *float64) {
	if err := j.validateSetBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference)SetConsumerGroupId(val *string) {
	if err := j.validateSetConsumerGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroupId",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference)SetInternalValue(val *PipesPipeSourceParametersManagedStreamingKafkaParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference)SetMaximumBatchingWindowInSeconds(val *float64) {
	if err := j.validateSetMaximumBatchingWindowInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBatchingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference)SetStartingPosition(val *string) {
	if err := j.validateSetStartingPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPosition",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference)SetTopicName(val *string) {
	if err := j.validateSetTopicNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicName",
		val,
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) PutCredentials(value *PipesPipeSourceParametersManagedStreamingKafkaParametersCredentials) {
	if err := p.validatePutCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCredentials",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) ResetBatchSize() {
	_jsii_.InvokeVoid(
		p,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) ResetConsumerGroupId() {
	_jsii_.InvokeVoid(
		p,
		"resetConsumerGroupId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) ResetCredentials() {
	_jsii_.InvokeVoid(
		p,
		"resetCredentials",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) ResetMaximumBatchingWindowInSeconds() {
	_jsii_.InvokeVoid(
		p,
		"resetMaximumBatchingWindowInSeconds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) ResetStartingPosition() {
	_jsii_.InvokeVoid(
		p,
		"resetStartingPosition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersManagedStreamingKafkaParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

