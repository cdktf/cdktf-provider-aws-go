// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference interface {
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
	InitialAudioGain() *float64
	SetInitialAudioGain(val *float64)
	InitialAudioGainInput() *float64
	InputEndAction() *string
	SetInputEndAction(val *string)
	InputEndActionInput() *string
	InputLossBehavior() MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference
	InputLossBehaviorInput() *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehavior
	InternalValue() *MedialiveChannelEncoderSettingsGlobalConfiguration
	SetInternalValue(val *MedialiveChannelEncoderSettingsGlobalConfiguration)
	OutputLockingMode() *string
	SetOutputLockingMode(val *string)
	OutputLockingModeInput() *string
	OutputTimingSource() *string
	SetOutputTimingSource(val *string)
	OutputTimingSourceInput() *string
	SupportLowFramerateInputs() *string
	SetSupportLowFramerateInputs(val *string)
	SupportLowFramerateInputsInput() *string
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
	PutInputLossBehavior(value *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehavior)
	ResetInitialAudioGain()
	ResetInputEndAction()
	ResetInputLossBehavior()
	ResetOutputLockingMode()
	ResetOutputTimingSource()
	ResetSupportLowFramerateInputs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) InitialAudioGain() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialAudioGain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) InitialAudioGainInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialAudioGainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) InputEndAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputEndAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) InputEndActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputEndActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) InputLossBehavior() MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference {
	var returns MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference
	_jsii_.Get(
		j,
		"inputLossBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) InputLossBehaviorInput() *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehavior {
	var returns *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehavior
	_jsii_.Get(
		j,
		"inputLossBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) InternalValue() *MedialiveChannelEncoderSettingsGlobalConfiguration {
	var returns *MedialiveChannelEncoderSettingsGlobalConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) OutputLockingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputLockingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) OutputLockingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputLockingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) OutputTimingSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTimingSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) OutputTimingSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTimingSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) SupportLowFramerateInputs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportLowFramerateInputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) SupportLowFramerateInputsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportLowFramerateInputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMedialiveChannelEncoderSettingsGlobalConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsGlobalConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsGlobalConfigurationOutputReference_Override(m MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference)SetInitialAudioGain(val *float64) {
	if err := j.validateSetInitialAudioGainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialAudioGain",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference)SetInputEndAction(val *string) {
	if err := j.validateSetInputEndActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputEndAction",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference)SetInternalValue(val *MedialiveChannelEncoderSettingsGlobalConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference)SetOutputLockingMode(val *string) {
	if err := j.validateSetOutputLockingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputLockingMode",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference)SetOutputTimingSource(val *string) {
	if err := j.validateSetOutputTimingSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputTimingSource",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference)SetSupportLowFramerateInputs(val *string) {
	if err := j.validateSetSupportLowFramerateInputsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportLowFramerateInputs",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) PutInputLossBehavior(value *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehavior) {
	if err := m.validatePutInputLossBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putInputLossBehavior",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) ResetInitialAudioGain() {
	_jsii_.InvokeVoid(
		m,
		"resetInitialAudioGain",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) ResetInputEndAction() {
	_jsii_.InvokeVoid(
		m,
		"resetInputEndAction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) ResetInputLossBehavior() {
	_jsii_.InvokeVoid(
		m,
		"resetInputLossBehavior",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) ResetOutputLockingMode() {
	_jsii_.InvokeVoid(
		m,
		"resetOutputLockingMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) ResetOutputTimingSource() {
	_jsii_.InvokeVoid(
		m,
		"resetOutputTimingSource",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) ResetSupportLowFramerateInputs() {
	_jsii_.InvokeVoid(
		m,
		"resetSupportLowFramerateInputs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

