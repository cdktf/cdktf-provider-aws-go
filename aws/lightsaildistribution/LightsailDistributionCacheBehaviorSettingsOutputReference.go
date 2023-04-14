package lightsaildistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v13/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v13/lightsaildistribution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LightsailDistributionCacheBehaviorSettingsOutputReference interface {
	cdktf.ComplexObject
	AllowedHttpMethods() *string
	SetAllowedHttpMethods(val *string)
	AllowedHttpMethodsInput() *string
	CachedHttpMethods() *string
	SetCachedHttpMethods(val *string)
	CachedHttpMethodsInput() *string
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
	DefaultTtl() *float64
	SetDefaultTtl(val *float64)
	DefaultTtlInput() *float64
	ForwardedCookies() LightsailDistributionCacheBehaviorSettingsForwardedCookiesOutputReference
	ForwardedCookiesInput() *LightsailDistributionCacheBehaviorSettingsForwardedCookies
	ForwardedHeaders() LightsailDistributionCacheBehaviorSettingsForwardedHeadersOutputReference
	ForwardedHeadersInput() *LightsailDistributionCacheBehaviorSettingsForwardedHeaders
	ForwardedQueryStrings() LightsailDistributionCacheBehaviorSettingsForwardedQueryStringsOutputReference
	ForwardedQueryStringsInput() *LightsailDistributionCacheBehaviorSettingsForwardedQueryStrings
	// Experimental.
	Fqn() *string
	InternalValue() *LightsailDistributionCacheBehaviorSettings
	SetInternalValue(val *LightsailDistributionCacheBehaviorSettings)
	MaximumTtl() *float64
	SetMaximumTtl(val *float64)
	MaximumTtlInput() *float64
	MinimumTtl() *float64
	SetMinimumTtl(val *float64)
	MinimumTtlInput() *float64
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
	PutForwardedCookies(value *LightsailDistributionCacheBehaviorSettingsForwardedCookies)
	PutForwardedHeaders(value *LightsailDistributionCacheBehaviorSettingsForwardedHeaders)
	PutForwardedQueryStrings(value *LightsailDistributionCacheBehaviorSettingsForwardedQueryStrings)
	ResetAllowedHttpMethods()
	ResetCachedHttpMethods()
	ResetDefaultTtl()
	ResetForwardedCookies()
	ResetForwardedHeaders()
	ResetForwardedQueryStrings()
	ResetMaximumTtl()
	ResetMinimumTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LightsailDistributionCacheBehaviorSettingsOutputReference
type jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) AllowedHttpMethods() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedHttpMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) AllowedHttpMethodsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedHttpMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) CachedHttpMethods() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachedHttpMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) CachedHttpMethodsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachedHttpMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ForwardedCookies() LightsailDistributionCacheBehaviorSettingsForwardedCookiesOutputReference {
	var returns LightsailDistributionCacheBehaviorSettingsForwardedCookiesOutputReference
	_jsii_.Get(
		j,
		"forwardedCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ForwardedCookiesInput() *LightsailDistributionCacheBehaviorSettingsForwardedCookies {
	var returns *LightsailDistributionCacheBehaviorSettingsForwardedCookies
	_jsii_.Get(
		j,
		"forwardedCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ForwardedHeaders() LightsailDistributionCacheBehaviorSettingsForwardedHeadersOutputReference {
	var returns LightsailDistributionCacheBehaviorSettingsForwardedHeadersOutputReference
	_jsii_.Get(
		j,
		"forwardedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ForwardedHeadersInput() *LightsailDistributionCacheBehaviorSettingsForwardedHeaders {
	var returns *LightsailDistributionCacheBehaviorSettingsForwardedHeaders
	_jsii_.Get(
		j,
		"forwardedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ForwardedQueryStrings() LightsailDistributionCacheBehaviorSettingsForwardedQueryStringsOutputReference {
	var returns LightsailDistributionCacheBehaviorSettingsForwardedQueryStringsOutputReference
	_jsii_.Get(
		j,
		"forwardedQueryStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ForwardedQueryStringsInput() *LightsailDistributionCacheBehaviorSettingsForwardedQueryStrings {
	var returns *LightsailDistributionCacheBehaviorSettingsForwardedQueryStrings
	_jsii_.Get(
		j,
		"forwardedQueryStringsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) InternalValue() *LightsailDistributionCacheBehaviorSettings {
	var returns *LightsailDistributionCacheBehaviorSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) MaximumTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) MaximumTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) MinimumTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) MinimumTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLightsailDistributionCacheBehaviorSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LightsailDistributionCacheBehaviorSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewLightsailDistributionCacheBehaviorSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lightsailDistribution.LightsailDistributionCacheBehaviorSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLightsailDistributionCacheBehaviorSettingsOutputReference_Override(l LightsailDistributionCacheBehaviorSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lightsailDistribution.LightsailDistributionCacheBehaviorSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference)SetAllowedHttpMethods(val *string) {
	if err := j.validateSetAllowedHttpMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedHttpMethods",
		val,
	)
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference)SetCachedHttpMethods(val *string) {
	if err := j.validateSetCachedHttpMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachedHttpMethods",
		val,
	)
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference)SetDefaultTtl(val *float64) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference)SetInternalValue(val *LightsailDistributionCacheBehaviorSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference)SetMaximumTtl(val *float64) {
	if err := j.validateSetMaximumTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumTtl",
		val,
	)
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference)SetMinimumTtl(val *float64) {
	if err := j.validateSetMinimumTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumTtl",
		val,
	)
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) PutForwardedCookies(value *LightsailDistributionCacheBehaviorSettingsForwardedCookies) {
	if err := l.validatePutForwardedCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putForwardedCookies",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) PutForwardedHeaders(value *LightsailDistributionCacheBehaviorSettingsForwardedHeaders) {
	if err := l.validatePutForwardedHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putForwardedHeaders",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) PutForwardedQueryStrings(value *LightsailDistributionCacheBehaviorSettingsForwardedQueryStrings) {
	if err := l.validatePutForwardedQueryStringsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putForwardedQueryStrings",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ResetAllowedHttpMethods() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowedHttpMethods",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ResetCachedHttpMethods() {
	_jsii_.InvokeVoid(
		l,
		"resetCachedHttpMethods",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ResetForwardedCookies() {
	_jsii_.InvokeVoid(
		l,
		"resetForwardedCookies",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ResetForwardedHeaders() {
	_jsii_.InvokeVoid(
		l,
		"resetForwardedHeaders",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ResetForwardedQueryStrings() {
	_jsii_.InvokeVoid(
		l,
		"resetForwardedQueryStrings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ResetMaximumTtl() {
	_jsii_.InvokeVoid(
		l,
		"resetMaximumTtl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ResetMinimumTtl() {
	_jsii_.InvokeVoid(
		l,
		"resetMinimumTtl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDistributionCacheBehaviorSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

