// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/wafv2webaclrulegroupassociation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference interface {
	cdktf.ComplexObject
	Allow() Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseAllowList
	AllowInput() interface{}
	Block() Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseBlockList
	BlockInput() interface{}
	Captcha() Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseCaptchaList
	CaptchaInput() interface{}
	Challenge() Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseChallengeList
	ChallengeInput() interface{}
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
	Count() Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseCountList
	CountInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAllow(value interface{})
	PutBlock(value interface{})
	PutCaptcha(value interface{})
	PutChallenge(value interface{})
	PutCount(value interface{})
	ResetAllow()
	ResetBlock()
	ResetCaptcha()
	ResetChallenge()
	ResetCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference
type jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) Allow() Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseAllowList {
	var returns Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseAllowList
	_jsii_.Get(
		j,
		"allow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) AllowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) Block() Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseBlockList {
	var returns Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseBlockList
	_jsii_.Get(
		j,
		"block",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) BlockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) Captcha() Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseCaptchaList {
	var returns Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseCaptchaList
	_jsii_.Get(
		j,
		"captcha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) CaptchaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captchaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) Challenge() Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseChallengeList {
	var returns Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseChallengeList
	_jsii_.Get(
		j,
		"challenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) ChallengeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"challengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) Count() Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseCountList {
	var returns Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseCountList
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) CountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAclRuleGroupAssociation.Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference_Override(w Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAclRuleGroupAssociation.Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) PutAllow(value interface{}) {
	if err := w.validatePutAllowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllow",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) PutBlock(value interface{}) {
	if err := w.validatePutBlockParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBlock",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) PutCaptcha(value interface{}) {
	if err := w.validatePutCaptchaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCaptcha",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) PutChallenge(value interface{}) {
	if err := w.validatePutChallengeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putChallenge",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) PutCount(value interface{}) {
	if err := w.validatePutCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCount",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) ResetAllow() {
	_jsii_.InvokeVoid(
		w,
		"resetAllow",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) ResetBlock() {
	_jsii_.InvokeVoid(
		w,
		"resetBlock",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) ResetCaptcha() {
	_jsii_.InvokeVoid(
		w,
		"resetCaptcha",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) ResetChallenge() {
	_jsii_.InvokeVoid(
		w,
		"resetChallenge",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		w,
		"resetCount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

