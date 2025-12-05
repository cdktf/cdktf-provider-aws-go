// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/athenaworkgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AthenaWorkgroupConfigurationOutputReference interface {
	cdktf.ComplexObject
	BytesScannedCutoffPerQuery() *float64
	SetBytesScannedCutoffPerQuery(val *float64)
	BytesScannedCutoffPerQueryInput() *float64
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
	EnforceWorkgroupConfiguration() interface{}
	SetEnforceWorkgroupConfiguration(val interface{})
	EnforceWorkgroupConfigurationInput() interface{}
	EngineVersion() AthenaWorkgroupConfigurationEngineVersionOutputReference
	EngineVersionInput() *AthenaWorkgroupConfigurationEngineVersion
	ExecutionRole() *string
	SetExecutionRole(val *string)
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	IdentityCenterConfiguration() AthenaWorkgroupConfigurationIdentityCenterConfigurationOutputReference
	IdentityCenterConfigurationInput() *AthenaWorkgroupConfigurationIdentityCenterConfiguration
	InternalValue() *AthenaWorkgroupConfiguration
	SetInternalValue(val *AthenaWorkgroupConfiguration)
	ManagedQueryResultsConfiguration() AthenaWorkgroupConfigurationManagedQueryResultsConfigurationOutputReference
	ManagedQueryResultsConfigurationInput() *AthenaWorkgroupConfigurationManagedQueryResultsConfiguration
	PublishCloudwatchMetricsEnabled() interface{}
	SetPublishCloudwatchMetricsEnabled(val interface{})
	PublishCloudwatchMetricsEnabledInput() interface{}
	RequesterPaysEnabled() interface{}
	SetRequesterPaysEnabled(val interface{})
	RequesterPaysEnabledInput() interface{}
	ResultConfiguration() AthenaWorkgroupConfigurationResultConfigurationOutputReference
	ResultConfigurationInput() *AthenaWorkgroupConfigurationResultConfiguration
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutEngineVersion(value *AthenaWorkgroupConfigurationEngineVersion)
	PutIdentityCenterConfiguration(value *AthenaWorkgroupConfigurationIdentityCenterConfiguration)
	PutManagedQueryResultsConfiguration(value *AthenaWorkgroupConfigurationManagedQueryResultsConfiguration)
	PutResultConfiguration(value *AthenaWorkgroupConfigurationResultConfiguration)
	ResetBytesScannedCutoffPerQuery()
	ResetEnforceWorkgroupConfiguration()
	ResetEngineVersion()
	ResetExecutionRole()
	ResetIdentityCenterConfiguration()
	ResetManagedQueryResultsConfiguration()
	ResetPublishCloudwatchMetricsEnabled()
	ResetRequesterPaysEnabled()
	ResetResultConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AthenaWorkgroupConfigurationOutputReference
type jsiiProxy_AthenaWorkgroupConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) BytesScannedCutoffPerQuery() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesScannedCutoffPerQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) BytesScannedCutoffPerQueryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesScannedCutoffPerQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) EnforceWorkgroupConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceWorkgroupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) EnforceWorkgroupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceWorkgroupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) EngineVersion() AthenaWorkgroupConfigurationEngineVersionOutputReference {
	var returns AthenaWorkgroupConfigurationEngineVersionOutputReference
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) EngineVersionInput() *AthenaWorkgroupConfigurationEngineVersion {
	var returns *AthenaWorkgroupConfigurationEngineVersion
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) IdentityCenterConfiguration() AthenaWorkgroupConfigurationIdentityCenterConfigurationOutputReference {
	var returns AthenaWorkgroupConfigurationIdentityCenterConfigurationOutputReference
	_jsii_.Get(
		j,
		"identityCenterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) IdentityCenterConfigurationInput() *AthenaWorkgroupConfigurationIdentityCenterConfiguration {
	var returns *AthenaWorkgroupConfigurationIdentityCenterConfiguration
	_jsii_.Get(
		j,
		"identityCenterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) InternalValue() *AthenaWorkgroupConfiguration {
	var returns *AthenaWorkgroupConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ManagedQueryResultsConfiguration() AthenaWorkgroupConfigurationManagedQueryResultsConfigurationOutputReference {
	var returns AthenaWorkgroupConfigurationManagedQueryResultsConfigurationOutputReference
	_jsii_.Get(
		j,
		"managedQueryResultsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ManagedQueryResultsConfigurationInput() *AthenaWorkgroupConfigurationManagedQueryResultsConfiguration {
	var returns *AthenaWorkgroupConfigurationManagedQueryResultsConfiguration
	_jsii_.Get(
		j,
		"managedQueryResultsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) PublishCloudwatchMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishCloudwatchMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) PublishCloudwatchMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishCloudwatchMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) RequesterPaysEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) RequesterPaysEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ResultConfiguration() AthenaWorkgroupConfigurationResultConfigurationOutputReference {
	var returns AthenaWorkgroupConfigurationResultConfigurationOutputReference
	_jsii_.Get(
		j,
		"resultConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ResultConfigurationInput() *AthenaWorkgroupConfigurationResultConfiguration {
	var returns *AthenaWorkgroupConfigurationResultConfiguration
	_jsii_.Get(
		j,
		"resultConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAthenaWorkgroupConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AthenaWorkgroupConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewAthenaWorkgroupConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AthenaWorkgroupConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.athenaWorkgroup.AthenaWorkgroupConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAthenaWorkgroupConfigurationOutputReference_Override(a AthenaWorkgroupConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.athenaWorkgroup.AthenaWorkgroupConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference)SetBytesScannedCutoffPerQuery(val *float64) {
	if err := j.validateSetBytesScannedCutoffPerQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bytesScannedCutoffPerQuery",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference)SetEnforceWorkgroupConfiguration(val interface{}) {
	if err := j.validateSetEnforceWorkgroupConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceWorkgroupConfiguration",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference)SetInternalValue(val *AthenaWorkgroupConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference)SetPublishCloudwatchMetricsEnabled(val interface{}) {
	if err := j.validateSetPublishCloudwatchMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishCloudwatchMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference)SetRequesterPaysEnabled(val interface{}) {
	if err := j.validateSetRequesterPaysEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requesterPaysEnabled",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkgroupConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) PutEngineVersion(value *AthenaWorkgroupConfigurationEngineVersion) {
	if err := a.validatePutEngineVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEngineVersion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) PutIdentityCenterConfiguration(value *AthenaWorkgroupConfigurationIdentityCenterConfiguration) {
	if err := a.validatePutIdentityCenterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIdentityCenterConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) PutManagedQueryResultsConfiguration(value *AthenaWorkgroupConfigurationManagedQueryResultsConfiguration) {
	if err := a.validatePutManagedQueryResultsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedQueryResultsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) PutResultConfiguration(value *AthenaWorkgroupConfigurationResultConfiguration) {
	if err := a.validatePutResultConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResultConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ResetBytesScannedCutoffPerQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetBytesScannedCutoffPerQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ResetEnforceWorkgroupConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEnforceWorkgroupConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ResetExecutionRole() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ResetIdentityCenterConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityCenterConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ResetManagedQueryResultsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedQueryResultsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ResetPublishCloudwatchMetricsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPublishCloudwatchMetricsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ResetRequesterPaysEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRequesterPaysEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ResetResultConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetResultConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AthenaWorkgroupConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

