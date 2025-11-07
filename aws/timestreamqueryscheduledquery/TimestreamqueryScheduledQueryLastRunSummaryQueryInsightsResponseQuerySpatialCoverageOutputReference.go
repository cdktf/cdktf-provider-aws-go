// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamqueryscheduledquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/timestreamqueryscheduledquery/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Max() TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxList
	MaxInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutMax(value interface{})
	ResetMax()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference
type jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) Max() TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxList {
	var returns TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxList
	_jsii_.Get(
		j,
		"max",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) MaxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference {
	_init_.Initialize()

	if err := validateNewTimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.timestreamqueryScheduledQuery.TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewTimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference_Override(t TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.timestreamqueryScheduledQuery.TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) PutMax(value interface{}) {
	if err := t.validatePutMaxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMax",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) ResetMax() {
	_jsii_.InvokeVoid(
		t,
		"resetMax",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (t *jsiiProxy_TimestreamqueryScheduledQueryLastRunSummaryQueryInsightsResponseQuerySpatialCoverageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

