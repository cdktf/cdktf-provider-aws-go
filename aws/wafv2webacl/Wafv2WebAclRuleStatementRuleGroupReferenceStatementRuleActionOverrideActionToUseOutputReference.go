// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference interface {
	cdktf.ComplexObject
	Allow() Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseAllowOutputReference
	AllowInput() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseAllow
	Block() Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlockOutputReference
	BlockInput() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlock
	Captcha() Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCaptchaOutputReference
	CaptchaInput() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCaptcha
	Challenge() Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference
	ChallengeInput() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallenge
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
	Count() Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCountOutputReference
	CountInput() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCount
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUse
	SetInternalValue(val *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUse)
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
	PutAllow(value *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseAllow)
	PutBlock(value *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlock)
	PutCaptcha(value *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCaptcha)
	PutChallenge(value *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallenge)
	PutCount(value *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCount)
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

// The jsii proxy struct for Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference
type jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) Allow() Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseAllowOutputReference {
	var returns Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseAllowOutputReference
	_jsii_.Get(
		j,
		"allow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) AllowInput() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseAllow {
	var returns *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseAllow
	_jsii_.Get(
		j,
		"allowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) Block() Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlockOutputReference {
	var returns Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlockOutputReference
	_jsii_.Get(
		j,
		"block",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) BlockInput() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlock {
	var returns *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlock
	_jsii_.Get(
		j,
		"blockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) Captcha() Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCaptchaOutputReference {
	var returns Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCaptchaOutputReference
	_jsii_.Get(
		j,
		"captcha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) CaptchaInput() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCaptcha {
	var returns *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCaptcha
	_jsii_.Get(
		j,
		"captchaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) Challenge() Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference {
	var returns Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference
	_jsii_.Get(
		j,
		"challenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) ChallengeInput() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallenge {
	var returns *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallenge
	_jsii_.Get(
		j,
		"challengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) Count() Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCountOutputReference {
	var returns Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCountOutputReference
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) CountInput() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCount {
	var returns *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCount
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) InternalValue() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUse {
	var returns *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUse
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference_Override(w Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference)SetInternalValue(val *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUse) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) PutAllow(value *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseAllow) {
	if err := w.validatePutAllowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllow",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) PutBlock(value *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlock) {
	if err := w.validatePutBlockParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBlock",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) PutCaptcha(value *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCaptcha) {
	if err := w.validatePutCaptchaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCaptcha",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) PutChallenge(value *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallenge) {
	if err := w.validatePutChallengeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putChallenge",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) PutCount(value *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCount) {
	if err := w.validatePutCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCount",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) ResetAllow() {
	_jsii_.InvokeVoid(
		w,
		"resetAllow",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) ResetBlock() {
	_jsii_.InvokeVoid(
		w,
		"resetBlock",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) ResetCaptcha() {
	_jsii_.InvokeVoid(
		w,
		"resetCaptcha",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) ResetChallenge() {
	_jsii_.InvokeVoid(
		w,
		"resetChallenge",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		w,
		"resetCount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

