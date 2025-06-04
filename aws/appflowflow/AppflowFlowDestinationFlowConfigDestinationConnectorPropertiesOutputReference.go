// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/appflowflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference interface {
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
	CustomConnector() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorOutputReference
	CustomConnectorInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnector
	CustomerProfiles() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfilesOutputReference
	CustomerProfilesInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfiles
	EventBridge() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridgeOutputReference
	EventBridgeInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridge
	// Experimental.
	Fqn() *string
	Honeycode() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycodeOutputReference
	HoneycodeInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycode
	InternalValue() *AppflowFlowDestinationFlowConfigDestinationConnectorProperties
	SetInternalValue(val *AppflowFlowDestinationFlowConfigDestinationConnectorProperties)
	LookoutMetrics() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetricsOutputReference
	LookoutMetricsInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetrics
	Marketo() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketoOutputReference
	MarketoInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketo
	Redshift() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshiftOutputReference
	RedshiftInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshift
	S3() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3OutputReference
	S3Input() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3
	Salesforce() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference
	SalesforceInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce
	SapoData() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataOutputReference
	SapoDataInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoData
	Snowflake() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflakeOutputReference
	SnowflakeInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflake
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Upsolver() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverOutputReference
	UpsolverInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolver
	Zendesk() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendeskOutputReference
	ZendeskInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendesk
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
	PutCustomConnector(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnector)
	PutCustomerProfiles(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfiles)
	PutEventBridge(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridge)
	PutHoneycode(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycode)
	PutLookoutMetrics(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetrics)
	PutMarketo(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketo)
	PutRedshift(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshift)
	PutS3(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3)
	PutSalesforce(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce)
	PutSapoData(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoData)
	PutSnowflake(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflake)
	PutUpsolver(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolver)
	PutZendesk(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendesk)
	ResetCustomConnector()
	ResetCustomerProfiles()
	ResetEventBridge()
	ResetHoneycode()
	ResetLookoutMetrics()
	ResetMarketo()
	ResetRedshift()
	ResetS3()
	ResetSalesforce()
	ResetSapoData()
	ResetSnowflake()
	ResetUpsolver()
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

// The jsii proxy struct for AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference
type jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) CustomConnector() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorOutputReference {
	var returns AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) CustomConnectorInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnector {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnector
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) CustomerProfiles() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfilesOutputReference {
	var returns AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfilesOutputReference
	_jsii_.Get(
		j,
		"customerProfiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) CustomerProfilesInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfiles {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfiles
	_jsii_.Get(
		j,
		"customerProfilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) EventBridge() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridgeOutputReference {
	var returns AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridgeOutputReference
	_jsii_.Get(
		j,
		"eventBridge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) EventBridgeInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridge {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridge
	_jsii_.Get(
		j,
		"eventBridgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) Honeycode() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycodeOutputReference {
	var returns AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycodeOutputReference
	_jsii_.Get(
		j,
		"honeycode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) HoneycodeInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycode {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycode
	_jsii_.Get(
		j,
		"honeycodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) InternalValue() *AppflowFlowDestinationFlowConfigDestinationConnectorProperties {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) LookoutMetrics() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetricsOutputReference {
	var returns AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetricsOutputReference
	_jsii_.Get(
		j,
		"lookoutMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) LookoutMetricsInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetrics {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetrics
	_jsii_.Get(
		j,
		"lookoutMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) Marketo() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketoOutputReference {
	var returns AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketoOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) MarketoInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketo {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketo
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) Redshift() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshiftOutputReference {
	var returns AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshiftOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) RedshiftInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshift {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshift
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) S3() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3OutputReference {
	var returns AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) S3Input() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3 {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) Salesforce() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference {
	var returns AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) SalesforceInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) SapoData() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataOutputReference {
	var returns AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) SapoDataInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoData {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoData
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) Snowflake() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflakeOutputReference {
	var returns AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflakeOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) SnowflakeInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflake {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflake
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) Upsolver() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverOutputReference {
	var returns AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverOutputReference
	_jsii_.Get(
		j,
		"upsolver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) UpsolverInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolver {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolver
	_jsii_.Get(
		j,
		"upsolverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) Zendesk() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendeskOutputReference {
	var returns AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendeskOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ZendeskInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendesk {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendesk
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


func NewAppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference_Override(a AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference)SetInternalValue(val *AppflowFlowDestinationFlowConfigDestinationConnectorProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) PutCustomConnector(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnector) {
	if err := a.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) PutCustomerProfiles(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfiles) {
	if err := a.validatePutCustomerProfilesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomerProfiles",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) PutEventBridge(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesEventBridge) {
	if err := a.validatePutEventBridgeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventBridge",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) PutHoneycode(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesHoneycode) {
	if err := a.validatePutHoneycodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHoneycode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) PutLookoutMetrics(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetrics) {
	if err := a.validatePutLookoutMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLookoutMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) PutMarketo(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesMarketo) {
	if err := a.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMarketo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) PutRedshift(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesRedshift) {
	if err := a.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshift",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) PutS3(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3) {
	if err := a.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) PutSalesforce(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce) {
	if err := a.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) PutSapoData(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSapoData) {
	if err := a.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSapoData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) PutSnowflake(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSnowflake) {
	if err := a.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) PutUpsolver(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolver) {
	if err := a.validatePutUpsolverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUpsolver",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) PutZendesk(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendesk) {
	if err := a.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZendesk",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ResetCustomerProfiles() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerProfiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ResetEventBridge() {
	_jsii_.InvokeVoid(
		a,
		"resetEventBridge",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ResetHoneycode() {
	_jsii_.InvokeVoid(
		a,
		"resetHoneycode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ResetLookoutMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetLookoutMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshift",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ResetUpsolver() {
	_jsii_.InvokeVoid(
		a,
		"resetUpsolver",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

