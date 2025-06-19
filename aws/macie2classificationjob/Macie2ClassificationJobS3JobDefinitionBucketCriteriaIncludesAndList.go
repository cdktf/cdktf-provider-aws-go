// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package macie2classificationjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/macie2classificationjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList interface {
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
	Get(index *float64) Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList
type jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMacie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList {
	_init_.Initialize()

	if err := validateNewMacie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList{}

	_jsii_.Create(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMacie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList_Override(m Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := m.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		m,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList) Get(index *float64) Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndOutputReference {
	if err := m.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesAndList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

