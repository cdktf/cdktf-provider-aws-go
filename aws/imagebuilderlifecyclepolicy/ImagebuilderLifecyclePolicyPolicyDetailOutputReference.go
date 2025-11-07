// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderlifecyclepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/imagebuilderlifecyclepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImagebuilderLifecyclePolicyPolicyDetailOutputReference interface {
	cdktf.ComplexObject
	Action() ImagebuilderLifecyclePolicyPolicyDetailActionList
	ActionInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExclusionRules() ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesList
	ExclusionRulesInput() interface{}
	Filter() ImagebuilderLifecyclePolicyPolicyDetailFilterList
	FilterInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutAction(value interface{})
	PutExclusionRules(value interface{})
	PutFilter(value interface{})
	ResetAction()
	ResetExclusionRules()
	ResetFilter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ImagebuilderLifecyclePolicyPolicyDetailOutputReference
type jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) Action() ImagebuilderLifecyclePolicyPolicyDetailActionList {
	var returns ImagebuilderLifecyclePolicyPolicyDetailActionList
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) ActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) ExclusionRules() ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesList {
	var returns ImagebuilderLifecyclePolicyPolicyDetailExclusionRulesList
	_jsii_.Get(
		j,
		"exclusionRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) ExclusionRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exclusionRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) Filter() ImagebuilderLifecyclePolicyPolicyDetailFilterList {
	var returns ImagebuilderLifecyclePolicyPolicyDetailFilterList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewImagebuilderLifecyclePolicyPolicyDetailOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ImagebuilderLifecyclePolicyPolicyDetailOutputReference {
	_init_.Initialize()

	if err := validateNewImagebuilderLifecyclePolicyPolicyDetailOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.imagebuilderLifecyclePolicy.ImagebuilderLifecyclePolicyPolicyDetailOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewImagebuilderLifecyclePolicyPolicyDetailOutputReference_Override(i ImagebuilderLifecyclePolicyPolicyDetailOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.imagebuilderLifecyclePolicy.ImagebuilderLifecyclePolicyPolicyDetailOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) PutAction(value interface{}) {
	if err := i.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAction",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) PutExclusionRules(value interface{}) {
	if err := i.validatePutExclusionRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putExclusionRules",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) PutFilter(value interface{}) {
	if err := i.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFilter",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		i,
		"resetAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) ResetExclusionRules() {
	_jsii_.InvokeVoid(
		i,
		"resetExclusionRules",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		i,
		"resetFilter",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_ImagebuilderLifecyclePolicyPolicyDetailOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

