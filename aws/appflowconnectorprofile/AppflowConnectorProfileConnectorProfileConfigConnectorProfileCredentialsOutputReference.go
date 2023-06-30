package appflowconnectorprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/appflowconnectorprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference interface {
	cdktf.ComplexObject
	Amplitude() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference
	AmplitudeInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude
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
	CustomConnector() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOutputReference
	CustomConnectorInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnector
	Datadog() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadogOutputReference
	DatadogInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadog
	Dynatrace() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatraceOutputReference
	DynatraceInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatrace
	// Experimental.
	Fqn() *string
	GoogleAnalytics() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOutputReference
	GoogleAnalyticsInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalytics
	Honeycode() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsHoneycodeOutputReference
	HoneycodeInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsHoneycode
	InforNexus() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexusOutputReference
	InforNexusInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexus
	InternalValue() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentials
	SetInternalValue(val *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentials)
	Marketo() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOutputReference
	MarketoInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketo
	Redshift() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshiftOutputReference
	RedshiftInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshift
	Salesforce() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforceOutputReference
	SalesforceInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforce
	SapoData() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataOutputReference
	SapoDataInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoData
	ServiceNow() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNowOutputReference
	ServiceNowInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNow
	Singular() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSingularOutputReference
	SingularInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSingular
	Slack() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlackOutputReference
	SlackInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlack
	Snowflake() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflakeOutputReference
	SnowflakeInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflake
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Trendmicro() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicroOutputReference
	TrendmicroInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicro
	Veeva() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsVeevaOutputReference
	VeevaInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsVeeva
	Zendesk() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendeskOutputReference
	ZendeskInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendesk
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
	PutAmplitude(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude)
	PutCustomConnector(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnector)
	PutDatadog(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadog)
	PutDynatrace(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatrace)
	PutGoogleAnalytics(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalytics)
	PutHoneycode(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsHoneycode)
	PutInforNexus(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexus)
	PutMarketo(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketo)
	PutRedshift(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshift)
	PutSalesforce(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforce)
	PutSapoData(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoData)
	PutServiceNow(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNow)
	PutSingular(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSingular)
	PutSlack(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlack)
	PutSnowflake(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflake)
	PutTrendmicro(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicro)
	PutVeeva(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsVeeva)
	PutZendesk(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendesk)
	ResetAmplitude()
	ResetCustomConnector()
	ResetDatadog()
	ResetDynatrace()
	ResetGoogleAnalytics()
	ResetHoneycode()
	ResetInforNexus()
	ResetMarketo()
	ResetRedshift()
	ResetSalesforce()
	ResetSapoData()
	ResetServiceNow()
	ResetSingular()
	ResetSlack()
	ResetSnowflake()
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

// The jsii proxy struct for AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference
type jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Amplitude() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) AmplitudeInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) CustomConnector() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) CustomConnectorInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnector {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnector
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Datadog() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadogOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadogOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) DatadogInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadog {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadog
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Dynatrace() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatraceOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatraceOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) DynatraceInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatrace {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatrace
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GoogleAnalytics() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOutputReference
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GoogleAnalyticsInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalytics {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalytics
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Honeycode() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsHoneycodeOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsHoneycodeOutputReference
	_jsii_.Get(
		j,
		"honeycode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) HoneycodeInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsHoneycode {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsHoneycode
	_jsii_.Get(
		j,
		"honeycodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) InforNexus() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexusOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexusOutputReference
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) InforNexusInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexus {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexus
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) InternalValue() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentials {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentials
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Marketo() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) MarketoInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketo {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketo
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Redshift() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshiftOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshiftOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) RedshiftInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshift {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshift
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Salesforce() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforceOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforceOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) SalesforceInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforce {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforce
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) SapoData() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) SapoDataInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoData {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoData
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ServiceNow() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNowOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ServiceNowInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNow {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNow
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Singular() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSingularOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSingularOutputReference
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) SingularInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSingular {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSingular
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Slack() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlackOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlackOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) SlackInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlack {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlack
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Snowflake() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflakeOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflakeOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) SnowflakeInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflake {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflake
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Trendmicro() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicroOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicroOutputReference
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) TrendmicroInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicro {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicro
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Veeva() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsVeevaOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsVeevaOutputReference
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) VeevaInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsVeeva {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsVeeva
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Zendesk() AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendeskOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendeskOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ZendeskInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendesk {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendesk
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


func NewAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appflowConnectorProfile.AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference_Override(a AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appflowConnectorProfile.AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference)SetInternalValue(val *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentials) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutAmplitude(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude) {
	if err := a.validatePutAmplitudeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmplitude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutCustomConnector(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnector) {
	if err := a.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutDatadog(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadog) {
	if err := a.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatadog",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutDynatrace(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatrace) {
	if err := a.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutGoogleAnalytics(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalytics) {
	if err := a.validatePutGoogleAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleAnalytics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutHoneycode(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsHoneycode) {
	if err := a.validatePutHoneycodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHoneycode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutInforNexus(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexus) {
	if err := a.validatePutInforNexusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInforNexus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutMarketo(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketo) {
	if err := a.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMarketo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutRedshift(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshift) {
	if err := a.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshift",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutSalesforce(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforce) {
	if err := a.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutSapoData(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoData) {
	if err := a.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSapoData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutServiceNow(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNow) {
	if err := a.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutSingular(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSingular) {
	if err := a.validatePutSingularParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSingular",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutSlack(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlack) {
	if err := a.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSlack",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutSnowflake(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflake) {
	if err := a.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutTrendmicro(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicro) {
	if err := a.validatePutTrendmicroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrendmicro",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutVeeva(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsVeeva) {
	if err := a.validatePutVeevaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVeeva",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) PutZendesk(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendesk) {
	if err := a.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZendesk",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		a,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		a,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetHoneycode() {
	_jsii_.InvokeVoid(
		a,
		"resetHoneycode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		a,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshift",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		a,
		"resetSingular",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		a,
		"resetSlack",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		a,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		a,
		"resetVeeva",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

