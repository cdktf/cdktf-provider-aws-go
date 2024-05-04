// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxontapvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/fsxontapvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxOntapVolumeAggregateConfigurationOutputReference interface {
	cdktf.ComplexObject
	Aggregates() *[]*string
	SetAggregates(val *[]*string)
	AggregatesInput() *[]*string
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
	ConstituentsPerAggregate() *float64
	SetConstituentsPerAggregate(val *float64)
	ConstituentsPerAggregateInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *FsxOntapVolumeAggregateConfiguration
	SetInternalValue(val *FsxOntapVolumeAggregateConfiguration)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalConstituents() *float64
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
	ResetAggregates()
	ResetConstituentsPerAggregate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FsxOntapVolumeAggregateConfigurationOutputReference
type jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) Aggregates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aggregates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) AggregatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aggregatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) ConstituentsPerAggregate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"constituentsPerAggregate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) ConstituentsPerAggregateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"constituentsPerAggregateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) InternalValue() *FsxOntapVolumeAggregateConfiguration {
	var returns *FsxOntapVolumeAggregateConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) TotalConstituents() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalConstituents",
		&returns,
	)
	return returns
}


func NewFsxOntapVolumeAggregateConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FsxOntapVolumeAggregateConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewFsxOntapVolumeAggregateConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.fsxOntapVolume.FsxOntapVolumeAggregateConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFsxOntapVolumeAggregateConfigurationOutputReference_Override(f FsxOntapVolumeAggregateConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.fsxOntapVolume.FsxOntapVolumeAggregateConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference)SetAggregates(val *[]*string) {
	if err := j.validateSetAggregatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregates",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference)SetConstituentsPerAggregate(val *float64) {
	if err := j.validateSetConstituentsPerAggregateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"constituentsPerAggregate",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference)SetInternalValue(val *FsxOntapVolumeAggregateConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) ResetAggregates() {
	_jsii_.InvokeVoid(
		f,
		"resetAggregates",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) ResetConstituentsPerAggregate() {
	_jsii_.InvokeVoid(
		f,
		"resetConstituentsPerAggregate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (f *jsiiProxy_FsxOntapVolumeAggregateConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

