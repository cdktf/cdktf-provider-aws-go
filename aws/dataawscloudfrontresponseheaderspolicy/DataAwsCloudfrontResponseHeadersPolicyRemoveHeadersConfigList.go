// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawscloudfrontresponseheaderspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/dataawscloudfrontresponseheaderspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList interface {
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
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList {
	_init_.Initialize()

	if err := validateNewDataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsCloudfrontResponseHeadersPolicy.DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList_Override(d DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsCloudfrontResponseHeadersPolicy.DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList) Get(index *float64) DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyRemoveHeadersConfigList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

