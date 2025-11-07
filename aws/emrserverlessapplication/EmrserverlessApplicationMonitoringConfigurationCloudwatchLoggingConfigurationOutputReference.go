// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/emrserverlessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EncryptionKeyArn() *string
	SetEncryptionKeyArn(val *string)
	EncryptionKeyArnInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfiguration
	SetInternalValue(val *EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfiguration)
	LogGroupName() *string
	SetLogGroupName(val *string)
	LogGroupNameInput() *string
	LogStreamNamePrefix() *string
	SetLogStreamNamePrefix(val *string)
	LogStreamNamePrefixInput() *string
	LogTypes() EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationLogTypesList
	LogTypesInput() interface{}
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
	PutLogTypes(value interface{})
	ResetEncryptionKeyArn()
	ResetLogGroupName()
	ResetLogStreamNamePrefix()
	ResetLogTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference
type jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) EncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) EncryptionKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) InternalValue() *EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfiguration {
	var returns *EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) LogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) LogStreamNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) LogStreamNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) LogTypes() EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationLogTypesList {
	var returns EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationLogTypesList
	_jsii_.Get(
		j,
		"logTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) LogTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewEmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrserverlessApplication.EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference_Override(e EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrserverlessApplication.EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference)SetEncryptionKeyArn(val *string) {
	if err := j.validateSetEncryptionKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference)SetInternalValue(val *EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference)SetLogGroupName(val *string) {
	if err := j.validateSetLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference)SetLogStreamNamePrefix(val *string) {
	if err := j.validateSetLogStreamNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStreamNamePrefix",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) PutLogTypes(value interface{}) {
	if err := e.validatePutLogTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLogTypes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) ResetEncryptionKeyArn() {
	_jsii_.InvokeVoid(
		e,
		"resetEncryptionKeyArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) ResetLogGroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetLogGroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) ResetLogStreamNamePrefix() {
	_jsii_.InvokeVoid(
		e,
		"resetLogStreamNamePrefix",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) ResetLogTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetLogTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

