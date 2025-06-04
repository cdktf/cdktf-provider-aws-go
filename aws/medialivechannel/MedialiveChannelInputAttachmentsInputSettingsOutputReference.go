// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelInputAttachmentsInputSettingsOutputReference interface {
	cdktf.ComplexObject
	AudioSelector() MedialiveChannelInputAttachmentsInputSettingsAudioSelectorList
	AudioSelectorInput() interface{}
	CaptionSelector() MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorList
	CaptionSelectorInput() interface{}
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
	DeblockFilter() *string
	SetDeblockFilter(val *string)
	DeblockFilterInput() *string
	DenoiseFilter() *string
	SetDenoiseFilter(val *string)
	DenoiseFilterInput() *string
	FilterStrength() *float64
	SetFilterStrength(val *float64)
	FilterStrengthInput() *float64
	// Experimental.
	Fqn() *string
	InputFilter() *string
	SetInputFilter(val *string)
	InputFilterInput() *string
	InternalValue() *MedialiveChannelInputAttachmentsInputSettings
	SetInternalValue(val *MedialiveChannelInputAttachmentsInputSettings)
	NetworkInputSettings() MedialiveChannelInputAttachmentsInputSettingsNetworkInputSettingsOutputReference
	NetworkInputSettingsInput() *MedialiveChannelInputAttachmentsInputSettingsNetworkInputSettings
	Scte35Pid() *float64
	SetScte35Pid(val *float64)
	Scte35PidInput() *float64
	Smpte2038DataPreference() *string
	SetSmpte2038DataPreference(val *string)
	Smpte2038DataPreferenceInput() *string
	SourceEndBehavior() *string
	SetSourceEndBehavior(val *string)
	SourceEndBehaviorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VideoSelector() MedialiveChannelInputAttachmentsInputSettingsVideoSelectorOutputReference
	VideoSelectorInput() *MedialiveChannelInputAttachmentsInputSettingsVideoSelector
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
	PutAudioSelector(value interface{})
	PutCaptionSelector(value interface{})
	PutNetworkInputSettings(value *MedialiveChannelInputAttachmentsInputSettingsNetworkInputSettings)
	PutVideoSelector(value *MedialiveChannelInputAttachmentsInputSettingsVideoSelector)
	ResetAudioSelector()
	ResetCaptionSelector()
	ResetDeblockFilter()
	ResetDenoiseFilter()
	ResetFilterStrength()
	ResetInputFilter()
	ResetNetworkInputSettings()
	ResetScte35Pid()
	ResetSmpte2038DataPreference()
	ResetSourceEndBehavior()
	ResetVideoSelector()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelInputAttachmentsInputSettingsOutputReference
type jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) AudioSelector() MedialiveChannelInputAttachmentsInputSettingsAudioSelectorList {
	var returns MedialiveChannelInputAttachmentsInputSettingsAudioSelectorList
	_jsii_.Get(
		j,
		"audioSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) AudioSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"audioSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) CaptionSelector() MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorList {
	var returns MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorList
	_jsii_.Get(
		j,
		"captionSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) CaptionSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captionSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) DeblockFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deblockFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) DeblockFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deblockFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) DenoiseFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denoiseFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) DenoiseFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denoiseFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) FilterStrength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filterStrength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) FilterStrengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filterStrengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) InputFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) InputFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) InternalValue() *MedialiveChannelInputAttachmentsInputSettings {
	var returns *MedialiveChannelInputAttachmentsInputSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) NetworkInputSettings() MedialiveChannelInputAttachmentsInputSettingsNetworkInputSettingsOutputReference {
	var returns MedialiveChannelInputAttachmentsInputSettingsNetworkInputSettingsOutputReference
	_jsii_.Get(
		j,
		"networkInputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) NetworkInputSettingsInput() *MedialiveChannelInputAttachmentsInputSettingsNetworkInputSettings {
	var returns *MedialiveChannelInputAttachmentsInputSettingsNetworkInputSettings
	_jsii_.Get(
		j,
		"networkInputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) Scte35Pid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scte35Pid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) Scte35PidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scte35PidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) Smpte2038DataPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smpte2038DataPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) Smpte2038DataPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smpte2038DataPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) SourceEndBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEndBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) SourceEndBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEndBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) VideoSelector() MedialiveChannelInputAttachmentsInputSettingsVideoSelectorOutputReference {
	var returns MedialiveChannelInputAttachmentsInputSettingsVideoSelectorOutputReference
	_jsii_.Get(
		j,
		"videoSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) VideoSelectorInput() *MedialiveChannelInputAttachmentsInputSettingsVideoSelector {
	var returns *MedialiveChannelInputAttachmentsInputSettingsVideoSelector
	_jsii_.Get(
		j,
		"videoSelectorInput",
		&returns,
	)
	return returns
}


func NewMedialiveChannelInputAttachmentsInputSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelInputAttachmentsInputSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelInputAttachmentsInputSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelInputAttachmentsInputSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelInputAttachmentsInputSettingsOutputReference_Override(m MedialiveChannelInputAttachmentsInputSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelInputAttachmentsInputSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference)SetDeblockFilter(val *string) {
	if err := j.validateSetDeblockFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deblockFilter",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference)SetDenoiseFilter(val *string) {
	if err := j.validateSetDenoiseFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"denoiseFilter",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference)SetFilterStrength(val *float64) {
	if err := j.validateSetFilterStrengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterStrength",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference)SetInputFilter(val *string) {
	if err := j.validateSetInputFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFilter",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference)SetInternalValue(val *MedialiveChannelInputAttachmentsInputSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference)SetScte35Pid(val *float64) {
	if err := j.validateSetScte35PidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scte35Pid",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference)SetSmpte2038DataPreference(val *string) {
	if err := j.validateSetSmpte2038DataPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smpte2038DataPreference",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference)SetSourceEndBehavior(val *string) {
	if err := j.validateSetSourceEndBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceEndBehavior",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) PutAudioSelector(value interface{}) {
	if err := m.validatePutAudioSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAudioSelector",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) PutCaptionSelector(value interface{}) {
	if err := m.validatePutCaptionSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCaptionSelector",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) PutNetworkInputSettings(value *MedialiveChannelInputAttachmentsInputSettingsNetworkInputSettings) {
	if err := m.validatePutNetworkInputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNetworkInputSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) PutVideoSelector(value *MedialiveChannelInputAttachmentsInputSettingsVideoSelector) {
	if err := m.validatePutVideoSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putVideoSelector",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ResetAudioSelector() {
	_jsii_.InvokeVoid(
		m,
		"resetAudioSelector",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ResetCaptionSelector() {
	_jsii_.InvokeVoid(
		m,
		"resetCaptionSelector",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ResetDeblockFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetDeblockFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ResetDenoiseFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetDenoiseFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ResetFilterStrength() {
	_jsii_.InvokeVoid(
		m,
		"resetFilterStrength",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ResetInputFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetInputFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ResetNetworkInputSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetNetworkInputSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ResetScte35Pid() {
	_jsii_.InvokeVoid(
		m,
		"resetScte35Pid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ResetSmpte2038DataPreference() {
	_jsii_.InvokeVoid(
		m,
		"resetSmpte2038DataPreference",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ResetSourceEndBehavior() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceEndBehavior",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ResetVideoSelector() {
	_jsii_.InvokeVoid(
		m,
		"resetVideoSelector",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsInputSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

