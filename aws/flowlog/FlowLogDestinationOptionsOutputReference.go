// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package flowlog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/flowlog/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FlowLogDestinationOptionsOutputReference interface {
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
	FileFormat() *string
	SetFileFormat(val *string)
	FileFormatInput() *string
	// Experimental.
	Fqn() *string
	HiveCompatiblePartitions() interface{}
	SetHiveCompatiblePartitions(val interface{})
	HiveCompatiblePartitionsInput() interface{}
	InternalValue() *FlowLogDestinationOptions
	SetInternalValue(val *FlowLogDestinationOptions)
	PerHourPartition() interface{}
	SetPerHourPartition(val interface{})
	PerHourPartitionInput() interface{}
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
	ResetFileFormat()
	ResetHiveCompatiblePartitions()
	ResetPerHourPartition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FlowLogDestinationOptionsOutputReference
type jsiiProxy_FlowLogDestinationOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference) FileFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference) FileFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference) HiveCompatiblePartitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hiveCompatiblePartitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference) HiveCompatiblePartitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hiveCompatiblePartitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference) InternalValue() *FlowLogDestinationOptions {
	var returns *FlowLogDestinationOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference) PerHourPartition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"perHourPartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference) PerHourPartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"perHourPartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFlowLogDestinationOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FlowLogDestinationOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewFlowLogDestinationOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FlowLogDestinationOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.flowLog.FlowLogDestinationOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFlowLogDestinationOptionsOutputReference_Override(f FlowLogDestinationOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.flowLog.FlowLogDestinationOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference)SetFileFormat(val *string) {
	if err := j.validateSetFileFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileFormat",
		val,
	)
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference)SetHiveCompatiblePartitions(val interface{}) {
	if err := j.validateSetHiveCompatiblePartitionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiveCompatiblePartitions",
		val,
	)
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference)SetInternalValue(val *FlowLogDestinationOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference)SetPerHourPartition(val interface{}) {
	if err := j.validateSetPerHourPartitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perHourPartition",
		val,
	)
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FlowLogDestinationOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) ResetFileFormat() {
	_jsii_.InvokeVoid(
		f,
		"resetFileFormat",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) ResetHiveCompatiblePartitions() {
	_jsii_.InvokeVoid(
		f,
		"resetHiveCompatiblePartitions",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) ResetPerHourPartition() {
	_jsii_.InvokeVoid(
		f,
		"resetPerHourPartition",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestinationOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

