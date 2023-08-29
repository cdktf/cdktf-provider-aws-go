// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/pipespipe/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipesPipeTargetParametersOutputReference interface {
	cdktf.ComplexObject
	BatchJobParameters() PipesPipeTargetParametersBatchJobParametersOutputReference
	BatchJobParametersInput() *PipesPipeTargetParametersBatchJobParameters
	CloudwatchLogsParameters() PipesPipeTargetParametersCloudwatchLogsParametersOutputReference
	CloudwatchLogsParametersInput() *PipesPipeTargetParametersCloudwatchLogsParameters
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
	EcsTaskParameters() PipesPipeTargetParametersEcsTaskParametersOutputReference
	EcsTaskParametersInput() *PipesPipeTargetParametersEcsTaskParameters
	EventbridgeEventBusParameters() PipesPipeTargetParametersEventbridgeEventBusParametersOutputReference
	EventbridgeEventBusParametersInput() *PipesPipeTargetParametersEventbridgeEventBusParameters
	// Experimental.
	Fqn() *string
	HttpParameters() PipesPipeTargetParametersHttpParametersOutputReference
	HttpParametersInput() *PipesPipeTargetParametersHttpParameters
	InputTemplate() *string
	SetInputTemplate(val *string)
	InputTemplateInput() *string
	InternalValue() *PipesPipeTargetParameters
	SetInternalValue(val *PipesPipeTargetParameters)
	KinesisStreamParameters() PipesPipeTargetParametersKinesisStreamParametersOutputReference
	KinesisStreamParametersInput() *PipesPipeTargetParametersKinesisStreamParameters
	LambdaFunctionParameters() PipesPipeTargetParametersLambdaFunctionParametersOutputReference
	LambdaFunctionParametersInput() *PipesPipeTargetParametersLambdaFunctionParameters
	RedshiftDataParameters() PipesPipeTargetParametersRedshiftDataParametersOutputReference
	RedshiftDataParametersInput() *PipesPipeTargetParametersRedshiftDataParameters
	SagemakerPipelineParameters() PipesPipeTargetParametersSagemakerPipelineParametersOutputReference
	SagemakerPipelineParametersInput() *PipesPipeTargetParametersSagemakerPipelineParameters
	SqsQueueParameters() PipesPipeTargetParametersSqsQueueParametersOutputReference
	SqsQueueParametersInput() *PipesPipeTargetParametersSqsQueueParameters
	StepFunctionStateMachineParameters() PipesPipeTargetParametersStepFunctionStateMachineParametersOutputReference
	StepFunctionStateMachineParametersInput() *PipesPipeTargetParametersStepFunctionStateMachineParameters
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
	PutBatchJobParameters(value *PipesPipeTargetParametersBatchJobParameters)
	PutCloudwatchLogsParameters(value *PipesPipeTargetParametersCloudwatchLogsParameters)
	PutEcsTaskParameters(value *PipesPipeTargetParametersEcsTaskParameters)
	PutEventbridgeEventBusParameters(value *PipesPipeTargetParametersEventbridgeEventBusParameters)
	PutHttpParameters(value *PipesPipeTargetParametersHttpParameters)
	PutKinesisStreamParameters(value *PipesPipeTargetParametersKinesisStreamParameters)
	PutLambdaFunctionParameters(value *PipesPipeTargetParametersLambdaFunctionParameters)
	PutRedshiftDataParameters(value *PipesPipeTargetParametersRedshiftDataParameters)
	PutSagemakerPipelineParameters(value *PipesPipeTargetParametersSagemakerPipelineParameters)
	PutSqsQueueParameters(value *PipesPipeTargetParametersSqsQueueParameters)
	PutStepFunctionStateMachineParameters(value *PipesPipeTargetParametersStepFunctionStateMachineParameters)
	ResetBatchJobParameters()
	ResetCloudwatchLogsParameters()
	ResetEcsTaskParameters()
	ResetEventbridgeEventBusParameters()
	ResetHttpParameters()
	ResetInputTemplate()
	ResetKinesisStreamParameters()
	ResetLambdaFunctionParameters()
	ResetRedshiftDataParameters()
	ResetSagemakerPipelineParameters()
	ResetSqsQueueParameters()
	ResetStepFunctionStateMachineParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipesPipeTargetParametersOutputReference
