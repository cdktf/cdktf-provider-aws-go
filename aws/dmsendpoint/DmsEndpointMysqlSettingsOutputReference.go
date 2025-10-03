// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dmsendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsEndpointMysqlSettingsOutputReference interface {
	cdktf.ComplexObject
	AfterConnectScript() *string
	SetAfterConnectScript(val *string)
	AfterConnectScriptInput() *string
	AuthenticationMethod() *string
	SetAuthenticationMethod(val *string)
	AuthenticationMethodInput() *string
	CleanSourceMetadataOnMismatch() interface{}
	SetCleanSourceMetadataOnMismatch(val interface{})
	CleanSourceMetadataOnMismatchInput() interface{}
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
	EventsPollInterval() *float64
	SetEventsPollInterval(val *float64)
	EventsPollIntervalInput() *float64
	ExecuteTimeout() *float64
	SetExecuteTimeout(val *float64)
	ExecuteTimeoutInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DmsEndpointMysqlSettings
	SetInternalValue(val *DmsEndpointMysqlSettings)
	MaxFileSize() *float64
	SetMaxFileSize(val *float64)
	MaxFileSizeInput() *float64
	ParallelLoadThreads() *float64
	SetParallelLoadThreads(val *float64)
	ParallelLoadThreadsInput() *float64
	ServerTimezone() *string
	SetServerTimezone(val *string)
	ServerTimezoneInput() *string
	ServiceAccessRoleArn() *string
	SetServiceAccessRoleArn(val *string)
	ServiceAccessRoleArnInput() *string
	TargetDbType() *string
	SetTargetDbType(val *string)
	TargetDbTypeInput() *string
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
	ResetAfterConnectScript()
	ResetAuthenticationMethod()
	ResetCleanSourceMetadataOnMismatch()
	ResetEventsPollInterval()
	ResetExecuteTimeout()
	ResetMaxFileSize()
	ResetParallelLoadThreads()
	ResetServerTimezone()
	ResetServiceAccessRoleArn()
	ResetTargetDbType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsEndpointMysqlSettingsOutputReference
type jsiiProxy_DmsEndpointMysqlSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) AfterConnectScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) AfterConnectScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) AuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) AuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) CleanSourceMetadataOnMismatch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanSourceMetadataOnMismatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) CleanSourceMetadataOnMismatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanSourceMetadataOnMismatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) EventsPollInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventsPollInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) EventsPollIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventsPollIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ExecuteTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ExecuteTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) InternalValue() *DmsEndpointMysqlSettings {
	var returns *DmsEndpointMysqlSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ParallelLoadThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelLoadThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ParallelLoadThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelLoadThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ServerTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ServerTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) TargetDbType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) TargetDbTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDmsEndpointMysqlSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DmsEndpointMysqlSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsEndpointMysqlSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpointMysqlSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpointMysqlSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsEndpointMysqlSettingsOutputReference_Override(d DmsEndpointMysqlSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpointMysqlSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetAfterConnectScript(val *string) {
	if err := j.validateSetAfterConnectScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterConnectScript",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetAuthenticationMethod(val *string) {
	if err := j.validateSetAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMethod",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetCleanSourceMetadataOnMismatch(val interface{}) {
	if err := j.validateSetCleanSourceMetadataOnMismatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanSourceMetadataOnMismatch",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetEventsPollInterval(val *float64) {
	if err := j.validateSetEventsPollIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventsPollInterval",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetExecuteTimeout(val *float64) {
	if err := j.validateSetExecuteTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executeTimeout",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetInternalValue(val *DmsEndpointMysqlSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetParallelLoadThreads(val *float64) {
	if err := j.validateSetParallelLoadThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelLoadThreads",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetServerTimezone(val *string) {
	if err := j.validateSetServerTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverTimezone",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetTargetDbType(val *string) {
	if err := j.validateSetTargetDbTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDbType",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMysqlSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ResetAfterConnectScript() {
	_jsii_.InvokeVoid(
		d,
		"resetAfterConnectScript",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ResetAuthenticationMethod() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthenticationMethod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ResetCleanSourceMetadataOnMismatch() {
	_jsii_.InvokeVoid(
		d,
		"resetCleanSourceMetadataOnMismatch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ResetEventsPollInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetEventsPollInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ResetExecuteTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetExecuteTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ResetParallelLoadThreads() {
	_jsii_.InvokeVoid(
		d,
		"resetParallelLoadThreads",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ResetServerTimezone() {
	_jsii_.InvokeVoid(
		d,
		"resetServerTimezone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ResetServiceAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ResetTargetDbType() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetDbType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMysqlSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

