// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamqueryscheduledquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/timestreamqueryscheduledquery/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList interface {
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
	Get(index *float64) TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList
type jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewTimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList {
	_init_.Initialize()

	if err := validateNewTimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList{}

	_jsii_.Create(
		"@cdktf/provider-aws.timestreamqueryScheduledQuery.TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewTimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList_Override(t TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.timestreamqueryScheduledQuery.TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		t,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
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

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList) Get(index *float64) TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxOutputReference {
	if err := t.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxOutputReference

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList) Resolve(context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQueryTemporalRangeMaxList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

