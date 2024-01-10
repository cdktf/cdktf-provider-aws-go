// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customerprofilesdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/customerprofilesdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomerprofilesDomainMatchingAutoMergingOutputReference interface {
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
	ConflictResolution() CustomerprofilesDomainMatchingAutoMergingConflictResolutionOutputReference
	ConflictResolutionInput() *CustomerprofilesDomainMatchingAutoMergingConflictResolution
	Consolidation() CustomerprofilesDomainMatchingAutoMergingConsolidationOutputReference
	ConsolidationInput() *CustomerprofilesDomainMatchingAutoMergingConsolidation
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *CustomerprofilesDomainMatchingAutoMerging
	SetInternalValue(val *CustomerprofilesDomainMatchingAutoMerging)
	MinAllowedConfidenceScoreForMerging() *float64
	SetMinAllowedConfidenceScoreForMerging(val *float64)
	MinAllowedConfidenceScoreForMergingInput() *float64
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
	PutConflictResolution(value *CustomerprofilesDomainMatchingAutoMergingConflictResolution)
	PutConsolidation(value *CustomerprofilesDomainMatchingAutoMergingConsolidation)
	ResetConflictResolution()
	ResetConsolidation()
	ResetMinAllowedConfidenceScoreForMerging()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomerprofilesDomainMatchingAutoMergingOutputReference
type jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) ConflictResolution() CustomerprofilesDomainMatchingAutoMergingConflictResolutionOutputReference {
	var returns CustomerprofilesDomainMatchingAutoMergingConflictResolutionOutputReference
	_jsii_.Get(
		j,
		"conflictResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) ConflictResolutionInput() *CustomerprofilesDomainMatchingAutoMergingConflictResolution {
	var returns *CustomerprofilesDomainMatchingAutoMergingConflictResolution
	_jsii_.Get(
		j,
		"conflictResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) Consolidation() CustomerprofilesDomainMatchingAutoMergingConsolidationOutputReference {
	var returns CustomerprofilesDomainMatchingAutoMergingConsolidationOutputReference
	_jsii_.Get(
		j,
		"consolidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) ConsolidationInput() *CustomerprofilesDomainMatchingAutoMergingConsolidation {
	var returns *CustomerprofilesDomainMatchingAutoMergingConsolidation
	_jsii_.Get(
		j,
		"consolidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) InternalValue() *CustomerprofilesDomainMatchingAutoMerging {
	var returns *CustomerprofilesDomainMatchingAutoMerging
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) MinAllowedConfidenceScoreForMerging() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAllowedConfidenceScoreForMerging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) MinAllowedConfidenceScoreForMergingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAllowedConfidenceScoreForMergingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCustomerprofilesDomainMatchingAutoMergingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CustomerprofilesDomainMatchingAutoMergingOutputReference {
	_init_.Initialize()

	if err := validateNewCustomerprofilesDomainMatchingAutoMergingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.customerprofilesDomain.CustomerprofilesDomainMatchingAutoMergingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomerprofilesDomainMatchingAutoMergingOutputReference_Override(c CustomerprofilesDomainMatchingAutoMergingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.customerprofilesDomain.CustomerprofilesDomainMatchingAutoMergingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference)SetInternalValue(val *CustomerprofilesDomainMatchingAutoMerging) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference)SetMinAllowedConfidenceScoreForMerging(val *float64) {
	if err := j.validateSetMinAllowedConfidenceScoreForMergingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minAllowedConfidenceScoreForMerging",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) PutConflictResolution(value *CustomerprofilesDomainMatchingAutoMergingConflictResolution) {
	if err := c.validatePutConflictResolutionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConflictResolution",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) PutConsolidation(value *CustomerprofilesDomainMatchingAutoMergingConsolidation) {
	if err := c.validatePutConsolidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConsolidation",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) ResetConflictResolution() {
	_jsii_.InvokeVoid(
		c,
		"resetConflictResolution",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) ResetConsolidation() {
	_jsii_.InvokeVoid(
		c,
		"resetConsolidation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) ResetMinAllowedConfidenceScoreForMerging() {
	_jsii_.InvokeVoid(
		c,
		"resetMinAllowedConfidenceScoreForMerging",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CustomerprofilesDomainMatchingAutoMergingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

