// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/pipespipe/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipesPipeTargetParametersBatchJobParametersOutputReference interface {
	cdktf.ComplexObject
	ArrayProperties() PipesPipeTargetParametersBatchJobParametersArrayPropertiesOutputReference
	ArrayPropertiesInput() *PipesPipeTargetParametersBatchJobParametersArrayProperties
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
	ContainerOverrides() PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference
	ContainerOverridesInput() *PipesPipeTargetParametersBatchJobParametersContainerOverrides
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DependsOn() PipesPipeTargetParametersBatchJobParametersDependsOnList
	DependsOnInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *PipesPipeTargetParametersBatchJobParameters
	SetInternalValue(val *PipesPipeTargetParametersBatchJobParameters)
	JobDefinition() *string
	SetJobDefinition(val *string)
	JobDefinitionInput() *string
	JobName() *string
	SetJobName(val *string)
	JobNameInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	RetryStrategy() PipesPipeTargetParametersBatchJobParametersRetryStrategyOutputReference
	RetryStrategyInput() *PipesPipeTargetParametersBatchJobParametersRetryStrategy
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
	PutArrayProperties(value *PipesPipeTargetParametersBatchJobParametersArrayProperties)
	PutContainerOverrides(value *PipesPipeTargetParametersBatchJobParametersContainerOverrides)
	PutDependsOn(value interface{})
	PutRetryStrategy(value *PipesPipeTargetParametersBatchJobParametersRetryStrategy)
	ResetArrayProperties()
	ResetContainerOverrides()
	ResetDependsOn()
	ResetParameters()
	ResetRetryStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipesPipeTargetParametersBatchJobParametersOutputReference
type jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) ArrayProperties() PipesPipeTargetParametersBatchJobParametersArrayPropertiesOutputReference {
	var returns PipesPipeTargetParametersBatchJobParametersArrayPropertiesOutputReference
	_jsii_.Get(
		j,
		"arrayProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) ArrayPropertiesInput() *PipesPipeTargetParametersBatchJobParametersArrayProperties {
	var returns *PipesPipeTargetParametersBatchJobParametersArrayProperties
	_jsii_.Get(
		j,
		"arrayPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) ContainerOverrides() PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference {
	var returns PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference
	_jsii_.Get(
		j,
		"containerOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) ContainerOverridesInput() *PipesPipeTargetParametersBatchJobParametersContainerOverrides {
	var returns *PipesPipeTargetParametersBatchJobParametersContainerOverrides
	_jsii_.Get(
		j,
		"containerOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) DependsOn() PipesPipeTargetParametersBatchJobParametersDependsOnList {
	var returns PipesPipeTargetParametersBatchJobParametersDependsOnList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) DependsOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) InternalValue() *PipesPipeTargetParametersBatchJobParameters {
	var returns *PipesPipeTargetParametersBatchJobParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) JobDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) JobDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) JobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) RetryStrategy() PipesPipeTargetParametersBatchJobParametersRetryStrategyOutputReference {
	var returns PipesPipeTargetParametersBatchJobParametersRetryStrategyOutputReference
	_jsii_.Get(
		j,
		"retryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) RetryStrategyInput() *PipesPipeTargetParametersBatchJobParametersRetryStrategy {
	var returns *PipesPipeTargetParametersBatchJobParametersRetryStrategy
	_jsii_.Get(
		j,
		"retryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipesPipeTargetParametersBatchJobParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PipesPipeTargetParametersBatchJobParametersOutputReference {
	_init_.Initialize()

	if err := validateNewPipesPipeTargetParametersBatchJobParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeTargetParametersBatchJobParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipesPipeTargetParametersBatchJobParametersOutputReference_Override(p PipesPipeTargetParametersBatchJobParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeTargetParametersBatchJobParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference)SetInternalValue(val *PipesPipeTargetParametersBatchJobParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference)SetJobDefinition(val *string) {
	if err := j.validateSetJobDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobDefinition",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference)SetJobName(val *string) {
	if err := j.validateSetJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobName",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) PutArrayProperties(value *PipesPipeTargetParametersBatchJobParametersArrayProperties) {
	if err := p.validatePutArrayPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putArrayProperties",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) PutContainerOverrides(value *PipesPipeTargetParametersBatchJobParametersContainerOverrides) {
	if err := p.validatePutContainerOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putContainerOverrides",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) PutDependsOn(value interface{}) {
	if err := p.validatePutDependsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDependsOn",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) PutRetryStrategy(value *PipesPipeTargetParametersBatchJobParametersRetryStrategy) {
	if err := p.validatePutRetryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRetryStrategy",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) ResetArrayProperties() {
	_jsii_.InvokeVoid(
		p,
		"resetArrayProperties",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) ResetContainerOverrides() {
	_jsii_.InvokeVoid(
		p,
		"resetContainerOverrides",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		p,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) ResetRetryStrategy() {
	_jsii_.InvokeVoid(
		p,
		"resetRetryStrategy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

