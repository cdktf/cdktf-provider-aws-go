// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dmsendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsEndpointPostgresSettingsOutputReference interface {
	cdktf.ComplexObject
	AfterConnectScript() *string
	SetAfterConnectScript(val *string)
	AfterConnectScriptInput() *string
	BabelfishDatabaseName() *string
	SetBabelfishDatabaseName(val *string)
	BabelfishDatabaseNameInput() *string
	CaptureDdls() interface{}
	SetCaptureDdls(val interface{})
	CaptureDdlsInput() interface{}
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
	DatabaseMode() *string
	SetDatabaseMode(val *string)
	DatabaseModeInput() *string
	DdlArtifactsSchema() *string
	SetDdlArtifactsSchema(val *string)
	DdlArtifactsSchemaInput() *string
	ExecuteTimeout() *float64
	SetExecuteTimeout(val *float64)
	ExecuteTimeoutInput() *float64
	FailTasksOnLobTruncation() interface{}
	SetFailTasksOnLobTruncation(val interface{})
	FailTasksOnLobTruncationInput() interface{}
	// Experimental.
	Fqn() *string
	HeartbeatEnable() interface{}
	SetHeartbeatEnable(val interface{})
	HeartbeatEnableInput() interface{}
	HeartbeatFrequency() *float64
	SetHeartbeatFrequency(val *float64)
	HeartbeatFrequencyInput() *float64
	HeartbeatSchema() *string
	SetHeartbeatSchema(val *string)
	HeartbeatSchemaInput() *string
	InternalValue() *DmsEndpointPostgresSettings
	SetInternalValue(val *DmsEndpointPostgresSettings)
	MapBooleanAsBoolean() interface{}
	SetMapBooleanAsBoolean(val interface{})
	MapBooleanAsBooleanInput() interface{}
	MapJsonbAsClob() interface{}
	SetMapJsonbAsClob(val interface{})
	MapJsonbAsClobInput() interface{}
	MapLongVarcharAs() *string
	SetMapLongVarcharAs(val *string)
	MapLongVarcharAsInput() *string
	MaxFileSize() *float64
	SetMaxFileSize(val *float64)
	MaxFileSizeInput() *float64
	PluginName() *string
	SetPluginName(val *string)
	PluginNameInput() *string
	SlotName() *string
	SetSlotName(val *string)
	SlotNameInput() *string
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
	ResetBabelfishDatabaseName()
	ResetCaptureDdls()
	ResetDatabaseMode()
	ResetDdlArtifactsSchema()
	ResetExecuteTimeout()
	ResetFailTasksOnLobTruncation()
	ResetHeartbeatEnable()
	ResetHeartbeatFrequency()
	ResetHeartbeatSchema()
	ResetMapBooleanAsBoolean()
	ResetMapJsonbAsClob()
	ResetMapLongVarcharAs()
	ResetMaxFileSize()
	ResetPluginName()
	ResetSlotName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsEndpointPostgresSettingsOutputReference
