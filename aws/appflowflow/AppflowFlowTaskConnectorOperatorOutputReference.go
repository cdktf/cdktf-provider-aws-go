// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appflowflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppflowFlowTaskConnectorOperatorOutputReference interface {
	cdktf.ComplexObject
	Amplitude() *string
	SetAmplitude(val *string)
	AmplitudeInput() *string
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
	CustomConnector() *string
	SetCustomConnector(val *string)
	CustomConnectorInput() *string
	Datadog() *string
	SetDatadog(val *string)
	DatadogInput() *string
	Dynatrace() *string
	SetDynatrace(val *string)
	DynatraceInput() *string
	// Experimental.
	Fqn() *string
	GoogleAnalytics() *string
	SetGoogleAnalytics(val *string)
	GoogleAnalyticsInput() *string
	InforNexus() *string
	SetInforNexus(val *string)
	InforNexusInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Marketo() *string
	SetMarketo(val *string)
	MarketoInput() *string
	S3() *string
	SetS3(val *string)
	S3Input() *string
	Salesforce() *string
	SetSalesforce(val *string)
	SalesforceInput() *string
	SapoData() *string
	SetSapoData(val *string)
	SapoDataInput() *string
	ServiceNow() *string
	SetServiceNow(val *string)
	ServiceNowInput() *string
	Singular() *string
	SetSingular(val *string)
	SingularInput() *string
	Slack() *string
	SetSlack(val *string)
	SlackInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Trendmicro() *string
	SetTrendmicro(val *string)
	TrendmicroInput() *string
	Veeva() *string
	SetVeeva(val *string)
	VeevaInput() *string
	Zendesk() *string
	SetZendesk(val *string)
	ZendeskInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetAmplitude()
	ResetCustomConnector()
	ResetDatadog()
	ResetDynatrace()
	ResetGoogleAnalytics()
	ResetInforNexus()
	ResetMarketo()
	ResetS3()
	ResetSalesforce()
	ResetSapoData()
	ResetServiceNow()
	ResetSingular()
	ResetSlack()
	ResetTrendmicro()
	ResetVeeva()
	ResetZendesk()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppflowFlowTaskConnectorOperatorOutputReference
type jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) Amplitude() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) AmplitudeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) CustomConnector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) CustomConnectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) Datadog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) DatadogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) Dynatrace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) DynatraceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) GoogleAnalytics() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) GoogleAnalyticsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) InforNexus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) InforNexusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) Marketo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) MarketoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) S3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) S3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) Salesforce() *string {
	var returns *string
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) SalesforceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) SapoData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) SapoDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ServiceNow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ServiceNowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) Singular() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) SingularInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) Slack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) SlackInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) Trendmicro() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) TrendmicroInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) Veeva() *string {
	var returns *string
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) VeevaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) Zendesk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ZendeskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


func NewAppflowFlowTaskConnectorOperatorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppflowFlowTaskConnectorOperatorOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowFlowTaskConnectorOperatorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTaskConnectorOperatorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppflowFlowTaskConnectorOperatorOutputReference_Override(a AppflowFlowTaskConnectorOperatorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTaskConnectorOperatorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetAmplitude(val *string) {
	if err := j.validateSetAmplitudeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amplitude",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetCustomConnector(val *string) {
	if err := j.validateSetCustomConnectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customConnector",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetDatadog(val *string) {
	if err := j.validateSetDatadogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datadog",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetDynatrace(val *string) {
	if err := j.validateSetDynatraceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynatrace",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetGoogleAnalytics(val *string) {
	if err := j.validateSetGoogleAnalyticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleAnalytics",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetInforNexus(val *string) {
	if err := j.validateSetInforNexusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inforNexus",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetMarketo(val *string) {
	if err := j.validateSetMarketoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"marketo",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetS3(val *string) {
	if err := j.validateSetS3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetSalesforce(val *string) {
	if err := j.validateSetSalesforceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"salesforce",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetSapoData(val *string) {
	if err := j.validateSetSapoDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sapoData",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetServiceNow(val *string) {
	if err := j.validateSetServiceNowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceNow",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetSingular(val *string) {
	if err := j.validateSetSingularParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singular",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetSlack(val *string) {
	if err := j.validateSetSlackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slack",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetTrendmicro(val *string) {
	if err := j.validateSetTrendmicroParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trendmicro",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetVeeva(val *string) {
	if err := j.validateSetVeevaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"veeva",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference)SetZendesk(val *string) {
	if err := j.validateSetZendeskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zendesk",
		val,
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		a,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		a,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		a,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		a,
		"resetSingular",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		a,
		"resetSlack",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		a,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		a,
		"resetVeeva",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskConnectorOperatorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

