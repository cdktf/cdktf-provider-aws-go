// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/quicksightdashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList interface {
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
	Get(index *float64) QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList
type jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewQuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList {
	_init_.Initialize()

	if err := validateNewQuicksightDashboardSourceEntitySourceTemplateDataSetReferencesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDashboard.QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewQuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList_Override(q QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDashboard.QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (q *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := q.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		q,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList) Get(index *float64) QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesOutputReference {
	if err := q.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesOutputReference

	_jsii_.Invoke(
		q,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightDashboardSourceEntitySourceTemplateDataSetReferencesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

