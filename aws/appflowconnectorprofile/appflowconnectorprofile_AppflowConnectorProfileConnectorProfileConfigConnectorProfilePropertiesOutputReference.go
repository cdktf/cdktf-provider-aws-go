package appflowconnectorprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/appflowconnectorprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference interface {
	cdktf.ComplexObject
	Amplitude() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesAmplitudeOutputReference
	AmplitudeInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesAmplitude
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
	CustomConnector() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnectorOutputReference
	CustomConnectorInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnector
	Datadog() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDatadogOutputReference
	DatadogInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDatadog
	Dynatrace() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDynatraceOutputReference
	DynatraceInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDynatrace
	// Experimental.
	Fqn() *string
	GoogleAnalytics() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsOutputReference
	GoogleAnalyticsInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesGoogleAnalytics
	Honeycode() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesHoneycodeOutputReference
	HoneycodeInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesHoneycode
	InforNexus() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesInforNexusOutputReference
	InforNexusInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesInforNexus
	InternalValue() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileProperties
	SetInternalValue(val *AppflowConnectorProfileConnectorProfileConfigConnectorProfileProperties)
	Marketo() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesMarketoOutputReference
	MarketoInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesMarketo
	Redshift() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesRedshiftOutputReference
	RedshiftInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesRedshift
	Salesforce() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSalesforceOutputReference
	SalesforceInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSalesforce
	SapoData() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference
	SapoDataInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData
	ServiceNow() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesServiceNowOutputReference
	ServiceNowInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesServiceNow
	Singular() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSingularOutputReference
	SingularInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSingular
	Slack() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSlackOutputReference
	SlackInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSlack
	Snowflake() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference
	SnowflakeInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Trendmicro() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesTrendmicroOutputReference
	TrendmicroInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesTrendmicro
	Veeva() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesVeevaOutputReference
	VeevaInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesVeeva
	Zendesk() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesZendeskOutputReference
	ZendeskInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesZendesk
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
	PutAmplitude(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesAmplitude)
	PutCustomConnector(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnector)
	PutDatadog(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDatadog)
	PutDynatrace(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDynatrace)
	PutGoogleAnalytics(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesGoogleAnalytics)
	PutHoneycode(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesHoneycode)
	PutInforNexus(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesInforNexus)
	PutMarketo(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesMarketo)
	PutRedshift(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesRedshift)
	PutSalesforce(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSalesforce)
	PutSapoData(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData)
	PutServiceNow(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesServiceNow)
	PutSingular(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSingular)
	PutSlack(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSlack)
	PutSnowflake(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake)
	PutTrendmicro(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesTrendmicro)
	PutVeeva(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesVeeva)
	PutZendesk(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesZendesk)
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

// The jsii proxy struct for AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference
type jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Amplitude() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesAmplitudeOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesAmplitudeOutputReference
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) AmplitudeInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesAmplitude {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesAmplitude
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) CustomConnector() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnectorOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnectorOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) CustomConnectorInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnector {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnector
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Datadog() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDatadogOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDatadogOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) DatadogInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDatadog {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDatadog
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Dynatrace() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDynatraceOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDynatraceOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) DynatraceInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDynatrace {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDynatrace
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) GoogleAnalytics() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsOutputReference
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) GoogleAnalyticsInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesGoogleAnalytics {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesGoogleAnalytics
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Honeycode() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesHoneycodeOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesHoneycodeOutputReference
	_jsii_.Get(
		j,
		"honeycode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) HoneycodeInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesHoneycode {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesHoneycode
	_jsii_.Get(
		j,
		"honeycodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) InforNexus() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesInforNexusOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesInforNexusOutputReference
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) InforNexusInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesInforNexus {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesInforNexus
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) InternalValue() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileProperties {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Marketo() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesMarketoOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesMarketoOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) MarketoInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesMarketo {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesMarketo
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Redshift() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesRedshiftOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesRedshiftOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) RedshiftInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesRedshift {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesRedshift
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Salesforce() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSalesforceOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSalesforceOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) SalesforceInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSalesforce {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSalesforce
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) SapoData() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) SapoDataInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ServiceNow() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesServiceNowOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ServiceNowInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesServiceNow {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesServiceNow
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Singular() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSingularOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSingularOutputReference
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) SingularInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSingular {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSingular
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Slack() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSlackOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSlackOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) SlackInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSlack {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSlack
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Snowflake() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflakeOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) SnowflakeInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Trendmicro() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesTrendmicroOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesTrendmicroOutputReference
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) TrendmicroInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesTrendmicro {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesTrendmicro
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Veeva() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesVeevaOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesVeevaOutputReference
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) VeevaInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesVeeva {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesVeeva
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Zendesk() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesZendeskOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesZendeskOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ZendeskInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesZendesk {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesZendesk
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


func NewAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appflowConnectorProfile.AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference_Override(a AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appflowConnectorProfile.AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference)SetInternalValue(val *AppflowConnectorProfileConnectorProfileConfigConnectorProfileProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutAmplitude(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesAmplitude) {
	if err := a.validatePutAmplitudeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmplitude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutCustomConnector(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnector) {
	if err := a.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutDatadog(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDatadog) {
	if err := a.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatadog",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutDynatrace(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDynatrace) {
	if err := a.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutGoogleAnalytics(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesGoogleAnalytics) {
	if err := a.validatePutGoogleAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleAnalytics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutHoneycode(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesHoneycode) {
	if err := a.validatePutHoneycodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHoneycode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutInforNexus(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesInforNexus) {
	if err := a.validatePutInforNexusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInforNexus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutMarketo(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesMarketo) {
	if err := a.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMarketo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutRedshift(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesRedshift) {
	if err := a.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshift",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutSalesforce(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSalesforce) {
	if err := a.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutSapoData(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData) {
	if err := a.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSapoData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutServiceNow(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesServiceNow) {
	if err := a.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutSingular(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSingular) {
	if err := a.validatePutSingularParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSingular",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutSlack(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSlack) {
	if err := a.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSlack",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutSnowflake(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake) {
	if err := a.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutTrendmicro(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesTrendmicro) {
	if err := a.validatePutTrendmicroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrendmicro",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutVeeva(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesVeeva) {
	if err := a.validatePutVeevaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVeeva",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) PutZendesk(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesZendesk) {
	if err := a.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZendesk",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		a,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		a,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetHoneycode() {
	_jsii_.InvokeVoid(
		a,
		"resetHoneycode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		a,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshift",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		a,
		"resetSingular",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		a,
		"resetSlack",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		a,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		a,
		"resetVeeva",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

