// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubinsight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/securityhubinsight/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList interface {
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
	Get(index *float64) SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList
type jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList {
	_init_.Initialize()

	if err := validateNewSecurityhubInsightFiltersResourceAwsS3BucketOwnerNameListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList{}

	_jsii_.Create(
		"@cdktf/provider-aws.securityhubInsight.SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList_Override(s SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.securityhubInsight.SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := s.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		s,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList) Get(index *float64) SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameOutputReference {
	if err := s.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

