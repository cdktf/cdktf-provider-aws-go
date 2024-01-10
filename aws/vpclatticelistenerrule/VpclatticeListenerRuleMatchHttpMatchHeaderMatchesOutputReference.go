// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpclatticelistenerrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/vpclatticelistenerrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference interface {
	cdktf.ComplexObject
	CaseSensitive() interface{}
	SetCaseSensitive(val interface{})
	CaseSensitiveInput() interface{}
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Match() VpclatticeListenerRuleMatchHttpMatchHeaderMatchesMatchOutputReference
	MatchInput() *VpclatticeListenerRuleMatchHttpMatchHeaderMatchesMatch
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutMatch(value *VpclatticeListenerRuleMatchHttpMatchHeaderMatchesMatch)
	ResetCaseSensitive()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference
type jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) CaseSensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) CaseSensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) Match() VpclatticeListenerRuleMatchHttpMatchHeaderMatchesMatchOutputReference {
	var returns VpclatticeListenerRuleMatchHttpMatchHeaderMatchesMatchOutputReference
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) MatchInput() *VpclatticeListenerRuleMatchHttpMatchHeaderMatchesMatch {
	var returns *VpclatticeListenerRuleMatchHttpMatchHeaderMatchesMatch
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference {
	_init_.Initialize()

	if err := validateNewVpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.vpclatticeListenerRule.VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference_Override(v VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.vpclatticeListenerRule.VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference)SetCaseSensitive(val interface{}) {
	if err := j.validateSetCaseSensitiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caseSensitive",
		val,
	)
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) PutMatch(value *VpclatticeListenerRuleMatchHttpMatchHeaderMatchesMatch) {
	if err := v.validatePutMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putMatch",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) ResetCaseSensitive() {
	_jsii_.InvokeVoid(
		v,
		"resetCaseSensitive",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchHeaderMatchesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

