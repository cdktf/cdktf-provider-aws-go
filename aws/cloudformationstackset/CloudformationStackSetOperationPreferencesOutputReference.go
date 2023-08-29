// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudformationstackset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/cloudformationstackset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudformationStackSetOperationPreferencesOutputReference interface {
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
	FailureToleranceCount() *float64
	SetFailureToleranceCount(val *float64)
	FailureToleranceCountInput() *float64
	FailureTolerancePercentage() *float64
	SetFailureTolerancePercentage(val *float64)
	FailureTolerancePercentageInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *CloudformationStackSetOperationPreferences
	SetInternalValue(val *CloudformationStackSetOperationPreferences)
	MaxConcurrentCount() *float64
	SetMaxConcurrentCount(val *float64)
	MaxConcurrentCountInput() *float64
	MaxConcurrentPercentage() *float64
	SetMaxConcurrentPercentage(val *float64)
	MaxConcurrentPercentageInput() *float64
	RegionConcurrencyType() *string
	SetRegionConcurrencyType(val *string)
	RegionConcurrencyTypeInput() *string
	RegionOrder() *[]*string
	SetRegionOrder(val *[]*string)
	RegionOrderInput() *[]*string
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
	ResetFailureToleranceCount()
	ResetFailureTolerancePercentage()
	ResetMaxConcurrentCount()
	ResetMaxConcurrentPercentage()
	ResetRegionConcurrencyType()
	ResetRegionOrder()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudformationStackSetOperationPreferencesOutputReference
type jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) FailureToleranceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureToleranceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) FailureToleranceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureToleranceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) FailureTolerancePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureTolerancePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) FailureTolerancePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureTolerancePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) InternalValue() *CloudformationStackSetOperationPreferences {
	var returns *CloudformationStackSetOperationPreferences
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) MaxConcurrentCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) MaxConcurrentCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) MaxConcurrentPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) MaxConcurrentPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) RegionConcurrencyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionConcurrencyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) RegionConcurrencyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionConcurrencyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) RegionOrder() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) RegionOrderInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudformationStackSetOperationPreferencesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudformationStackSetOperationPreferencesOutputReference {
	_init_.Initialize()

	if err := validateNewCloudformationStackSetOperationPreferencesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudformationStackSet.CloudformationStackSetOperationPreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudformationStackSetOperationPreferencesOutputReference_Override(c CloudformationStackSetOperationPreferencesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudformationStackSet.CloudformationStackSetOperationPreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference)SetFailureToleranceCount(val *float64) {
	if err := j.validateSetFailureToleranceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureToleranceCount",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference)SetFailureTolerancePercentage(val *float64) {
	if err := j.validateSetFailureTolerancePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureTolerancePercentage",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference)SetInternalValue(val *CloudformationStackSetOperationPreferences) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference)SetMaxConcurrentCount(val *float64) {
	if err := j.validateSetMaxConcurrentCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentCount",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference)SetMaxConcurrentPercentage(val *float64) {
	if err := j.validateSetMaxConcurrentPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentPercentage",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference)SetRegionConcurrencyType(val *string) {
	if err := j.validateSetRegionConcurrencyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionConcurrencyType",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference)SetRegionOrder(val *[]*string) {
	if err := j.validateSetRegionOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionOrder",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) ResetFailureToleranceCount() {
	_jsii_.InvokeVoid(
		c,
		"resetFailureToleranceCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) ResetFailureTolerancePercentage() {
	_jsii_.InvokeVoid(
		c,
		"resetFailureTolerancePercentage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) ResetMaxConcurrentCount() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxConcurrentCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) ResetMaxConcurrentPercentage() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxConcurrentPercentage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) ResetRegionConcurrencyType() {
	_jsii_.InvokeVoid(
		c,
		"resetRegionConcurrencyType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) ResetRegionOrder() {
	_jsii_.InvokeVoid(
		c,
		"resetRegionOrder",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudformationStackSetOperationPreferencesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

