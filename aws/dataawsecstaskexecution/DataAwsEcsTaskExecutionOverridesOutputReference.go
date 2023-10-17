// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsecstaskexecution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/dataawsecstaskexecution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsEcsTaskExecutionOverridesOutputReference interface {
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
	ContainerOverrides() DataAwsEcsTaskExecutionOverridesContainerOverridesList
	ContainerOverridesInput() interface{}
	Cpu() *string
	SetCpu(val *string)
	CpuInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	// Experimental.
	Fqn() *string
	InferenceAcceleratorOverrides() DataAwsEcsTaskExecutionOverridesInferenceAcceleratorOverridesList
	InferenceAcceleratorOverridesInput() interface{}
	InternalValue() *DataAwsEcsTaskExecutionOverrides
	SetInternalValue(val *DataAwsEcsTaskExecutionOverrides)
	Memory() *string
	SetMemory(val *string)
	MemoryInput() *string
	TaskRoleArn() *string
	SetTaskRoleArn(val *string)
	TaskRoleArnInput() *string
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
	PutContainerOverrides(value interface{})
	PutInferenceAcceleratorOverrides(value interface{})
	ResetContainerOverrides()
	ResetCpu()
	ResetExecutionRoleArn()
	ResetInferenceAcceleratorOverrides()
	ResetMemory()
	ResetTaskRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsEcsTaskExecutionOverridesOutputReference
type jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) ContainerOverrides() DataAwsEcsTaskExecutionOverridesContainerOverridesList {
	var returns DataAwsEcsTaskExecutionOverridesContainerOverridesList
	_jsii_.Get(
		j,
		"containerOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) ContainerOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) Cpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) CpuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) InferenceAcceleratorOverrides() DataAwsEcsTaskExecutionOverridesInferenceAcceleratorOverridesList {
	var returns DataAwsEcsTaskExecutionOverridesInferenceAcceleratorOverridesList
	_jsii_.Get(
		j,
		"inferenceAcceleratorOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) InferenceAcceleratorOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceAcceleratorOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) InternalValue() *DataAwsEcsTaskExecutionOverrides {
	var returns *DataAwsEcsTaskExecutionOverrides
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) Memory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) MemoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) TaskRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) TaskRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsEcsTaskExecutionOverridesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsEcsTaskExecutionOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsEcsTaskExecutionOverridesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsEcsTaskExecution.DataAwsEcsTaskExecutionOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsEcsTaskExecutionOverridesOutputReference_Override(d DataAwsEcsTaskExecutionOverridesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsEcsTaskExecution.DataAwsEcsTaskExecutionOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference)SetCpu(val *string) {
	if err := j.validateSetCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference)SetInternalValue(val *DataAwsEcsTaskExecutionOverrides) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference)SetMemory(val *string) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference)SetTaskRoleArn(val *string) {
	if err := j.validateSetTaskRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskRoleArn",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) PutContainerOverrides(value interface{}) {
	if err := d.validatePutContainerOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContainerOverrides",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) PutInferenceAcceleratorOverrides(value interface{}) {
	if err := d.validatePutInferenceAcceleratorOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInferenceAcceleratorOverrides",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) ResetContainerOverrides() {
	_jsii_.InvokeVoid(
		d,
		"resetContainerOverrides",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) ResetCpu() {
	_jsii_.InvokeVoid(
		d,
		"resetCpu",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) ResetExecutionRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetExecutionRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) ResetInferenceAcceleratorOverrides() {
	_jsii_.InvokeVoid(
		d,
		"resetInferenceAcceleratorOverrides",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		d,
		"resetMemory",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) ResetTaskRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetTaskRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecutionOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

