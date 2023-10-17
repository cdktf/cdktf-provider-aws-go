// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recoveryreadinessresourceset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/route53recoveryreadinessresourceset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53RecoveryreadinessResourceSetResourcesList interface {
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
	Get(index *float64) Route53RecoveryreadinessResourceSetResourcesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Route53RecoveryreadinessResourceSetResourcesList
type jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewRoute53RecoveryreadinessResourceSetResourcesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) Route53RecoveryreadinessResourceSetResourcesList {
	_init_.Initialize()

	if err := validateNewRoute53RecoveryreadinessResourceSetResourcesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList{}

	_jsii_.Create(
		"@cdktf/provider-aws.route53RecoveryreadinessResourceSet.Route53RecoveryreadinessResourceSetResourcesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewRoute53RecoveryreadinessResourceSetResourcesList_Override(r Route53RecoveryreadinessResourceSetResourcesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.route53RecoveryreadinessResourceSet.Route53RecoveryreadinessResourceSetResourcesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		r,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList) Get(index *float64) Route53RecoveryreadinessResourceSetResourcesOutputReference {
	if err := r.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns Route53RecoveryreadinessResourceSetResourcesOutputReference

	_jsii_.Invoke(
		r,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

