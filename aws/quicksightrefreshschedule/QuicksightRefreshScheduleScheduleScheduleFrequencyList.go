// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightrefreshschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/quicksightrefreshschedule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightRefreshScheduleScheduleScheduleFrequencyList interface {
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
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightRefreshScheduleScheduleScheduleFrequencyList
type jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewQuicksightRefreshScheduleScheduleScheduleFrequencyList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) QuicksightRefreshScheduleScheduleScheduleFrequencyList {
	_init_.Initialize()

	if err := validateNewQuicksightRefreshScheduleScheduleScheduleFrequencyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightRefreshSchedule.QuicksightRefreshScheduleScheduleScheduleFrequencyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewQuicksightRefreshScheduleScheduleScheduleFrequencyList_Override(q QuicksightRefreshScheduleScheduleScheduleFrequencyList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightRefreshSchedule.QuicksightRefreshScheduleScheduleScheduleFrequencyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList) Get(index *float64) QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference {
	if err := q.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference

	_jsii_.Invoke(
		q,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

