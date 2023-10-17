// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *Wafv2WebAclRuleStatementManagedRuleGroupStatement
	SetInternalValue(val *Wafv2WebAclRuleStatementManagedRuleGroupStatement)
	ManagedRuleGroupConfigs() Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsList
	ManagedRuleGroupConfigsInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	RuleActionOverride() Wafv2WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideList
	RuleActionOverrideInput() interface{}
	ScopeDownStatement() Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementOutputReference
	ScopeDownStatementInput() *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatement
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VendorName() *string
	SetVendorName(val *string)
	VendorNameInput() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutManagedRuleGroupConfigs(value interface{})
	PutRuleActionOverride(value interface{})
	PutScopeDownStatement(value *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatement)
	ResetManagedRuleGroupConfigs()
	ResetRuleActionOverride()
	ResetScopeDownStatement()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference
type jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) InternalValue() *Wafv2WebAclRuleStatementManagedRuleGroupStatement {
	var returns *Wafv2WebAclRuleStatementManagedRuleGroupStatement
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) ManagedRuleGroupConfigs() Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsList {
	var returns Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsList
	_jsii_.Get(
		j,
		"managedRuleGroupConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) ManagedRuleGroupConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedRuleGroupConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) RuleActionOverride() Wafv2WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideList {
	var returns Wafv2WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideList
	_jsii_.Get(
		j,
		"ruleActionOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) RuleActionOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleActionOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) ScopeDownStatement() Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementOutputReference {
	var returns Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementOutputReference
	_jsii_.Get(
		j,
		"scopeDownStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) ScopeDownStatementInput() *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatement {
	var returns *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatement
	_jsii_.Get(
		j,
		"scopeDownStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) VendorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vendorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) VendorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vendorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleStatementManagedRuleGroupStatementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference_Override(w Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference)SetInternalValue(val *Wafv2WebAclRuleStatementManagedRuleGroupStatement) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference)SetVendorName(val *string) {
	if err := j.validateSetVendorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vendorName",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) PutManagedRuleGroupConfigs(value interface{}) {
	if err := w.validatePutManagedRuleGroupConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putManagedRuleGroupConfigs",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) PutRuleActionOverride(value interface{}) {
	if err := w.validatePutRuleActionOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRuleActionOverride",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) PutScopeDownStatement(value *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatement) {
	if err := w.validatePutScopeDownStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putScopeDownStatement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) ResetManagedRuleGroupConfigs() {
	_jsii_.InvokeVoid(
		w,
		"resetManagedRuleGroupConfigs",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) ResetRuleActionOverride() {
	_jsii_.InvokeVoid(
		w,
		"resetRuleActionOverride",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) ResetScopeDownStatement() {
	_jsii_.InvokeVoid(
		w,
		"resetScopeDownStatement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		w,
		"resetVersion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

