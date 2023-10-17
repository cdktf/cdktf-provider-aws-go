// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/kinesisanalyticsv2application/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference interface {
	cdktf.ComplexObject
	ApplicationCodeConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationOutputReference
	ApplicationCodeConfigurationInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfiguration
	ApplicationSnapshotConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfigurationOutputReference
	ApplicationSnapshotConfigurationInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfiguration
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
	EnvironmentProperties() Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentPropertiesOutputReference
	EnvironmentPropertiesInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentProperties
	FlinkApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationOutputReference
	FlinkApplicationConfigurationInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfiguration
	// Experimental.
	Fqn() *string
	InternalValue() *Kinesisanalyticsv2ApplicationApplicationConfiguration
	SetInternalValue(val *Kinesisanalyticsv2ApplicationApplicationConfiguration)
	RunConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationRunConfigurationOutputReference
	RunConfigurationInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationRunConfiguration
	SqlApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationOutputReference
	SqlApplicationConfigurationInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfiguration
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationVpcConfigurationOutputReference
	VpcConfigurationInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationVpcConfiguration
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
	PutApplicationCodeConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfiguration)
	PutApplicationSnapshotConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfiguration)
	PutEnvironmentProperties(value *Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentProperties)
	PutFlinkApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfiguration)
	PutRunConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationRunConfiguration)
	PutSqlApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfiguration)
	PutVpcConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationVpcConfiguration)
	ResetApplicationSnapshotConfiguration()
	ResetEnvironmentProperties()
	ResetFlinkApplicationConfiguration()
	ResetRunConfiguration()
	ResetSqlApplicationConfiguration()
	ResetVpcConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference
type jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ApplicationCodeConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationOutputReference
	_jsii_.Get(
		j,
		"applicationCodeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ApplicationCodeConfigurationInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfiguration {
	var returns *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfiguration
	_jsii_.Get(
		j,
		"applicationCodeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ApplicationSnapshotConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfigurationOutputReference
	_jsii_.Get(
		j,
		"applicationSnapshotConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ApplicationSnapshotConfigurationInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfiguration {
	var returns *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfiguration
	_jsii_.Get(
		j,
		"applicationSnapshotConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) EnvironmentProperties() Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentPropertiesOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentPropertiesOutputReference
	_jsii_.Get(
		j,
		"environmentProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) EnvironmentPropertiesInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentProperties {
	var returns *Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentProperties
	_jsii_.Get(
		j,
		"environmentPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) FlinkApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfigurationOutputReference
	_jsii_.Get(
		j,
		"flinkApplicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) FlinkApplicationConfigurationInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfiguration {
	var returns *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfiguration
	_jsii_.Get(
		j,
		"flinkApplicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) InternalValue() *Kinesisanalyticsv2ApplicationApplicationConfiguration {
	var returns *Kinesisanalyticsv2ApplicationApplicationConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) RunConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationRunConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationRunConfigurationOutputReference
	_jsii_.Get(
		j,
		"runConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) RunConfigurationInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationRunConfiguration {
	var returns *Kinesisanalyticsv2ApplicationApplicationConfigurationRunConfiguration
	_jsii_.Get(
		j,
		"runConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) SqlApplicationConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationOutputReference
	_jsii_.Get(
		j,
		"sqlApplicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) SqlApplicationConfigurationInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfiguration {
	var returns *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfiguration
	_jsii_.Get(
		j,
		"sqlApplicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) VpcConfiguration() Kinesisanalyticsv2ApplicationApplicationConfigurationVpcConfigurationOutputReference {
	var returns Kinesisanalyticsv2ApplicationApplicationConfigurationVpcConfigurationOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) VpcConfigurationInput() *Kinesisanalyticsv2ApplicationApplicationConfigurationVpcConfiguration {
	var returns *Kinesisanalyticsv2ApplicationApplicationConfigurationVpcConfiguration
	_jsii_.Get(
		j,
		"vpcConfigurationInput",
		&returns,
	)
	return returns
}


func NewKinesisanalyticsv2ApplicationApplicationConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisanalyticsv2ApplicationApplicationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisanalyticsv2Application.Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisanalyticsv2ApplicationApplicationConfigurationOutputReference_Override(k Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisanalyticsv2Application.Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference)SetInternalValue(val *Kinesisanalyticsv2ApplicationApplicationConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutApplicationCodeConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfiguration) {
	if err := k.validatePutApplicationCodeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putApplicationCodeConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutApplicationSnapshotConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfiguration) {
	if err := k.validatePutApplicationSnapshotConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putApplicationSnapshotConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutEnvironmentProperties(value *Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentProperties) {
	if err := k.validatePutEnvironmentPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putEnvironmentProperties",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutFlinkApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfiguration) {
	if err := k.validatePutFlinkApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putFlinkApplicationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutRunConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationRunConfiguration) {
	if err := k.validatePutRunConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putRunConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutSqlApplicationConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfiguration) {
	if err := k.validatePutSqlApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSqlApplicationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) PutVpcConfiguration(value *Kinesisanalyticsv2ApplicationApplicationConfigurationVpcConfiguration) {
	if err := k.validatePutVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putVpcConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetApplicationSnapshotConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetApplicationSnapshotConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetEnvironmentProperties() {
	_jsii_.InvokeVoid(
		k,
		"resetEnvironmentProperties",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetFlinkApplicationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetFlinkApplicationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetRunConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetRunConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetSqlApplicationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSqlApplicationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ResetVpcConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetVpcConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

