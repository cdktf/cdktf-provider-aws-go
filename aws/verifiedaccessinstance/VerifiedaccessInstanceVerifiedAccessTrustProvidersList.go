// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedaccessinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/verifiedaccessinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VerifiedaccessInstanceVerifiedAccessTrustProvidersList interface {
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
	Get(index *float64) VerifiedaccessInstanceVerifiedAccessTrustProvidersOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VerifiedaccessInstanceVerifiedAccessTrustProvidersList
type jsiiProxy_VerifiedaccessInstanceVerifiedAccessTrustProvidersList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VerifiedaccessInstanceVerifiedAccessTrustProvidersList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceVerifiedAccessTrustProvidersList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceVerifiedAccessTrustProvidersList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceVerifiedAccessTrustProvidersList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceVerifiedAccessTrustProvidersList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVerifiedaccessInstanceVerifiedAccessTrustProvidersList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VerifiedaccessInstanceVerifiedAccessTrustProvidersList {
	_init_.Initialize()

	if err := validateNewVerifiedaccessInstanceVerifiedAccessTrustProvidersListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VerifiedaccessInstanceVerifiedAccessTrustProvidersList{}

	_jsii_.Create(
		"@cdktf/provider-aws.verifiedaccessInstance.VerifiedaccessInstanceVerifiedAccessTrustProvidersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVerifiedaccessInstanceVerifiedAccessTrustProvidersList_Override(v VerifiedaccessInstanceVerifiedAccessTrustProvidersList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.verifiedaccessInstance.VerifiedaccessInstanceVerifiedAccessTrustProvidersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VerifiedaccessInstanceVerifiedAccessTrustProvidersList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessInstanceVerifiedAccessTrustProvidersList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessInstanceVerifiedAccessTrustProvidersList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VerifiedaccessInstanceVerifiedAccessTrustProvidersList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceVerifiedAccessTrustProvidersList) Get(index *float64) VerifiedaccessInstanceVerifiedAccessTrustProvidersOutputReference {
	if err := v.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns VerifiedaccessInstanceVerifiedAccessTrustProvidersOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceVerifiedAccessTrustProvidersList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceVerifiedAccessTrustProvidersList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

