// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/kinesisanalyticsapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList interface {
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
	Get(index *float64) KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList
type jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewKinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList {
	_init_.Initialize()

	if err := validateNewKinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisAnalyticsApplication.KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewKinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList_Override(k KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisAnalyticsApplication.KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		k,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList) Get(index *float64) KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsOutputReference {
	if err := k.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsOutputReference

	_jsii_.Invoke(
		k,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

