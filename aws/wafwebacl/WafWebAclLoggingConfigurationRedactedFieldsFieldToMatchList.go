// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafwebacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/wafwebacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList interface {
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
	Get(index *float64) WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList
type jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList {
	_init_.Initialize()

	if err := validateNewWafWebAclLoggingConfigurationRedactedFieldsFieldToMatchListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafWebAcl.WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList_Override(w WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafWebAcl.WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := w.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		w,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList) Get(index *float64) WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchOutputReference {
	if err := w.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafWebAclLoggingConfigurationRedactedFieldsFieldToMatchList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

