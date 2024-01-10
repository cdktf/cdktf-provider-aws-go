// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rumappmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/rumappmonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RumAppMonitorAppMonitorConfigurationOutputReference interface {
	cdktf.ComplexObject
	AllowCookies() interface{}
	SetAllowCookies(val interface{})
	AllowCookiesInput() interface{}
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
	EnableXray() interface{}
	SetEnableXray(val interface{})
	EnableXrayInput() interface{}
	ExcludedPages() *[]*string
	SetExcludedPages(val *[]*string)
	ExcludedPagesInput() *[]*string
	FavoritePages() *[]*string
	SetFavoritePages(val *[]*string)
	FavoritePagesInput() *[]*string
	// Experimental.
	Fqn() *string
	GuestRoleArn() *string
	SetGuestRoleArn(val *string)
	GuestRoleArnInput() *string
	IdentityPoolId() *string
	SetIdentityPoolId(val *string)
	IdentityPoolIdInput() *string
	IncludedPages() *[]*string
	SetIncludedPages(val *[]*string)
	IncludedPagesInput() *[]*string
	InternalValue() *RumAppMonitorAppMonitorConfiguration
	SetInternalValue(val *RumAppMonitorAppMonitorConfiguration)
	SessionSampleRate() *float64
	SetSessionSampleRate(val *float64)
	SessionSampleRateInput() *float64
	Telemetries() *[]*string
	SetTelemetries(val *[]*string)
	TelemetriesInput() *[]*string
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
	ResetAllowCookies()
	ResetEnableXray()
	ResetExcludedPages()
	ResetFavoritePages()
	ResetGuestRoleArn()
	ResetIdentityPoolId()
	ResetIncludedPages()
	ResetSessionSampleRate()
	ResetTelemetries()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RumAppMonitorAppMonitorConfigurationOutputReference
type jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) AllowCookies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) AllowCookiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) EnableXray() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableXray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) EnableXrayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableXrayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ExcludedPages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ExcludedPagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) FavoritePages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"favoritePages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) FavoritePagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"favoritePagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) GuestRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) GuestRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) IdentityPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) IncludedPages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedPages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) IncludedPagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedPagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) InternalValue() *RumAppMonitorAppMonitorConfiguration {
	var returns *RumAppMonitorAppMonitorConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) SessionSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) SessionSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) Telemetries() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"telemetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) TelemetriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"telemetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRumAppMonitorAppMonitorConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RumAppMonitorAppMonitorConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewRumAppMonitorAppMonitorConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.rumAppMonitor.RumAppMonitorAppMonitorConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRumAppMonitorAppMonitorConfigurationOutputReference_Override(r RumAppMonitorAppMonitorConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.rumAppMonitor.RumAppMonitorAppMonitorConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference)SetAllowCookies(val interface{}) {
	if err := j.validateSetAllowCookiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowCookies",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference)SetEnableXray(val interface{}) {
	if err := j.validateSetEnableXrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableXray",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference)SetExcludedPages(val *[]*string) {
	if err := j.validateSetExcludedPagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedPages",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference)SetFavoritePages(val *[]*string) {
	if err := j.validateSetFavoritePagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"favoritePages",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference)SetGuestRoleArn(val *string) {
	if err := j.validateSetGuestRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestRoleArn",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference)SetIdentityPoolId(val *string) {
	if err := j.validateSetIdentityPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityPoolId",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference)SetIncludedPages(val *[]*string) {
	if err := j.validateSetIncludedPagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedPages",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference)SetInternalValue(val *RumAppMonitorAppMonitorConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference)SetSessionSampleRate(val *float64) {
	if err := j.validateSetSessionSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionSampleRate",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference)SetTelemetries(val *[]*string) {
	if err := j.validateSetTelemetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"telemetries",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ResetAllowCookies() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowCookies",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ResetEnableXray() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableXray",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ResetExcludedPages() {
	_jsii_.InvokeVoid(
		r,
		"resetExcludedPages",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ResetFavoritePages() {
	_jsii_.InvokeVoid(
		r,
		"resetFavoritePages",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ResetGuestRoleArn() {
	_jsii_.InvokeVoid(
		r,
		"resetGuestRoleArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ResetIdentityPoolId() {
	_jsii_.InvokeVoid(
		r,
		"resetIdentityPoolId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ResetIncludedPages() {
	_jsii_.InvokeVoid(
		r,
		"resetIncludedPages",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ResetSessionSampleRate() {
	_jsii_.InvokeVoid(
		r,
		"resetSessionSampleRate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ResetTelemetries() {
	_jsii_.InvokeVoid(
		r,
		"resetTelemetries",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

