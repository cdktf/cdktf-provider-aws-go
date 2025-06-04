// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamqueryscheduledquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/timestreamqueryscheduledquery/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	Get(index *float64) TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList
type jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewTimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList {
	_init_.Initialize()

	if err := validateNewTimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList{}

	_jsii_.Create(
		"@cdktf/provider-aws.timestreamqueryScheduledQuery.TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewTimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList_Override(t TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.timestreamqueryScheduledQuery.TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		t,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := t.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		t,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList) Get(index *float64) TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingOutputReference {
	if err := t.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingOutputReference

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMappingList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

