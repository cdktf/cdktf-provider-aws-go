// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inspector2filter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/inspector2filter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList interface {
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
	Get(index *float64) Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList
type jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewInspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList {
	_init_.Initialize()

	if err := validateNewInspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList{}

	_jsii_.Create(
		"@cdktf/provider-aws.inspector2Filter.Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewInspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList_Override(i Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.inspector2Filter.Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		i,
	)
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := i.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		i,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList) Get(index *float64) Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnOutputReference {
	if err := i.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnOutputReference

	_jsii_.Invoke(
		i,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList) Resolve(context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

