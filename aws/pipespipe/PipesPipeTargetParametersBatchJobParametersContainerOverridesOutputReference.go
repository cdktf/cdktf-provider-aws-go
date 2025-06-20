// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/pipespipe/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference interface {
	cdktf.ComplexObject
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	Environment() PipesPipeTargetParametersBatchJobParametersContainerOverridesEnvironmentList
	EnvironmentInput() interface{}
	// Experimental.
	Fqn() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() *PipesPipeTargetParametersBatchJobParametersContainerOverrides
	SetInternalValue(val *PipesPipeTargetParametersBatchJobParametersContainerOverrides)
	ResourceRequirement() PipesPipeTargetParametersBatchJobParametersContainerOverridesResourceRequirementList
	ResourceRequirementInput() interface{}
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
	PutEnvironment(value interface{})
	PutResourceRequirement(value interface{})
	ResetCommand()
	ResetEnvironment()
	ResetInstanceType()
	ResetResourceRequirement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference
type jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) Environment() PipesPipeTargetParametersBatchJobParametersContainerOverridesEnvironmentList {
	var returns PipesPipeTargetParametersBatchJobParametersContainerOverridesEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) InternalValue() *PipesPipeTargetParametersBatchJobParametersContainerOverrides {
	var returns *PipesPipeTargetParametersBatchJobParametersContainerOverrides
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) ResourceRequirement() PipesPipeTargetParametersBatchJobParametersContainerOverridesResourceRequirementList {
	var returns PipesPipeTargetParametersBatchJobParametersContainerOverridesResourceRequirementList
	_jsii_.Get(
		j,
		"resourceRequirement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) ResourceRequirementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRequirementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewPipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference_Override(p PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference)SetInternalValue(val *PipesPipeTargetParametersBatchJobParametersContainerOverrides) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) PutEnvironment(value interface{}) {
	if err := p.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) PutResourceRequirement(value interface{}) {
	if err := p.validatePutResourceRequirementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putResourceRequirement",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		p,
		"resetCommand",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		p,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		p,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) ResetResourceRequirement() {
	_jsii_.InvokeVoid(
		p,
		"resetResourceRequirement",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipesPipeTargetParametersBatchJobParametersContainerOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

