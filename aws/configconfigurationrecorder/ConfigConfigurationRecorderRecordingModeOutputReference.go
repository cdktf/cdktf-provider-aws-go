// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configconfigurationrecorder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/configconfigurationrecorder/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigConfigurationRecorderRecordingModeOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *ConfigConfigurationRecorderRecordingMode
	SetInternalValue(val *ConfigConfigurationRecorderRecordingMode)
	RecordingFrequency() *string
	SetRecordingFrequency(val *string)
	RecordingFrequencyInput() *string
	RecordingModeOverride() ConfigConfigurationRecorderRecordingModeRecordingModeOverrideOutputReference
	RecordingModeOverrideInput() *ConfigConfigurationRecorderRecordingModeRecordingModeOverride
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
	PutRecordingModeOverride(value *ConfigConfigurationRecorderRecordingModeRecordingModeOverride)
	ResetRecordingFrequency()
	ResetRecordingModeOverride()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigConfigurationRecorderRecordingModeOutputReference
type jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) InternalValue() *ConfigConfigurationRecorderRecordingMode {
	var returns *ConfigConfigurationRecorderRecordingMode
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) RecordingFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordingFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) RecordingFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordingFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) RecordingModeOverride() ConfigConfigurationRecorderRecordingModeRecordingModeOverrideOutputReference {
	var returns ConfigConfigurationRecorderRecordingModeRecordingModeOverrideOutputReference
	_jsii_.Get(
		j,
		"recordingModeOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) RecordingModeOverrideInput() *ConfigConfigurationRecorderRecordingModeRecordingModeOverride {
	var returns *ConfigConfigurationRecorderRecordingModeRecordingModeOverride
	_jsii_.Get(
		j,
		"recordingModeOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigConfigurationRecorderRecordingModeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConfigConfigurationRecorderRecordingModeOutputReference {
	_init_.Initialize()

	if err := validateNewConfigConfigurationRecorderRecordingModeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.configConfigurationRecorder.ConfigConfigurationRecorderRecordingModeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConfigConfigurationRecorderRecordingModeOutputReference_Override(c ConfigConfigurationRecorderRecordingModeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.configConfigurationRecorder.ConfigConfigurationRecorderRecordingModeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference)SetInternalValue(val *ConfigConfigurationRecorderRecordingMode) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference)SetRecordingFrequency(val *string) {
	if err := j.validateSetRecordingFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordingFrequency",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) PutRecordingModeOverride(value *ConfigConfigurationRecorderRecordingModeRecordingModeOverride) {
	if err := c.validatePutRecordingModeOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRecordingModeOverride",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) ResetRecordingFrequency() {
	_jsii_.InvokeVoid(
		c,
		"resetRecordingFrequency",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) ResetRecordingModeOverride() {
	_jsii_.InvokeVoid(
		c,
		"resetRecordingModeOverride",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingModeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

