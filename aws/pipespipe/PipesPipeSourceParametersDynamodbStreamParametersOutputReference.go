// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/pipespipe/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipesPipeSourceParametersDynamodbStreamParametersOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeadLetterConfig() PipesPipeSourceParametersDynamodbStreamParametersDeadLetterConfigOutputReference
	DeadLetterConfigInput() *PipesPipeSourceParametersDynamodbStreamParametersDeadLetterConfig
	// Experimental.
	Fqn() *string
	InternalValue() *PipesPipeSourceParametersDynamodbStreamParameters
	SetInternalValue(val *PipesPipeSourceParametersDynamodbStreamParameters)
	MaximumBatchingWindowInSeconds() *float64
	SetMaximumBatchingWindowInSeconds(val *float64)
	MaximumBatchingWindowInSecondsInput() *float64
	MaximumRecordAgeInSeconds() *float64
	SetMaximumRecordAgeInSeconds(val *float64)
	MaximumRecordAgeInSecondsInput() *float64
	MaximumRetryAttempts() *float64
	SetMaximumRetryAttempts(val *float64)
	MaximumRetryAttemptsInput() *float64
	OnPartialBatchItemFailure() *string
	SetOnPartialBatchItemFailure(val *string)
	OnPartialBatchItemFailureInput() *string
	ParallelizationFactor() *float64
	SetParallelizationFactor(val *float64)
	ParallelizationFactorInput() *float64
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
	PutDeadLetterConfig(value *PipesPipeSourceParametersDynamodbStreamParametersDeadLetterConfig)
	ResetBatchSize()
	ResetDeadLetterConfig()
	ResetMaximumBatchingWindowInSeconds()
	ResetMaximumRecordAgeInSeconds()
	ResetMaximumRetryAttempts()
	ResetOnPartialBatchItemFailure()
	ResetParallelizationFactor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipesPipeSourceParametersDynamodbStreamParametersOutputReference
type jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) BatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) DeadLetterConfig() PipesPipeSourceParametersDynamodbStreamParametersDeadLetterConfigOutputReference {
	var returns PipesPipeSourceParametersDynamodbStreamParametersDeadLetterConfigOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) DeadLetterConfigInput() *PipesPipeSourceParametersDynamodbStreamParametersDeadLetterConfig {
	var returns *PipesPipeSourceParametersDynamodbStreamParametersDeadLetterConfig
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) InternalValue() *PipesPipeSourceParametersDynamodbStreamParameters {
	var returns *PipesPipeSourceParametersDynamodbStreamParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) MaximumBatchingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) MaximumBatchingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) MaximumRecordAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) MaximumRecordAgeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) MaximumRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) MaximumRetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) OnPartialBatchItemFailure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onPartialBatchItemFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) OnPartialBatchItemFailureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onPartialBatchItemFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) ParallelizationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) ParallelizationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) StartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) StartingPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipesPipeSourceParametersDynamodbStreamParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PipesPipeSourceParametersDynamodbStreamParametersOutputReference {
	_init_.Initialize()

	if err := validateNewPipesPipeSourceParametersDynamodbStreamParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeSourceParametersDynamodbStreamParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipesPipeSourceParametersDynamodbStreamParametersOutputReference_Override(p PipesPipeSourceParametersDynamodbStreamParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeSourceParametersDynamodbStreamParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference)SetBatchSize(val *float64) {
	if err := j.validateSetBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference)SetInternalValue(val *PipesPipeSourceParametersDynamodbStreamParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference)SetMaximumBatchingWindowInSeconds(val *float64) {
	if err := j.validateSetMaximumBatchingWindowInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBatchingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference)SetMaximumRecordAgeInSeconds(val *float64) {
	if err := j.validateSetMaximumRecordAgeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRecordAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference)SetMaximumRetryAttempts(val *float64) {
	if err := j.validateSetMaximumRetryAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference)SetOnPartialBatchItemFailure(val *string) {
	if err := j.validateSetOnPartialBatchItemFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onPartialBatchItemFailure",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference)SetParallelizationFactor(val *float64) {
	if err := j.validateSetParallelizationFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelizationFactor",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference)SetStartingPosition(val *string) {
	if err := j.validateSetStartingPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPosition",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) PutDeadLetterConfig(value *PipesPipeSourceParametersDynamodbStreamParametersDeadLetterConfig) {
	if err := p.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) ResetBatchSize() {
	_jsii_.InvokeVoid(
		p,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		p,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) ResetMaximumBatchingWindowInSeconds() {
	_jsii_.InvokeVoid(
		p,
		"resetMaximumBatchingWindowInSeconds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) ResetMaximumRecordAgeInSeconds() {
	_jsii_.InvokeVoid(
		p,
		"resetMaximumRecordAgeInSeconds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) ResetMaximumRetryAttempts() {
	_jsii_.InvokeVoid(
		p,
		"resetMaximumRetryAttempts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) ResetOnPartialBatchItemFailure() {
	_jsii_.InvokeVoid(
		p,
		"resetOnPartialBatchItemFailure",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) ResetParallelizationFactor() {
	_jsii_.InvokeVoid(
		p,
		"resetParallelizationFactor",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipesPipeSourceParametersDynamodbStreamParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

