// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewaymethodsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/apigatewaymethodsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiGatewayMethodSettingsSettingsOutputReference interface {
	cdktf.ComplexObject
	CacheDataEncrypted() interface{}
	SetCacheDataEncrypted(val interface{})
	CacheDataEncryptedInput() interface{}
	CacheTtlInSeconds() *float64
	SetCacheTtlInSeconds(val *float64)
	CacheTtlInSecondsInput() *float64
	CachingEnabled() interface{}
	SetCachingEnabled(val interface{})
	CachingEnabledInput() interface{}
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
	DataTraceEnabled() interface{}
	SetDataTraceEnabled(val interface{})
	DataTraceEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ApiGatewayMethodSettingsSettings
	SetInternalValue(val *ApiGatewayMethodSettingsSettings)
	LoggingLevel() *string
	SetLoggingLevel(val *string)
	LoggingLevelInput() *string
	MetricsEnabled() interface{}
	SetMetricsEnabled(val interface{})
	MetricsEnabledInput() interface{}
	RequireAuthorizationForCacheControl() interface{}
	SetRequireAuthorizationForCacheControl(val interface{})
	RequireAuthorizationForCacheControlInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThrottlingBurstLimit() *float64
	SetThrottlingBurstLimit(val *float64)
	ThrottlingBurstLimitInput() *float64
	ThrottlingRateLimit() *float64
	SetThrottlingRateLimit(val *float64)
	ThrottlingRateLimitInput() *float64
	UnauthorizedCacheControlHeaderStrategy() *string
	SetUnauthorizedCacheControlHeaderStrategy(val *string)
	UnauthorizedCacheControlHeaderStrategyInput() *string
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
	ResetCacheDataEncrypted()
	ResetCacheTtlInSeconds()
	ResetCachingEnabled()
	ResetDataTraceEnabled()
	ResetLoggingLevel()
	ResetMetricsEnabled()
	ResetRequireAuthorizationForCacheControl()
	ResetThrottlingBurstLimit()
	ResetThrottlingRateLimit()
	ResetUnauthorizedCacheControlHeaderStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiGatewayMethodSettingsSettingsOutputReference
type jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) CacheDataEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheDataEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) CacheDataEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheDataEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) CacheTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) CacheTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) CachingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cachingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) CachingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cachingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) DataTraceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTraceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) DataTraceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTraceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) InternalValue() *ApiGatewayMethodSettingsSettings {
	var returns *ApiGatewayMethodSettingsSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) LoggingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) LoggingLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) MetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) MetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) RequireAuthorizationForCacheControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAuthorizationForCacheControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) RequireAuthorizationForCacheControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAuthorizationForCacheControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ThrottlingBurstLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingBurstLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ThrottlingBurstLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingBurstLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ThrottlingRateLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingRateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ThrottlingRateLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingRateLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) UnauthorizedCacheControlHeaderStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthorizedCacheControlHeaderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) UnauthorizedCacheControlHeaderStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthorizedCacheControlHeaderStrategyInput",
		&returns,
	)
	return returns
}


func NewApiGatewayMethodSettingsSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiGatewayMethodSettingsSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewApiGatewayMethodSettingsSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.apiGatewayMethodSettings.ApiGatewayMethodSettingsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiGatewayMethodSettingsSettingsOutputReference_Override(a ApiGatewayMethodSettingsSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.apiGatewayMethodSettings.ApiGatewayMethodSettingsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetCacheDataEncrypted(val interface{}) {
	if err := j.validateSetCacheDataEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheDataEncrypted",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetCacheTtlInSeconds(val *float64) {
	if err := j.validateSetCacheTtlInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetCachingEnabled(val interface{}) {
	if err := j.validateSetCachingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachingEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetDataTraceEnabled(val interface{}) {
	if err := j.validateSetDataTraceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTraceEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetInternalValue(val *ApiGatewayMethodSettingsSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetLoggingLevel(val *string) {
	if err := j.validateSetLoggingLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingLevel",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetMetricsEnabled(val interface{}) {
	if err := j.validateSetMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetRequireAuthorizationForCacheControl(val interface{}) {
	if err := j.validateSetRequireAuthorizationForCacheControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAuthorizationForCacheControl",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetThrottlingBurstLimit(val *float64) {
	if err := j.validateSetThrottlingBurstLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttlingBurstLimit",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetThrottlingRateLimit(val *float64) {
	if err := j.validateSetThrottlingRateLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttlingRateLimit",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference)SetUnauthorizedCacheControlHeaderStrategy(val *string) {
	if err := j.validateSetUnauthorizedCacheControlHeaderStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unauthorizedCacheControlHeaderStrategy",
		val,
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetCacheDataEncrypted() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheDataEncrypted",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetCacheTtlInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheTtlInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetCachingEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetCachingEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetDataTraceEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDataTraceEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetLoggingLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetLoggingLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetMetricsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetRequireAuthorizationForCacheControl() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireAuthorizationForCacheControl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetThrottlingBurstLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottlingBurstLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetThrottlingRateLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottlingRateLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ResetUnauthorizedCacheControlHeaderStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetUnauthorizedCacheControlHeaderStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

