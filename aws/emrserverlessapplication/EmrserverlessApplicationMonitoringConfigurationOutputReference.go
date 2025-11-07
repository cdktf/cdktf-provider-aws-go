// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/emrserverlessapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrserverlessApplicationMonitoringConfigurationOutputReference interface {
	cdktf.ComplexObject
	CloudwatchLoggingConfiguration() EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference
	CloudwatchLoggingConfigurationInput() *EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfiguration
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
	// Experimental.
	Fqn() *string
	InternalValue() *EmrserverlessApplicationMonitoringConfiguration
	SetInternalValue(val *EmrserverlessApplicationMonitoringConfiguration)
	ManagedPersistenceMonitoringConfiguration() EmrserverlessApplicationMonitoringConfigurationManagedPersistenceMonitoringConfigurationOutputReference
	ManagedPersistenceMonitoringConfigurationInput() *EmrserverlessApplicationMonitoringConfigurationManagedPersistenceMonitoringConfiguration
	PrometheusMonitoringConfiguration() EmrserverlessApplicationMonitoringConfigurationPrometheusMonitoringConfigurationOutputReference
	PrometheusMonitoringConfigurationInput() *EmrserverlessApplicationMonitoringConfigurationPrometheusMonitoringConfiguration
	S3MonitoringConfiguration() EmrserverlessApplicationMonitoringConfigurationS3MonitoringConfigurationOutputReference
	S3MonitoringConfigurationInput() *EmrserverlessApplicationMonitoringConfigurationS3MonitoringConfiguration
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
	PutCloudwatchLoggingConfiguration(value *EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfiguration)
	PutManagedPersistenceMonitoringConfiguration(value *EmrserverlessApplicationMonitoringConfigurationManagedPersistenceMonitoringConfiguration)
	PutPrometheusMonitoringConfiguration(value *EmrserverlessApplicationMonitoringConfigurationPrometheusMonitoringConfiguration)
	PutS3MonitoringConfiguration(value *EmrserverlessApplicationMonitoringConfigurationS3MonitoringConfiguration)
	ResetCloudwatchLoggingConfiguration()
	ResetManagedPersistenceMonitoringConfiguration()
	ResetPrometheusMonitoringConfiguration()
	ResetS3MonitoringConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrserverlessApplicationMonitoringConfigurationOutputReference
type jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) CloudwatchLoggingConfiguration() EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference {
	var returns EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) CloudwatchLoggingConfigurationInput() *EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfiguration {
	var returns *EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfiguration
	_jsii_.Get(
		j,
		"cloudwatchLoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) InternalValue() *EmrserverlessApplicationMonitoringConfiguration {
	var returns *EmrserverlessApplicationMonitoringConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) ManagedPersistenceMonitoringConfiguration() EmrserverlessApplicationMonitoringConfigurationManagedPersistenceMonitoringConfigurationOutputReference {
	var returns EmrserverlessApplicationMonitoringConfigurationManagedPersistenceMonitoringConfigurationOutputReference
	_jsii_.Get(
		j,
		"managedPersistenceMonitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) ManagedPersistenceMonitoringConfigurationInput() *EmrserverlessApplicationMonitoringConfigurationManagedPersistenceMonitoringConfiguration {
	var returns *EmrserverlessApplicationMonitoringConfigurationManagedPersistenceMonitoringConfiguration
	_jsii_.Get(
		j,
		"managedPersistenceMonitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) PrometheusMonitoringConfiguration() EmrserverlessApplicationMonitoringConfigurationPrometheusMonitoringConfigurationOutputReference {
	var returns EmrserverlessApplicationMonitoringConfigurationPrometheusMonitoringConfigurationOutputReference
	_jsii_.Get(
		j,
		"prometheusMonitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) PrometheusMonitoringConfigurationInput() *EmrserverlessApplicationMonitoringConfigurationPrometheusMonitoringConfiguration {
	var returns *EmrserverlessApplicationMonitoringConfigurationPrometheusMonitoringConfiguration
	_jsii_.Get(
		j,
		"prometheusMonitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) S3MonitoringConfiguration() EmrserverlessApplicationMonitoringConfigurationS3MonitoringConfigurationOutputReference {
	var returns EmrserverlessApplicationMonitoringConfigurationS3MonitoringConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3MonitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) S3MonitoringConfigurationInput() *EmrserverlessApplicationMonitoringConfigurationS3MonitoringConfiguration {
	var returns *EmrserverlessApplicationMonitoringConfigurationS3MonitoringConfiguration
	_jsii_.Get(
		j,
		"s3MonitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEmrserverlessApplicationMonitoringConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EmrserverlessApplicationMonitoringConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewEmrserverlessApplicationMonitoringConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrserverlessApplication.EmrserverlessApplicationMonitoringConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEmrserverlessApplicationMonitoringConfigurationOutputReference_Override(e EmrserverlessApplicationMonitoringConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrserverlessApplication.EmrserverlessApplicationMonitoringConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference)SetInternalValue(val *EmrserverlessApplicationMonitoringConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) PutCloudwatchLoggingConfiguration(value *EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfiguration) {
	if err := e.validatePutCloudwatchLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCloudwatchLoggingConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) PutManagedPersistenceMonitoringConfiguration(value *EmrserverlessApplicationMonitoringConfigurationManagedPersistenceMonitoringConfiguration) {
	if err := e.validatePutManagedPersistenceMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putManagedPersistenceMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) PutPrometheusMonitoringConfiguration(value *EmrserverlessApplicationMonitoringConfigurationPrometheusMonitoringConfiguration) {
	if err := e.validatePutPrometheusMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPrometheusMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) PutS3MonitoringConfiguration(value *EmrserverlessApplicationMonitoringConfigurationS3MonitoringConfiguration) {
	if err := e.validatePutS3MonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putS3MonitoringConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) ResetCloudwatchLoggingConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetCloudwatchLoggingConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) ResetManagedPersistenceMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetManagedPersistenceMonitoringConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) ResetPrometheusMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetPrometheusMonitoringConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) ResetS3MonitoringConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetS3MonitoringConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EmrserverlessApplicationMonitoringConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

