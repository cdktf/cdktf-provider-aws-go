// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawscodeguruprofilerprofilinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/dataawscodeguruprofilerprofilinggroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList
type jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList {
	_init_.Initialize()

	if err := validateNewDataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsCodeguruprofilerProfilingGroup.DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList_Override(d DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsCodeguruprofilerProfilingGroup.DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList) Get(index *float64) DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsCodeguruprofilerProfilingGroupProfilingStatusLatestAggregatedProfileList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

