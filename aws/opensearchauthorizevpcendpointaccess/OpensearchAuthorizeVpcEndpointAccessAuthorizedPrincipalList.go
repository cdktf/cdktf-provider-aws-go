// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchauthorizevpcendpointaccess

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/opensearchauthorizevpcendpointaccess/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList interface {
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
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList
type jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewOpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList {
	_init_.Initialize()

	if err := validateNewOpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList{}

	_jsii_.Create(
		"@cdktf/provider-aws.opensearchAuthorizeVpcEndpointAccess.OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewOpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList_Override(o OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opensearchAuthorizeVpcEndpointAccess.OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := o.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		o,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList) Get(index *float64) OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalOutputReference {
	if err := o.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchAuthorizeVpcEndpointAccessAuthorizedPrincipalList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

