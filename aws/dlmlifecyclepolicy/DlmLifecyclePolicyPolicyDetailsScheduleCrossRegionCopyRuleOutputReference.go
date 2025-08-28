// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dlmlifecyclepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference interface {
	cdktf.ComplexObject
	CmkArn() *string
	SetCmkArn(val *string)
	CmkArnInput() *string
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
	CopyTags() interface{}
	SetCopyTags(val interface{})
	CopyTagsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeprecateRule() DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference
	DeprecateRuleInput() *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule
	Encrypted() interface{}
	SetEncrypted(val interface{})
	EncryptedInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RetainRule() DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference
	RetainRuleInput() *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
	TargetRegion() *string
	SetTargetRegion(val *string)
	TargetRegionInput() *string
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
	PutDeprecateRule(value *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule)
	PutRetainRule(value *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule)
	ResetCmkArn()
	ResetCopyTags()
	ResetDeprecateRule()
	ResetRetainRule()
	ResetTarget()
	ResetTargetRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference
type jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) CmkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) CmkArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmkArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) CopyTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) CopyTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) DeprecateRule() DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleOutputReference
	_jsii_.Get(
		j,
		"deprecateRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) DeprecateRuleInput() *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule {
	var returns *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule
	_jsii_.Get(
		j,
		"deprecateRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) Encrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) EncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) RetainRule() DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleOutputReference
	_jsii_.Get(
		j,
		"retainRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) RetainRuleInput() *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule {
	var returns *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule
	_jsii_.Get(
		j,
		"retainRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) TargetRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) TargetRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference {
	_init_.Initialize()

	if err := validateNewDlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dlmLifecyclePolicy.DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference_Override(d DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dlmLifecyclePolicy.DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference)SetCmkArn(val *string) {
	if err := j.validateSetCmkArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cmkArn",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference)SetCopyTags(val interface{}) {
	if err := j.validateSetCopyTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTags",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference)SetEncrypted(val interface{}) {
	if err := j.validateSetEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference)SetTargetRegion(val *string) {
	if err := j.validateSetTargetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRegion",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) PutDeprecateRule(value *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule) {
	if err := d.validatePutDeprecateRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDeprecateRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) PutRetainRule(value *DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule) {
	if err := d.validatePutRetainRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRetainRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) ResetCmkArn() {
	_jsii_.InvokeVoid(
		d,
		"resetCmkArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) ResetCopyTags() {
	_jsii_.InvokeVoid(
		d,
		"resetCopyTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) ResetDeprecateRule() {
	_jsii_.InvokeVoid(
		d,
		"resetDeprecateRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) ResetRetainRule() {
	_jsii_.InvokeVoid(
		d,
		"resetRetainRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) ResetTargetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

