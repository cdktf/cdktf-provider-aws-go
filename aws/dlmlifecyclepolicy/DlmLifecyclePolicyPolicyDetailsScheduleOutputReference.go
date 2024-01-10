// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dlmlifecyclepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DlmLifecyclePolicyPolicyDetailsScheduleOutputReference interface {
	cdktf.ComplexObject
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
	CreateRule() DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference
	CreateRuleInput() *DlmLifecyclePolicyPolicyDetailsScheduleCreateRule
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CrossRegionCopyRule() DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleList
	CrossRegionCopyRuleInput() interface{}
	DeprecateRule() DlmLifecyclePolicyPolicyDetailsScheduleDeprecateRuleOutputReference
	DeprecateRuleInput() *DlmLifecyclePolicyPolicyDetailsScheduleDeprecateRule
	FastRestoreRule() DlmLifecyclePolicyPolicyDetailsScheduleFastRestoreRuleOutputReference
	FastRestoreRuleInput() *DlmLifecyclePolicyPolicyDetailsScheduleFastRestoreRule
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	RetainRule() DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference
	RetainRuleInput() *DlmLifecyclePolicyPolicyDetailsScheduleRetainRule
	ShareRule() DlmLifecyclePolicyPolicyDetailsScheduleShareRuleOutputReference
	ShareRuleInput() *DlmLifecyclePolicyPolicyDetailsScheduleShareRule
	TagsToAdd() *map[string]*string
	SetTagsToAdd(val *map[string]*string)
	TagsToAddInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VariableTags() *map[string]*string
	SetVariableTags(val *map[string]*string)
	VariableTagsInput() *map[string]*string
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
	PutCreateRule(value *DlmLifecyclePolicyPolicyDetailsScheduleCreateRule)
	PutCrossRegionCopyRule(value interface{})
	PutDeprecateRule(value *DlmLifecyclePolicyPolicyDetailsScheduleDeprecateRule)
	PutFastRestoreRule(value *DlmLifecyclePolicyPolicyDetailsScheduleFastRestoreRule)
	PutRetainRule(value *DlmLifecyclePolicyPolicyDetailsScheduleRetainRule)
	PutShareRule(value *DlmLifecyclePolicyPolicyDetailsScheduleShareRule)
	ResetCopyTags()
	ResetCrossRegionCopyRule()
	ResetDeprecateRule()
	ResetFastRestoreRule()
	ResetShareRule()
	ResetTagsToAdd()
	ResetVariableTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DlmLifecyclePolicyPolicyDetailsScheduleOutputReference
type jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) CopyTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) CopyTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) CreateRule() DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsScheduleCreateRuleOutputReference
	_jsii_.Get(
		j,
		"createRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) CreateRuleInput() *DlmLifecyclePolicyPolicyDetailsScheduleCreateRule {
	var returns *DlmLifecyclePolicyPolicyDetailsScheduleCreateRule
	_jsii_.Get(
		j,
		"createRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) CrossRegionCopyRule() DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleList {
	var returns DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleList
	_jsii_.Get(
		j,
		"crossRegionCopyRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) CrossRegionCopyRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossRegionCopyRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) DeprecateRule() DlmLifecyclePolicyPolicyDetailsScheduleDeprecateRuleOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsScheduleDeprecateRuleOutputReference
	_jsii_.Get(
		j,
		"deprecateRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) DeprecateRuleInput() *DlmLifecyclePolicyPolicyDetailsScheduleDeprecateRule {
	var returns *DlmLifecyclePolicyPolicyDetailsScheduleDeprecateRule
	_jsii_.Get(
		j,
		"deprecateRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) FastRestoreRule() DlmLifecyclePolicyPolicyDetailsScheduleFastRestoreRuleOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsScheduleFastRestoreRuleOutputReference
	_jsii_.Get(
		j,
		"fastRestoreRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) FastRestoreRuleInput() *DlmLifecyclePolicyPolicyDetailsScheduleFastRestoreRule {
	var returns *DlmLifecyclePolicyPolicyDetailsScheduleFastRestoreRule
	_jsii_.Get(
		j,
		"fastRestoreRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) RetainRule() DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsScheduleRetainRuleOutputReference
	_jsii_.Get(
		j,
		"retainRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) RetainRuleInput() *DlmLifecyclePolicyPolicyDetailsScheduleRetainRule {
	var returns *DlmLifecyclePolicyPolicyDetailsScheduleRetainRule
	_jsii_.Get(
		j,
		"retainRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) ShareRule() DlmLifecyclePolicyPolicyDetailsScheduleShareRuleOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsScheduleShareRuleOutputReference
	_jsii_.Get(
		j,
		"shareRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) ShareRuleInput() *DlmLifecyclePolicyPolicyDetailsScheduleShareRule {
	var returns *DlmLifecyclePolicyPolicyDetailsScheduleShareRule
	_jsii_.Get(
		j,
		"shareRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) TagsToAdd() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsToAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) TagsToAddInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsToAddInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) VariableTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variableTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) VariableTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"variableTagsInput",
		&returns,
	)
	return returns
}


func NewDlmLifecyclePolicyPolicyDetailsScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DlmLifecyclePolicyPolicyDetailsScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewDlmLifecyclePolicyPolicyDetailsScheduleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dlmLifecyclePolicy.DlmLifecyclePolicyPolicyDetailsScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDlmLifecyclePolicyPolicyDetailsScheduleOutputReference_Override(d DlmLifecyclePolicyPolicyDetailsScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dlmLifecyclePolicy.DlmLifecyclePolicyPolicyDetailsScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference)SetCopyTags(val interface{}) {
	if err := j.validateSetCopyTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTags",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference)SetTagsToAdd(val *map[string]*string) {
	if err := j.validateSetTagsToAddParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsToAdd",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference)SetVariableTags(val *map[string]*string) {
	if err := j.validateSetVariableTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variableTags",
		val,
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) PutCreateRule(value *DlmLifecyclePolicyPolicyDetailsScheduleCreateRule) {
	if err := d.validatePutCreateRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCreateRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) PutCrossRegionCopyRule(value interface{}) {
	if err := d.validatePutCrossRegionCopyRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCrossRegionCopyRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) PutDeprecateRule(value *DlmLifecyclePolicyPolicyDetailsScheduleDeprecateRule) {
	if err := d.validatePutDeprecateRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDeprecateRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) PutFastRestoreRule(value *DlmLifecyclePolicyPolicyDetailsScheduleFastRestoreRule) {
	if err := d.validatePutFastRestoreRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFastRestoreRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) PutRetainRule(value *DlmLifecyclePolicyPolicyDetailsScheduleRetainRule) {
	if err := d.validatePutRetainRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRetainRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) PutShareRule(value *DlmLifecyclePolicyPolicyDetailsScheduleShareRule) {
	if err := d.validatePutShareRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putShareRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) ResetCopyTags() {
	_jsii_.InvokeVoid(
		d,
		"resetCopyTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) ResetCrossRegionCopyRule() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossRegionCopyRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) ResetDeprecateRule() {
	_jsii_.InvokeVoid(
		d,
		"resetDeprecateRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) ResetFastRestoreRule() {
	_jsii_.InvokeVoid(
		d,
		"resetFastRestoreRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) ResetShareRule() {
	_jsii_.InvokeVoid(
		d,
		"resetShareRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) ResetTagsToAdd() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsToAdd",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) ResetVariableTags() {
	_jsii_.InvokeVoid(
		d,
		"resetVariableTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

