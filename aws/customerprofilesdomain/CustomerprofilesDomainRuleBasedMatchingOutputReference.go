// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customerprofilesdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/customerprofilesdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomerprofilesDomainRuleBasedMatchingOutputReference interface {
	cdktf.ComplexObject
	AttributeTypesSelector() CustomerprofilesDomainRuleBasedMatchingAttributeTypesSelectorOutputReference
	AttributeTypesSelectorInput() *CustomerprofilesDomainRuleBasedMatchingAttributeTypesSelector
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
	ConflictResolution() CustomerprofilesDomainRuleBasedMatchingConflictResolutionOutputReference
	ConflictResolutionInput() *CustomerprofilesDomainRuleBasedMatchingConflictResolution
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ExportingConfig() CustomerprofilesDomainRuleBasedMatchingExportingConfigOutputReference
	ExportingConfigInput() *CustomerprofilesDomainRuleBasedMatchingExportingConfig
	// Experimental.
	Fqn() *string
	InternalValue() *CustomerprofilesDomainRuleBasedMatching
	SetInternalValue(val *CustomerprofilesDomainRuleBasedMatching)
	MatchingRules() CustomerprofilesDomainRuleBasedMatchingMatchingRulesList
	MatchingRulesInput() interface{}
	MaxAllowedRuleLevelForMatching() *float64
	SetMaxAllowedRuleLevelForMatching(val *float64)
	MaxAllowedRuleLevelForMatchingInput() *float64
	MaxAllowedRuleLevelForMerging() *float64
	SetMaxAllowedRuleLevelForMerging(val *float64)
	MaxAllowedRuleLevelForMergingInput() *float64
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
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
	PutAttributeTypesSelector(value *CustomerprofilesDomainRuleBasedMatchingAttributeTypesSelector)
	PutConflictResolution(value *CustomerprofilesDomainRuleBasedMatchingConflictResolution)
	PutExportingConfig(value *CustomerprofilesDomainRuleBasedMatchingExportingConfig)
	PutMatchingRules(value interface{})
	ResetAttributeTypesSelector()
	ResetConflictResolution()
	ResetExportingConfig()
	ResetMatchingRules()
	ResetMaxAllowedRuleLevelForMatching()
	ResetMaxAllowedRuleLevelForMerging()
	ResetStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomerprofilesDomainRuleBasedMatchingOutputReference
type jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) AttributeTypesSelector() CustomerprofilesDomainRuleBasedMatchingAttributeTypesSelectorOutputReference {
	var returns CustomerprofilesDomainRuleBasedMatchingAttributeTypesSelectorOutputReference
	_jsii_.Get(
		j,
		"attributeTypesSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) AttributeTypesSelectorInput() *CustomerprofilesDomainRuleBasedMatchingAttributeTypesSelector {
	var returns *CustomerprofilesDomainRuleBasedMatchingAttributeTypesSelector
	_jsii_.Get(
		j,
		"attributeTypesSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ConflictResolution() CustomerprofilesDomainRuleBasedMatchingConflictResolutionOutputReference {
	var returns CustomerprofilesDomainRuleBasedMatchingConflictResolutionOutputReference
	_jsii_.Get(
		j,
		"conflictResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ConflictResolutionInput() *CustomerprofilesDomainRuleBasedMatchingConflictResolution {
	var returns *CustomerprofilesDomainRuleBasedMatchingConflictResolution
	_jsii_.Get(
		j,
		"conflictResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ExportingConfig() CustomerprofilesDomainRuleBasedMatchingExportingConfigOutputReference {
	var returns CustomerprofilesDomainRuleBasedMatchingExportingConfigOutputReference
	_jsii_.Get(
		j,
		"exportingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ExportingConfigInput() *CustomerprofilesDomainRuleBasedMatchingExportingConfig {
	var returns *CustomerprofilesDomainRuleBasedMatchingExportingConfig
	_jsii_.Get(
		j,
		"exportingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) InternalValue() *CustomerprofilesDomainRuleBasedMatching {
	var returns *CustomerprofilesDomainRuleBasedMatching
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) MatchingRules() CustomerprofilesDomainRuleBasedMatchingMatchingRulesList {
	var returns CustomerprofilesDomainRuleBasedMatchingMatchingRulesList
	_jsii_.Get(
		j,
		"matchingRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) MatchingRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchingRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) MaxAllowedRuleLevelForMatching() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedRuleLevelForMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) MaxAllowedRuleLevelForMatchingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedRuleLevelForMatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) MaxAllowedRuleLevelForMerging() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedRuleLevelForMerging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) MaxAllowedRuleLevelForMergingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedRuleLevelForMergingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCustomerprofilesDomainRuleBasedMatchingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CustomerprofilesDomainRuleBasedMatchingOutputReference {
	_init_.Initialize()

	if err := validateNewCustomerprofilesDomainRuleBasedMatchingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.customerprofilesDomain.CustomerprofilesDomainRuleBasedMatchingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomerprofilesDomainRuleBasedMatchingOutputReference_Override(c CustomerprofilesDomainRuleBasedMatchingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.customerprofilesDomain.CustomerprofilesDomainRuleBasedMatchingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference)SetInternalValue(val *CustomerprofilesDomainRuleBasedMatching) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference)SetMaxAllowedRuleLevelForMatching(val *float64) {
	if err := j.validateSetMaxAllowedRuleLevelForMatchingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAllowedRuleLevelForMatching",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference)SetMaxAllowedRuleLevelForMerging(val *float64) {
	if err := j.validateSetMaxAllowedRuleLevelForMergingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAllowedRuleLevelForMerging",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) PutAttributeTypesSelector(value *CustomerprofilesDomainRuleBasedMatchingAttributeTypesSelector) {
	if err := c.validatePutAttributeTypesSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAttributeTypesSelector",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) PutConflictResolution(value *CustomerprofilesDomainRuleBasedMatchingConflictResolution) {
	if err := c.validatePutConflictResolutionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConflictResolution",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) PutExportingConfig(value *CustomerprofilesDomainRuleBasedMatchingExportingConfig) {
	if err := c.validatePutExportingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExportingConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) PutMatchingRules(value interface{}) {
	if err := c.validatePutMatchingRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMatchingRules",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ResetAttributeTypesSelector() {
	_jsii_.InvokeVoid(
		c,
		"resetAttributeTypesSelector",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ResetConflictResolution() {
	_jsii_.InvokeVoid(
		c,
		"resetConflictResolution",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ResetExportingConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetExportingConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ResetMatchingRules() {
	_jsii_.InvokeVoid(
		c,
		"resetMatchingRules",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ResetMaxAllowedRuleLevelForMatching() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxAllowedRuleLevelForMatching",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ResetMaxAllowedRuleLevelForMerging() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxAllowedRuleLevelForMerging",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		c,
		"resetStatus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainRuleBasedMatchingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

