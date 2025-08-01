// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package xraygroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/xraygroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type XrayGroupInsightsConfigurationOutputReference interface {
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
	InsightsEnabled() interface{}
	SetInsightsEnabled(val interface{})
	InsightsEnabledInput() interface{}
	InternalValue() *XrayGroupInsightsConfiguration
	SetInternalValue(val *XrayGroupInsightsConfiguration)
	NotificationsEnabled() interface{}
	SetNotificationsEnabled(val interface{})
	NotificationsEnabledInput() interface{}
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
	ResetNotificationsEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for XrayGroupInsightsConfigurationOutputReference
type jsiiProxy_XrayGroupInsightsConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) InsightsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insightsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) InsightsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insightsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) InternalValue() *XrayGroupInsightsConfiguration {
	var returns *XrayGroupInsightsConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) NotificationsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) NotificationsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewXrayGroupInsightsConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) XrayGroupInsightsConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewXrayGroupInsightsConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_XrayGroupInsightsConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.xrayGroup.XrayGroupInsightsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewXrayGroupInsightsConfigurationOutputReference_Override(x XrayGroupInsightsConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.xrayGroup.XrayGroupInsightsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		x,
	)
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference)SetInsightsEnabled(val interface{}) {
	if err := j.validateSetInsightsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insightsEnabled",
		val,
	)
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference)SetInternalValue(val *XrayGroupInsightsConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference)SetNotificationsEnabled(val interface{}) {
	if err := j.validateSetNotificationsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationsEnabled",
		val,
	)
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_XrayGroupInsightsConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := x.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		x,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := x.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		x,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := x.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		x,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := x.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		x,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := x.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		x,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := x.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		x,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := x.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		x,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := x.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		x,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := x.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		x,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		x,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := x.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		x,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) ResetNotificationsEnabled() {
	_jsii_.InvokeVoid(
		x,
		"resetNotificationsEnabled",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := x.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		x,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayGroupInsightsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

