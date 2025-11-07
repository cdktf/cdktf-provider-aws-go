// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recordsexclusive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/route53recordsexclusive/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList interface {
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
	Get(index *float64) Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList
type jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewRoute53RecordsExclusiveResourceRecordSetCidrRoutingConfigList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList {
	_init_.Initialize()

	if err := validateNewRoute53RecordsExclusiveResourceRecordSetCidrRoutingConfigListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList{}

	_jsii_.Create(
		"@cdktf/provider-aws.route53RecordsExclusive.Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewRoute53RecordsExclusiveResourceRecordSetCidrRoutingConfigList_Override(r Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.route53RecordsExclusive.Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		r,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := r.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		r,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList) Get(index *float64) Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigOutputReference {
	if err := r.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigOutputReference

	_jsii_.Invoke(
		r,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList) Resolve(context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