type jsiiProxy_PipesPipeTargetParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) BatchJobParameters() PipesPipeTargetParametersBatchJobParametersOutputReference {
	var returns PipesPipeTargetParametersBatchJobParametersOutputReference
	_jsii_.Get(
		j,
		"batchJobParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) BatchJobParametersInput() *PipesPipeTargetParametersBatchJobParameters {
	var returns *PipesPipeTargetParametersBatchJobParameters
	_jsii_.Get(
		j,
		"batchJobParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) CloudwatchLogsParameters() PipesPipeTargetParametersCloudwatchLogsParametersOutputReference {
	var returns PipesPipeTargetParametersCloudwatchLogsParametersOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) CloudwatchLogsParametersInput() *PipesPipeTargetParametersCloudwatchLogsParameters {
	var returns *PipesPipeTargetParametersCloudwatchLogsParameters
	_jsii_.Get(
		j,
		"cloudwatchLogsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) EcsTaskParameters() PipesPipeTargetParametersEcsTaskParametersOutputReference {
	var returns PipesPipeTargetParametersEcsTaskParametersOutputReference
	_jsii_.Get(
		j,
		"ecsTaskParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) EcsTaskParametersInput() *PipesPipeTargetParametersEcsTaskParameters {
	var returns *PipesPipeTargetParametersEcsTaskParameters
	_jsii_.Get(
		j,
		"ecsTaskParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) EventbridgeEventBusParameters() PipesPipeTargetParametersEventbridgeEventBusParametersOutputReference {
	var returns PipesPipeTargetParametersEventbridgeEventBusParametersOutputReference
	_jsii_.Get(
		j,
		"eventbridgeEventBusParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) EventbridgeEventBusParametersInput() *PipesPipeTargetParametersEventbridgeEventBusParameters {
	var returns *PipesPipeTargetParametersEventbridgeEventBusParameters
	_jsii_.Get(
		j,
		"eventbridgeEventBusParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) HttpParameters() PipesPipeTargetParametersHttpParametersOutputReference {
	var returns PipesPipeTargetParametersHttpParametersOutputReference
	_jsii_.Get(
		j,
		"httpParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) HttpParametersInput() *PipesPipeTargetParametersHttpParameters {
	var returns *PipesPipeTargetParametersHttpParameters
	_jsii_.Get(
		j,
		"httpParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) InputTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) InputTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) InternalValue() *PipesPipeTargetParameters {
	var returns *PipesPipeTargetParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) KinesisStreamParameters() PipesPipeTargetParametersKinesisStreamParametersOutputReference {
	var returns PipesPipeTargetParametersKinesisStreamParametersOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) KinesisStreamParametersInput() *PipesPipeTargetParametersKinesisStreamParameters {
	var returns *PipesPipeTargetParametersKinesisStreamParameters
	_jsii_.Get(
		j,
		"kinesisStreamParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) LambdaFunctionParameters() PipesPipeTargetParametersLambdaFunctionParametersOutputReference {
	var returns PipesPipeTargetParametersLambdaFunctionParametersOutputReference
	_jsii_.Get(
		j,
		"lambdaFunctionParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) LambdaFunctionParametersInput() *PipesPipeTargetParametersLambdaFunctionParameters {
	var returns *PipesPipeTargetParametersLambdaFunctionParameters
	_jsii_.Get(
		j,
		"lambdaFunctionParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) RedshiftDataParameters() PipesPipeTargetParametersRedshiftDataParametersOutputReference {
	var returns PipesPipeTargetParametersRedshiftDataParametersOutputReference
	_jsii_.Get(
		j,
		"redshiftDataParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) RedshiftDataParametersInput() *PipesPipeTargetParametersRedshiftDataParameters {
	var returns *PipesPipeTargetParametersRedshiftDataParameters
	_jsii_.Get(
		j,
		"redshiftDataParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) SagemakerPipelineParameters() PipesPipeTargetParametersSagemakerPipelineParametersOutputReference {
	var returns PipesPipeTargetParametersSagemakerPipelineParametersOutputReference
	_jsii_.Get(
		j,
		"sagemakerPipelineParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) SagemakerPipelineParametersInput() *PipesPipeTargetParametersSagemakerPipelineParameters {
	var returns *PipesPipeTargetParametersSagemakerPipelineParameters
	_jsii_.Get(
		j,
		"sagemakerPipelineParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) SqsQueueParameters() PipesPipeTargetParametersSqsQueueParametersOutputReference {
	var returns PipesPipeTargetParametersSqsQueueParametersOutputReference
	_jsii_.Get(
		j,
		"sqsQueueParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) SqsQueueParametersInput() *PipesPipeTargetParametersSqsQueueParameters {
	var returns *PipesPipeTargetParametersSqsQueueParameters
	_jsii_.Get(
		j,
		"sqsQueueParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) StepFunctionStateMachineParameters() PipesPipeTargetParametersStepFunctionStateMachineParametersOutputReference {
	var returns PipesPipeTargetParametersStepFunctionStateMachineParametersOutputReference
	_jsii_.Get(
		j,
		"stepFunctionStateMachineParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) StepFunctionStateMachineParametersInput() *PipesPipeTargetParametersStepFunctionStateMachineParameters {
	var returns *PipesPipeTargetParametersStepFunctionStateMachineParameters
	_jsii_.Get(
		j,
		"stepFunctionStateMachineParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipesPipeTargetParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PipesPipeTargetParametersOutputReference {
	_init_.Initialize()

	if err := validateNewPipesPipeTargetParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipesPipeTargetParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeTargetParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipesPipeTargetParametersOutputReference_Override(p PipesPipeTargetParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeTargetParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference)SetInputTemplate(val *string) {
	if err := j.validateSetInputTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputTemplate",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference)SetInternalValue(val *PipesPipeTargetParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) PutBatchJobParameters(value *PipesPipeTargetParametersBatchJobParameters) {
	if err := p.validatePutBatchJobParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putBatchJobParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) PutCloudwatchLogsParameters(value *PipesPipeTargetParametersCloudwatchLogsParameters) {
	if err := p.validatePutCloudwatchLogsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCloudwatchLogsParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) PutEcsTaskParameters(value *PipesPipeTargetParametersEcsTaskParameters) {
	if err := p.validatePutEcsTaskParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEcsTaskParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) PutEventbridgeEventBusParameters(value *PipesPipeTargetParametersEventbridgeEventBusParameters) {
	if err := p.validatePutEventbridgeEventBusParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventbridgeEventBusParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) PutHttpParameters(value *PipesPipeTargetParametersHttpParameters) {
	if err := p.validatePutHttpParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHttpParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) PutKinesisStreamParameters(value *PipesPipeTargetParametersKinesisStreamParameters) {
	if err := p.validatePutKinesisStreamParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putKinesisStreamParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) PutLambdaFunctionParameters(value *PipesPipeTargetParametersLambdaFunctionParameters) {
	if err := p.validatePutLambdaFunctionParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLambdaFunctionParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) PutRedshiftDataParameters(value *PipesPipeTargetParametersRedshiftDataParameters) {
	if err := p.validatePutRedshiftDataParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRedshiftDataParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) PutSagemakerPipelineParameters(value *PipesPipeTargetParametersSagemakerPipelineParameters) {
	if err := p.validatePutSagemakerPipelineParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSagemakerPipelineParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) PutSqsQueueParameters(value *PipesPipeTargetParametersSqsQueueParameters) {
	if err := p.validatePutSqsQueueParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSqsQueueParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) PutStepFunctionStateMachineParameters(value *PipesPipeTargetParametersStepFunctionStateMachineParameters) {
	if err := p.validatePutStepFunctionStateMachineParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putStepFunctionStateMachineParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) ResetBatchJobParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetBatchJobParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) ResetCloudwatchLogsParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetCloudwatchLogsParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) ResetEcsTaskParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetEcsTaskParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) ResetEventbridgeEventBusParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetEventbridgeEventBusParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) ResetHttpParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetHttpParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) ResetInputTemplate() {
	_jsii_.InvokeVoid(
		p,
		"resetInputTemplate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) ResetKinesisStreamParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetKinesisStreamParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) ResetLambdaFunctionParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetLambdaFunctionParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) ResetRedshiftDataParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetRedshiftDataParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) ResetSagemakerPipelineParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetSagemakerPipelineParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) ResetSqsQueueParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetSqsQueueParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) ResetStepFunctionStateMachineParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetStepFunctionStateMachineParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipesPipeTargetParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

