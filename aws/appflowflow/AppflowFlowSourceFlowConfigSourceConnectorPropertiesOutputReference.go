package appflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/appflowflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference interface {
	cdktf.ComplexObject
	Amplitude() AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitudeOutputReference
	AmplitudeInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitude
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
	CustomConnector() AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnectorOutputReference
	CustomConnectorInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnector
	Datadog() AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadogOutputReference
	DatadogInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadog
	Dynatrace() AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatraceOutputReference
	DynatraceInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatrace
	// Experimental.
	Fqn() *string
	GoogleAnalytics() AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalyticsOutputReference
	GoogleAnalyticsInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalytics
	InforNexus() AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexusOutputReference
	InforNexusInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexus
	InternalValue() *AppflowFlowSourceFlowConfigSourceConnectorProperties
	SetInternalValue(val *AppflowFlowSourceFlowConfigSourceConnectorProperties)
	Marketo() AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketoOutputReference
	MarketoInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketo
	S3() AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3OutputReference
	S3Input() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3
	Salesforce() AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforceOutputReference
	SalesforceInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforce
	SapoData() AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference
	SapoDataInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData
	ServiceNow() AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNowOutputReference
	ServiceNowInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNow
	Singular() AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingularOutputReference
	SingularInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingular
	Slack() AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlackOutputReference
	SlackInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlack
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Trendmicro() AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicroOutputReference
	TrendmicroInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicro
	Veeva() AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference
	VeevaInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva
	Zendesk() AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendeskOutputReference
	ZendeskInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendesk
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
	PutAmplitude(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitude)
	PutCustomConnector(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnector)
	PutDatadog(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadog)
	PutDynatrace(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatrace)
	PutGoogleAnalytics(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalytics)
	PutInforNexus(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexus)
	PutMarketo(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketo)
	PutS3(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3)
	PutSalesforce(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforce)
	PutSapoData(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData)
	PutServiceNow(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNow)
	PutSingular(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingular)
	PutSlack(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlack)
	PutTrendmicro(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicro)
	PutVeeva(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva)
	PutZendesk(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendesk)
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
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference
type jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Amplitude() AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitudeOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitudeOutputReference
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) AmplitudeInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitude {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitude
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) CustomConnector() AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnectorOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnectorOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) CustomConnectorInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnector {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnector
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Datadog() AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadogOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadogOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) DatadogInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadog {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadog
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Dynatrace() AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatraceOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatraceOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) DynatraceInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatrace {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatrace
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GoogleAnalytics() AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalyticsOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalyticsOutputReference
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GoogleAnalyticsInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalytics {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalytics
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) InforNexus() AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexusOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexusOutputReference
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) InforNexusInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexus {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexus
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) InternalValue() *AppflowFlowSourceFlowConfigSourceConnectorProperties {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Marketo() AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketoOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketoOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) MarketoInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketo {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketo
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) S3() AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3OutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) S3Input() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3 {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Salesforce() AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforceOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforceOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) SalesforceInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforce {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforce
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) SapoData() AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) SapoDataInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ServiceNow() AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNowOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ServiceNowInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNow {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNow
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Singular() AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingularOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingularOutputReference
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) SingularInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingular {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingular
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Slack() AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlackOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlackOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) SlackInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlack {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlack
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Trendmicro() AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicroOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicroOutputReference
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) TrendmicroInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicro {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicro
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Veeva() AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) VeevaInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Zendesk() AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendeskOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendeskOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ZendeskInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendesk {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendesk
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


func NewAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference_Override(a AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference)SetInternalValue(val *AppflowFlowSourceFlowConfigSourceConnectorProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutAmplitude(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitude) {
	if err := a.validatePutAmplitudeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmplitude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutCustomConnector(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnector) {
	if err := a.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutDatadog(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadog) {
	if err := a.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatadog",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutDynatrace(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatrace) {
	if err := a.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutGoogleAnalytics(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalytics) {
	if err := a.validatePutGoogleAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleAnalytics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutInforNexus(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexus) {
	if err := a.validatePutInforNexusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInforNexus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutMarketo(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketo) {
	if err := a.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMarketo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutS3(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3) {
	if err := a.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutSalesforce(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforce) {
	if err := a.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutSapoData(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData) {
	if err := a.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSapoData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutServiceNow(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNow) {
	if err := a.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutSingular(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingular) {
	if err := a.validatePutSingularParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSingular",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutSlack(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlack) {
	if err := a.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSlack",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutTrendmicro(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicro) {
	if err := a.validatePutTrendmicroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrendmicro",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutVeeva(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva) {
	if err := a.validatePutVeevaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVeeva",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) PutZendesk(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendesk) {
	if err := a.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZendesk",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		a,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		a,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		a,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		a,
		"resetSingular",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		a,
		"resetSlack",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		a,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		a,
		"resetVeeva",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

