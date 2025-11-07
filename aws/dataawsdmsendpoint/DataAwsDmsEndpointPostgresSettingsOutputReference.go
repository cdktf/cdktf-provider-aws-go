// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsdmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dataawsdmsendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsDmsEndpointPostgresSettingsOutputReference interface {
	cdktf.ComplexObject
	AfterConnectScript() *string
	AuthenticationMethod() *string
	BabelfishDatabaseName() *string
	CaptureDdls() cdktf.IResolvable
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
	DdlArtifactsSchema() *string
	ExecuteTimeout() *float64
	FailTasksOnLobTruncation() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	HeartbeatEnable() cdktf.IResolvable
	HeartbeatFrequency() *float64
	HeartbeatSchema() *string
	InternalValue() *DataAwsDmsEndpointPostgresSettings
	SetInternalValue(val *DataAwsDmsEndpointPostgresSettings)
	MapBooleanAsBoolean() cdktf.IResolvable
	MapJsonbAsClob() cdktf.IResolvable
	MapLongVarcharAs() *string
	MaxFileSize() *float64
	PluginName() *string
	ServiceAccessRoleArn() *string
	SlotName() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsDmsEndpointPostgresSettingsOutputReference
type jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) AfterConnectScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) AuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) BabelfishDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"babelfishDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) CaptureDdls() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"captureDdls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) DatabaseMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) DdlArtifactsSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddlArtifactsSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) ExecuteTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) FailTasksOnLobTruncation() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"failTasksOnLobTruncation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) HeartbeatEnable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"heartbeatEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) HeartbeatFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) HeartbeatSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"heartbeatSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) InternalValue() *DataAwsDmsEndpointPostgresSettings {
	var returns *DataAwsDmsEndpointPostgresSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) MapBooleanAsBoolean() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mapBooleanAsBoolean",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) MapJsonbAsClob() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mapJsonbAsClob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) MapLongVarcharAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapLongVarcharAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) PluginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) SlotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsDmsEndpointPostgresSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsDmsEndpointPostgresSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsDmsEndpointPostgresSettingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsDmsEndpoint.DataAwsDmsEndpointPostgresSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsDmsEndpointPostgresSettingsOutputReference_Override(d DataAwsDmsEndpointPostgresSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsDmsEndpoint.DataAwsDmsEndpointPostgresSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference)SetInternalValue(val *DataAwsDmsEndpointPostgresSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpointPostgresSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

