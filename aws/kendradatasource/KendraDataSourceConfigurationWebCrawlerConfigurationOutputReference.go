// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/kendradatasource/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference interface {
	cdktf.ComplexObject
	AuthenticationConfiguration() KendraDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationOutputReference
	AuthenticationConfigurationInput() *KendraDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfiguration
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
	CrawlDepth() *float64
	SetCrawlDepth(val *float64)
	CrawlDepthInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KendraDataSourceConfigurationWebCrawlerConfiguration
	SetInternalValue(val *KendraDataSourceConfigurationWebCrawlerConfiguration)
	MaxContentSizePerPageInMegaBytes() *float64
	SetMaxContentSizePerPageInMegaBytes(val *float64)
	MaxContentSizePerPageInMegaBytesInput() *float64
	MaxLinksPerPage() *float64
	SetMaxLinksPerPage(val *float64)
	MaxLinksPerPageInput() *float64
	MaxUrlsPerMinuteCrawlRate() *float64
	SetMaxUrlsPerMinuteCrawlRate(val *float64)
	MaxUrlsPerMinuteCrawlRateInput() *float64
	ProxyConfiguration() KendraDataSourceConfigurationWebCrawlerConfigurationProxyConfigurationOutputReference
	ProxyConfigurationInput() *KendraDataSourceConfigurationWebCrawlerConfigurationProxyConfiguration
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlExclusionPatterns() *[]*string
	SetUrlExclusionPatterns(val *[]*string)
	UrlExclusionPatternsInput() *[]*string
	UrlInclusionPatterns() *[]*string
	SetUrlInclusionPatterns(val *[]*string)
	UrlInclusionPatternsInput() *[]*string
	Urls() KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference
	UrlsInput() *KendraDataSourceConfigurationWebCrawlerConfigurationUrls
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
	PutAuthenticationConfiguration(value *KendraDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfiguration)
	PutProxyConfiguration(value *KendraDataSourceConfigurationWebCrawlerConfigurationProxyConfiguration)
	PutUrls(value *KendraDataSourceConfigurationWebCrawlerConfigurationUrls)
	ResetAuthenticationConfiguration()
	ResetCrawlDepth()
	ResetMaxContentSizePerPageInMegaBytes()
	ResetMaxLinksPerPage()
	ResetMaxUrlsPerMinuteCrawlRate()
	ResetProxyConfiguration()
	ResetUrlExclusionPatterns()
	ResetUrlInclusionPatterns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference
type jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) AuthenticationConfiguration() KendraDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationOutputReference {
	var returns KendraDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationOutputReference
	_jsii_.Get(
		j,
		"authenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) AuthenticationConfigurationInput() *KendraDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfiguration {
	var returns *KendraDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfiguration
	_jsii_.Get(
		j,
		"authenticationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) CrawlDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"crawlDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) CrawlDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"crawlDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) InternalValue() *KendraDataSourceConfigurationWebCrawlerConfiguration {
	var returns *KendraDataSourceConfigurationWebCrawlerConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) MaxContentSizePerPageInMegaBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxContentSizePerPageInMegaBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) MaxContentSizePerPageInMegaBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxContentSizePerPageInMegaBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) MaxLinksPerPage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLinksPerPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) MaxLinksPerPageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLinksPerPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) MaxUrlsPerMinuteCrawlRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUrlsPerMinuteCrawlRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) MaxUrlsPerMinuteCrawlRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUrlsPerMinuteCrawlRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) ProxyConfiguration() KendraDataSourceConfigurationWebCrawlerConfigurationProxyConfigurationOutputReference {
	var returns KendraDataSourceConfigurationWebCrawlerConfigurationProxyConfigurationOutputReference
	_jsii_.Get(
		j,
		"proxyConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) ProxyConfigurationInput() *KendraDataSourceConfigurationWebCrawlerConfigurationProxyConfiguration {
	var returns *KendraDataSourceConfigurationWebCrawlerConfigurationProxyConfiguration
	_jsii_.Get(
		j,
		"proxyConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) UrlExclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlExclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) UrlExclusionPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlExclusionPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) UrlInclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlInclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) UrlInclusionPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlInclusionPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) Urls() KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference {
	var returns KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference
	_jsii_.Get(
		j,
		"urls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) UrlsInput() *KendraDataSourceConfigurationWebCrawlerConfigurationUrls {
	var returns *KendraDataSourceConfigurationWebCrawlerConfigurationUrls
	_jsii_.Get(
		j,
		"urlsInput",
		&returns,
	)
	return returns
}


func NewKendraDataSourceConfigurationWebCrawlerConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceConfigurationWebCrawlerConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kendraDataSource.KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceConfigurationWebCrawlerConfigurationOutputReference_Override(k KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kendraDataSource.KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetCrawlDepth(val *float64) {
	if err := j.validateSetCrawlDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlDepth",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetInternalValue(val *KendraDataSourceConfigurationWebCrawlerConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetMaxContentSizePerPageInMegaBytes(val *float64) {
	if err := j.validateSetMaxContentSizePerPageInMegaBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxContentSizePerPageInMegaBytes",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetMaxLinksPerPage(val *float64) {
	if err := j.validateSetMaxLinksPerPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLinksPerPage",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetMaxUrlsPerMinuteCrawlRate(val *float64) {
	if err := j.validateSetMaxUrlsPerMinuteCrawlRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUrlsPerMinuteCrawlRate",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetUrlExclusionPatterns(val *[]*string) {
	if err := j.validateSetUrlExclusionPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlExclusionPatterns",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference)SetUrlInclusionPatterns(val *[]*string) {
	if err := j.validateSetUrlInclusionPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlInclusionPatterns",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) PutAuthenticationConfiguration(value *KendraDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfiguration) {
	if err := k.validatePutAuthenticationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAuthenticationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) PutProxyConfiguration(value *KendraDataSourceConfigurationWebCrawlerConfigurationProxyConfiguration) {
	if err := k.validatePutProxyConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putProxyConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) PutUrls(value *KendraDataSourceConfigurationWebCrawlerConfigurationUrls) {
	if err := k.validatePutUrlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putUrls",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetAuthenticationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetAuthenticationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetCrawlDepth() {
	_jsii_.InvokeVoid(
		k,
		"resetCrawlDepth",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetMaxContentSizePerPageInMegaBytes() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxContentSizePerPageInMegaBytes",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetMaxLinksPerPage() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxLinksPerPage",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetMaxUrlsPerMinuteCrawlRate() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxUrlsPerMinuteCrawlRate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetProxyConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetProxyConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetUrlExclusionPatterns() {
	_jsii_.InvokeVoid(
		k,
		"resetUrlExclusionPatterns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) ResetUrlInclusionPatterns() {
	_jsii_.InvokeVoid(
		k,
		"resetUrlInclusionPatterns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