type jsiiProxy_DmsEndpointPostgresSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) AfterConnectScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) AfterConnectScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) BabelfishDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"babelfishDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) BabelfishDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"babelfishDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) CaptureDdls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captureDdls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) CaptureDdlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captureDdlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) DatabaseMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) DatabaseModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) DdlArtifactsSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddlArtifactsSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) DdlArtifactsSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddlArtifactsSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ExecuteTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ExecuteTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) FailTasksOnLobTruncation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failTasksOnLobTruncation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) FailTasksOnLobTruncationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failTasksOnLobTruncationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) HeartbeatEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"heartbeatEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) HeartbeatEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"heartbeatEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) HeartbeatFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) HeartbeatFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) HeartbeatSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"heartbeatSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) HeartbeatSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"heartbeatSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) InternalValue() *DmsEndpointPostgresSettings {
	var returns *DmsEndpointPostgresSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) MapBooleanAsBoolean() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapBooleanAsBoolean",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) MapBooleanAsBooleanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapBooleanAsBooleanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) MapJsonbAsClob() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapJsonbAsClob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) MapJsonbAsClobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapJsonbAsClobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) MapLongVarcharAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapLongVarcharAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) MapLongVarcharAsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapLongVarcharAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) PluginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) PluginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) SlotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) SlotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDmsEndpointPostgresSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DmsEndpointPostgresSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsEndpointPostgresSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpointPostgresSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpointPostgresSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsEndpointPostgresSettingsOutputReference_Override(d DmsEndpointPostgresSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpointPostgresSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetAfterConnectScript(val *string) {
	if err := j.validateSetAfterConnectScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterConnectScript",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetBabelfishDatabaseName(val *string) {
	if err := j.validateSetBabelfishDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"babelfishDatabaseName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetCaptureDdls(val interface{}) {
	if err := j.validateSetCaptureDdlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captureDdls",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetDatabaseMode(val *string) {
	if err := j.validateSetDatabaseModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseMode",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetDdlArtifactsSchema(val *string) {
	if err := j.validateSetDdlArtifactsSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ddlArtifactsSchema",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetExecuteTimeout(val *float64) {
	if err := j.validateSetExecuteTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executeTimeout",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetFailTasksOnLobTruncation(val interface{}) {
	if err := j.validateSetFailTasksOnLobTruncationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failTasksOnLobTruncation",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetHeartbeatEnable(val interface{}) {
	if err := j.validateSetHeartbeatEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatEnable",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetHeartbeatFrequency(val *float64) {
	if err := j.validateSetHeartbeatFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatFrequency",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetHeartbeatSchema(val *string) {
	if err := j.validateSetHeartbeatSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatSchema",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetInternalValue(val *DmsEndpointPostgresSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetMapBooleanAsBoolean(val interface{}) {
	if err := j.validateSetMapBooleanAsBooleanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapBooleanAsBoolean",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetMapJsonbAsClob(val interface{}) {
	if err := j.validateSetMapJsonbAsClobParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapJsonbAsClob",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetMapLongVarcharAs(val *string) {
	if err := j.validateSetMapLongVarcharAsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapLongVarcharAs",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetPluginName(val *string) {
	if err := j.validateSetPluginNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetSlotName(val *string) {
	if err := j.validateSetSlotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slotName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointPostgresSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetAfterConnectScript() {
	_jsii_.InvokeVoid(
		d,
		"resetAfterConnectScript",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetBabelfishDatabaseName() {
	_jsii_.InvokeVoid(
		d,
		"resetBabelfishDatabaseName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetCaptureDdls() {
	_jsii_.InvokeVoid(
		d,
		"resetCaptureDdls",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetDatabaseMode() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetDdlArtifactsSchema() {
	_jsii_.InvokeVoid(
		d,
		"resetDdlArtifactsSchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetExecuteTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetExecuteTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetFailTasksOnLobTruncation() {
	_jsii_.InvokeVoid(
		d,
		"resetFailTasksOnLobTruncation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetHeartbeatEnable() {
	_jsii_.InvokeVoid(
		d,
		"resetHeartbeatEnable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetHeartbeatFrequency() {
	_jsii_.InvokeVoid(
		d,
		"resetHeartbeatFrequency",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetHeartbeatSchema() {
	_jsii_.InvokeVoid(
		d,
		"resetHeartbeatSchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetMapBooleanAsBoolean() {
	_jsii_.InvokeVoid(
		d,
		"resetMapBooleanAsBoolean",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetMapJsonbAsClob() {
	_jsii_.InvokeVoid(
		d,
		"resetMapJsonbAsClob",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetMapLongVarcharAs() {
	_jsii_.InvokeVoid(
		d,
		"resetMapLongVarcharAs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetPluginName() {
	_jsii_.InvokeVoid(
		d,
		"resetPluginName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ResetSlotName() {
	_jsii_.InvokeVoid(
		d,
		"resetSlotName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DmsEndpointPostgresSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

