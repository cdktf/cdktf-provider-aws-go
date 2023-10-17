// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/pipespipe/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference interface {
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
	Cpu() *float64
	SetCpu(val *float64)
	CpuInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Environment() PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironmentList
	EnvironmentFile() PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironmentFileList
	EnvironmentFileInput() interface{}
	EnvironmentInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Memory() *float64
	SetMemory(val *float64)
	MemoryInput() *float64
	MemoryReservation() *float64
	SetMemoryReservation(val *float64)
	MemoryReservationInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	ResourceRequirement() PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideResourceRequirementList
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
	PutEnvironmentFile(value interface{})
	PutResourceRequirement(value interface{})
	ResetCommand()
	ResetCpu()
	ResetEnvironment()
	ResetEnvironmentFile()
	ResetMemory()
	ResetMemoryReservation()
	ResetName()
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

// The jsii proxy struct for PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference
type jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) CpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) Environment() PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironmentList {
	var returns PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) EnvironmentFile() PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironmentFileList {
	var returns PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironmentFileList
	_jsii_.Get(
		j,
		"environmentFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) EnvironmentFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) MemoryReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) ResourceRequirement() PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideResourceRequirementList {
	var returns PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideResourceRequirementList
	_jsii_.Get(
		j,
		"resourceRequirement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) ResourceRequirementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRequirementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference {
	_init_.Initialize()

	if err := validateNewPipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference_Override(p PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference)SetCpu(val *float64) {
	if err := j.validateSetCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference)SetMemoryReservation(val *float64) {
	if err := j.validateSetMemoryReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryReservation",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) PutEnvironment(value interface{}) {
	if err := p.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) PutEnvironmentFile(value interface{}) {
	if err := p.validatePutEnvironmentFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEnvironmentFile",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) PutResourceRequirement(value interface{}) {
	if err := p.validatePutResourceRequirementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putResourceRequirement",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		p,
		"resetCommand",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) ResetCpu() {
	_jsii_.InvokeVoid(
		p,
		"resetCpu",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		p,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) ResetEnvironmentFile() {
	_jsii_.InvokeVoid(
		p,
		"resetEnvironmentFile",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		p,
		"resetMemory",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) ResetMemoryReservation() {
	_jsii_.InvokeVoid(
		p,
		"resetMemoryReservation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) ResetResourceRequirement() {
	_jsii_.InvokeVoid(
		p,
		"resetResourceRequirement",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipesPipeTargetParametersEcsTaskParametersOverridesContainerOverrideOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

